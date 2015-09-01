package core

import (
	"encoding/json"
	"errors"
	"time"

	"../utils"

	dapi "./docker-api"
	dguard "github.com/90TechSAS/libgo-docker-guard"
)

/*
	List of probe's containers
*/
var (
	ContainerList map[string]*dguard.Container
)

/*
	Initialize Core
*/
func Init() {
	// API Client
	l.Verbose("Init api-client")
	InitAPIClient()
	l.Verbose("api-client OK")

	// Test Docker API
	l.Verbose("Test Docker API")
	TestDockerAPI()
	l.Verbose("Docker API OK")

	// Init ContainersList
	ContainerList = make(map[string]*dguard.Container)

	// Run HTTP Server
	l.Verbose("Run HTTP Server")
	go RunHTTPServer()

	// Run Stats controller
	go StatsController()

	// Run Storage controller
	go StorageController()

	// Refresh container list
	for {
		RefreshContainerList()
		time.Sleep(time.Second * time.Duration(DGConfig.DockerGuard.RefreshContainerListInterval))
	}
}

/*
	Test Docker API connectivity
*/
func TestDockerAPI() {
	var status int                       // HTTP status returned
	var body string                      // HTTP body returned
	var err error                        // Error handling
	var dockerVersion dapi.DockerVersion // DockerVersion struct

	// Get /version on API
	status, body = HTTPReq("/version")
	if status != 200 {
		l.Critical("Can't get docker version, status:", status)
	}

	// Parse returned json
	err = json.Unmarshal([]byte(body), &dockerVersion)
	if err != nil {
		l.Critical("Parsing docker version error:", err)
	}

	// Display version infos
	l.Info("Docker API connection OK:")
	l.Info("\tAPIVersion:\t", dockerVersion.APIVersion)
	l.Info("\tArch:\t\t", dockerVersion.Arch)
	l.Info("\tExperimental:\t", dockerVersion.Experimental)
	l.Info("\tGitCommit:\t", dockerVersion.GitCommit)
	l.Info("\tGoVersion:\t", dockerVersion.GoVersion)
	l.Info("\tKernelVersion:\t", dockerVersion.KernelVersion)
	l.Info("\tOs:\t\t", dockerVersion.Os)
	l.Info("\tVersion:\t", dockerVersion.Version)
}

/*
	Refresh core.ContainerList
*/
func RefreshContainerList() error {
	var tmpContainerArray []dapi.ContainerShort               // Temporary container array
	var tmpContainerList = make(map[string]*dguard.Container) // Temporary container list
	var tmpContainer *dguard.Container                        // Temporary container
	var status int                                            // HTTP status returned
	var body string                                           // HTTP body returned
	var err error                                             // Error handling
	var ok bool                                               // map getter returned

	// Get container list
	l.Debug("RefreshContainerList: Get tmpContainerArray")
	status, body = HTTPReq("/containers/json?all=1")
	if status != 200 {
		l.Error("RefreshContainerList: Can't get container list, status:", status)
		return errors.New("Can't get container list, status: " + utils.I2S(status))
	}

	// Parse returned json
	err = json.Unmarshal([]byte(body), &tmpContainerArray)
	if err != nil {
		l.Error("RefreshContainerList: Parsing container list error:", err)
		return err
	}
	l.Debug("RefreshContainerList: Get tmpContainerArray OK")

	// Create container list
	l.Debug("RefreshContainerList: tmpContainerList")
	for _, dapiContainer := range tmpContainerArray {
		var tmpDAPIContainer dapi.Container  // Temporary DAPI container
		var tmpDGSContainer dguard.Container // Temporary DGS container

		// Get container infos
		status, body = HTTPReq("/containers/" + dapiContainer.ID + "/json")
		if status != 200 {
			l.Error("RefreshContainerList: Can't get container list, status:", status)
			continue
		}

		// Parse returned json
		err = json.Unmarshal([]byte(body), &tmpDAPIContainer)
		if err != nil {
			l.Error("RefreshContainerList: Parsing container list error:", err)
			continue
		}

		// Set tmpDGSContainer fields
		tmpDGSContainer.ID = tmpDAPIContainer.ID
		tmpDGSContainer.Hostname = tmpDAPIContainer.Config.Hostname
		tmpDGSContainer.Image = tmpDAPIContainer.Config.Image
		tmpDGSContainer.IPAddress = tmpDAPIContainer.NetworkSettings.IPAddress
		tmpDGSContainer.MacAddress = tmpDAPIContainer.NetworkSettings.MacAddress
		tmpDGSContainer.SizeRootFs = 0
		tmpDGSContainer.SizeRw = 0
		tmpDGSContainer.MemoryUsed = 0
		tmpDGSContainer.Running = tmpDAPIContainer.State.Running

		// Add tmpDGSContainer to the tmpContainerList
		tmpContainerList[tmpDGSContainer.ID] = &tmpDGSContainer

		// if tmpDGSContainer already exists: update status
		tmpC, ok := ContainerList[tmpDGSContainer.ID]
		if ok {
			tmpC.Running = tmpDGSContainer.Running
			ContainerResetTime(tmpDGSContainer.ID)
			ContainerList[tmpDGSContainer.ID] = tmpC
		}
	}
	l.Debug("RefreshContainerList: tmpContainerList OK")

	// Check new containers
	l.Debug("RefreshContainerList: Check new containers")
	for _, tmpContainer = range tmpContainerList {
		// If containers does not exist: add it
		_, ok = ContainerList[tmpContainer.ID]
		if !ok {
			// Add tmpContainer in ContainerList
			ContainerList[tmpContainer.ID] = tmpContainer
			l.Debug("RefreshContainerList: Container", tmpContainer.ID, "added in ContainerList")
		}
	}
	l.Debug("RefreshContainerList: Check new containers OK")

	// Check removed containers
	l.Debug("RefreshContainerList: Check removed containers")
	for _, tmpContainer = range ContainerList {
		// If containers does not exist: delete it
		_, ok = tmpContainerList[tmpContainer.ID]
		if !ok {
			// Remove tmpContainer in ContainerList
			delete(ContainerList, tmpContainer.ID)
			l.Debug("RefreshContainerList: Container", tmpContainer.ID, "removed in ContainerList")
		}
	}
	l.Debug("RefreshContainerList: Check removed containers OK")

	return nil
}

/*
	Set the RootFs size of a container in the ContainerList
*/
func SetContainerSizeRootFs(id string, size float64) bool {
	var ok bool
	_, ok = ContainerList[id]
	if ok {
		ContainerList[id].SizeRootFs = size
		return true
	}
	return false
}

/*
	Set the Rw size of a container in the ContainerList
*/
func SetContainerSizeRw(id string, size float64) bool {
	var ok bool
	_, ok = ContainerList[id]
	if ok {
		ContainerList[id].SizeRw = size
		return true
	}
	return false
}

/*
	Set the memory used size of a container in the ContainerList
*/
func SetContainerMemoryUsed(id string, size float64) bool {
	var ok bool
	_, ok = ContainerList[id]
	if ok {
		ContainerList[id].MemoryUsed = size
		return true
	}
	return false
}

/*
	Set the memory used size of a container in the ContainerList
*/
func ContainerResetTime(id string) bool {
	var ok bool
	_, ok = ContainerList[id]
	if ok {
		ContainerList[id].Time = float64(time.Now().Unix())
		return true
	}
	return false
}

/*
	Set the net bandwith RX of a container in the ContainerList
*/
func SetContainerNetBandwithRX(id string, size float64) bool {
	var ok bool
	_, ok = ContainerList[id]
	if ok {
		ContainerList[id].NetBandwithRX = size
		return true
	}
	return false
}

/*
	Set the net bandwith TX of a container in the ContainerList
*/
func SetContainerNetBandwithTX(id string, size float64) bool {
	var ok bool
	_, ok = ContainerList[id]
	if ok {
		ContainerList[id].NetBandwithTX = size
		return true
	}
	return false
}

/*
	Set the cpu usage of a container in the ContainerList
*/
func SetContainerCPUUsage(id string, size float64) bool {
	var ok bool
	_, ok = ContainerList[id]
	if ok {
		ContainerList[id].CPUUsage = size
		return true
	}
	return false
}

package core

import (
	"encoding/json"
	"fmt"
	"net/http"

	dapi "./docker-api"

	dguard "github.com/90TechSAS/libgo-docker-guard"
)

func HTTPHandlerList(w http.ResponseWriter, r *http.Request) {
	var returnStr string
	var tmpContainerList []dguard.Container // Temporary DGC Container list

	// Refresh Container list
	RefreshContainerList()

	// Browse container list
	for _, container := range ContainerList {
		var tmpDAPIContainer dapi.Container  // Temporary Docker API Container
		var tmpDGSContainer dguard.Container // Temporary DGS Container
		var status int                       // HTTP status returned
		var body string                      // HTTP body returned
		var err error                        // Error handling

		// Get container info
		status, body = HTTPReq("/containers/" + container.ID + "/json")
		if status != 200 {
			l.Error("Can't get docker container (", container.ID, ") info, status:", status)
		}

		// Parse returned json
		err = json.Unmarshal([]byte(body), &tmpDAPIContainer)
		if err != nil {
			l.Error("Parsing docker container (", container.ID, ") info:", err)
		}

		// Set tmpDGSContainer
		tmpDGSContainer.ID = tmpDAPIContainer.ID
		tmpDGSContainer.Hostname = tmpDAPIContainer.Config.Hostname
		tmpDGSContainer.Image = tmpDAPIContainer.Image
		tmpDGSContainer.IPAddress = tmpDAPIContainer.NetworkSettings.IPAddress
		tmpDGSContainer.MacAddress = tmpDAPIContainer.NetworkSettings.MacAddress
		tmpDGSContainer.State.Dead = tmpDAPIContainer.State.Dead
		tmpDGSContainer.State.Error = tmpDAPIContainer.State.Error
		tmpDGSContainer.State.ExitCode = tmpDAPIContainer.State.ExitCode
		tmpDGSContainer.State.FinishedAt = tmpDAPIContainer.State.FinishedAt
		tmpDGSContainer.State.OOMKilled = tmpDAPIContainer.State.OOMKilled
		tmpDGSContainer.State.Paused = tmpDAPIContainer.State.Paused
		tmpDGSContainer.State.Pid = tmpDAPIContainer.State.Pid
		tmpDGSContainer.State.Restarting = tmpDAPIContainer.State.Restarting
		tmpDGSContainer.State.Running = tmpDAPIContainer.State.Running
		tmpDGSContainer.State.StartedAt = tmpDAPIContainer.State.StartedAt

		containerStorage, ok := ContainerStorage[tmpDGSContainer.ID]
		if ok {
			tmpDGSContainer.SizeRootFs = containerStorage[0]
			tmpDGSContainer.SizeRw = containerStorage[1]
		} else {
			tmpDGSContainer.SizeRootFs = 0
			tmpDGSContainer.SizeRw = 0
		}

		// Add tmpDGSContainer to tmpContainerList
		tmpContainerList = append(tmpContainerList, tmpDGSContainer)
	}

	// tmpContainerList => json
	tmpJson, _ := json.Marshal(tmpContainerList)

	// Add json to the returned string
	returnStr = string(tmpJson)

	fmt.Fprint(w, returnStr)
}

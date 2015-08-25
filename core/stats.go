package core

import (
	"encoding/json"
	"sync"
	"time"

	dapi "./docker-api"
)

func StatsController() {
	var status int                               // HTTP status returned
	var body string                              // HTTP body returned
	var err error                                // Error handling
	var tmpContainerArray []*dapi.ContainerShort // Temporary container array
	var wg sync.WaitGroup

	for {
		// Get container list
		l.Debug("StatsController: Get tmpContainerArray")
		status, body = HTTPReq("/containers/json?all=1")
		if status != 200 {
			l.Error("StatsController: Can't get container list, status:", status)
			time.Sleep(time.Second * 5)
			continue
		}

		// Parse returned json
		err = json.Unmarshal([]byte(body), &tmpContainerArray)
		if err != nil {
			l.Error("StatsController: Parsing container list error:", err)
			time.Sleep(time.Second * 5)
			continue
		}
		l.Debug("StatsController: Get tmpContainerArray OK")

		// If no container => stop
		if len(tmpContainerArray) == 0 {
			time.Sleep(time.Second * time.Duration(DGConfig.DockerGuard.StorageControllerPause))
			continue
		}

		// Get containers' stats
		for i := 0; i < len(tmpContainerArray); i++ {
			tmpContainer := tmpContainerArray[i]

			// Get asynchronously stats
			wg.Add(1)
			go GetStats(tmpContainer, &wg)

			// Pause 1sec * StorageControllerInterval
			time.Sleep(time.Second * time.Duration(DGConfig.DockerGuard.StorageControllerInterval))
		}

		// Wait stats getters
		wg.Wait()

		// Pause 1sec * StorageControllerPause
		l.Silly("StatsController: End getting containers storage usage")
		time.Sleep(time.Second * time.Duration(DGConfig.DockerGuard.StorageControllerPause))
	}

}

/*
	Get container's stats
*/
func GetStats(container *dapi.ContainerShort, wg *sync.WaitGroup) {
	var tmpDAPIContainerS dapi.ContainerStats // Temporary DAPI container stats
	var status int                            // HTTP status returned
	var body string                           // HTTP body returned
	var err error                             // Error handling

	defer func() {
		wg.Done()
	}()

	// Get container stats
	status, body = HTTPReq("/containers/" + container.ID + "/stats?stream=0")
	if status != 200 {
		l.Error("StatsController: Can't get docker container (", container.ID, ") stats, status:", status)
		return
	} else {
		// Parse container returned json
		err = json.Unmarshal([]byte(body), &tmpDAPIContainerS)
		if err != nil {
			l.Error("StatsController: Parsing docker container (", container.ID, ") stats:", err)
			return
		}
	}

	// Add values to map
	if status == 200 {
		l.Debug("StatsController: Add values to map")
		SetContainerMemoryUsed(container.ID, float64(tmpDAPIContainerS.MemoryStats.Usage))
		SetContainerNetBandwithRX(container.ID, float64(tmpDAPIContainerS.Network.RxBytes))
		SetContainerNetBandwithTX(container.ID, float64(tmpDAPIContainerS.Network.TxBytes))
		SetContainerCPUUsage(container.ID, float64(tmpDAPIContainerS.CPUStats.CPUUsage.TotalUsage))
		ContainerResetTime(container.ID)
		l.Debug("StatsController: Add values to map OK")
	}
}

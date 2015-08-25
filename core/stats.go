package core

import (
	"encoding/json"
	"time"

	dapi "./docker-api"
)

func StatsController() {
	var status int                              // HTTP status returned
	var body string                             // HTTP body returned
	var err error                               // Error handling
	var tmpContainerArray []dapi.ContainerShort // Temporary container array
	var tmpDAPIContainerS dapi.ContainerStats   // Temporary DAPI container stats

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

		for i := 0; i < len(tmpContainerArray); i++ {
			l.Silly("StatsController: Get", tmpContainerArray[i], "storage usage")

			// Get container stats
			status, body = HTTPReq("/containers/" + tmpContainerArray[i].ID + "/stats?stream=0")
			if status != 200 {
				l.Error("StatsController: Can't get docker container (", tmpContainerArray[i].ID, ") stats, status:", status)
				continue
			} else {
				// Parse container returned json
				err = json.Unmarshal([]byte(body), &tmpDAPIContainerS)
				if err != nil {
					l.Error("StatsController: Parsing docker container (", tmpContainerArray[i].ID, ") stats:", err)
					continue
				}
			}

			// Add values to map
			if status == 200 {
				l.Debug("StatsController: Add values to map")
				SetContainerMemoryUsed(tmpContainerArray[i].ID, float64(tmpDAPIContainerS.MemoryStats.Usage))
				SetContainerNetBandwithRX(tmpContainerArray[i].ID, float64(tmpDAPIContainerS.Network.RxBytes))
				SetContainerNetBandwithTX(tmpContainerArray[i].ID, float64(tmpDAPIContainerS.Network.TxBytes))
				SetContainerCPUUsage(tmpContainerArray[i].ID, float64(tmpDAPIContainerS.CPUStats.CPUUsage.TotalUsage))
				ContainerResetTime(tmpContainerArray[i].ID)
				l.Debug("StatsController: Add values to map OK")
			}

			// Pause 1sec * StorageControllerInterval
			time.Sleep(time.Second * time.Duration(DGConfig.DockerGuard.StorageControllerInterval))
		}

		// Pause 1sec * StorageControllerPause
		l.Silly("StatsController: End getting containers storage usage")
		time.Sleep(time.Second * time.Duration(DGConfig.DockerGuard.StorageControllerPause))
	}

}

package core

import (
	"encoding/json"
	"time"

	dapi "./docker-api"
)

func StatsController() {
	var status int                                 // HTTP status returned
	var body string                                // HTTP body returned
	var err error                                  // Error handling
	var tmpContainerArray []dapi.ContainerShort    // Temporary container array
	var tmpDAPIContainer []dapi.ContainerShortSize // Temporary DAPI container
	var tmpDAPIContainerS dapi.ContainerStats      // Temporary DAPI container stats
	var previousContainerID string                 // Previous Container ID

	for {
		// Get container list
		l.Debug("StatsController: Get tmpContainerArray")
		status, body = HTTPReq("/containers/json?all=1")
		if status != 200 {
			l.Error("Can't get container list, status:", status)
			time.Sleep(time.Second * 5)
			continue
		}

		// Parse returned json
		err = json.Unmarshal([]byte(body), &tmpContainerArray)
		if err != nil {
			l.Error("Parsing container list error:", err)
			time.Sleep(time.Second * 5)
			continue
		}
		l.Debug("StatsController: Get tmpContainerArray OK")

		// If no container => stop
		if len(tmpContainerArray) == 0 {
			time.Sleep(time.Second * time.Duration(DGConfig.DockerGuard.StorageControllerPause))
			continue
		}

		// Get first container info
		l.Debug("StatsController: Get first container info")
		status, body = HTTPReq("/containers/json?all=1&size=1&limit=1")
		if status != 200 {
			l.Error("Can't get docker container (", tmpContainerArray[0].ID, ") info, status:", status)
		}

		// Parse first container returned json
		err = json.Unmarshal([]byte(body), &tmpDAPIContainer)
		if err != nil {
			l.Error("Parsing docker container (", tmpContainerArray[0].ID, ") info:", err)
		}
		l.Debug("StatsController: Get first container info OK")

		// Get first container stats
		l.Debug("StatsController: Get first container stats")
		status, body = HTTPReq("/containers/" + tmpContainerArray[0].ID + "/stats?stream=0")
		if status != 200 {
			l.Error("Can't get docker container (", tmpContainerArray[0].ID, ") stats, status:", status)
		}

		// Parse first container returned json
		err = json.Unmarshal([]byte(body), &tmpDAPIContainerS)
		if err != nil {
			l.Error("Parsing docker container (", tmpContainerArray[0].ID, ") stats:", err)
		}
		l.Debug("StatsController: Get first container stats OK")

		// Add values to map
		l.Debug("StatsController: Add values to map")
		previousContainerID = tmpContainerArray[0].ID
		SetContainerSizeRootFs(previousContainerID, tmpDAPIContainer[0].SizeRootFs)
		SetContainerSizeRw(previousContainerID, tmpDAPIContainer[0].SizeRw)
		SetContainerMemoryUsed(previousContainerID, float64(tmpDAPIContainerS.MemoryStats.Usage))
		l.Debug("StatsController: Add values to map OK")
		for i := 1; i < len(tmpContainerArray); i++ {
			l.Silly("Get", tmpContainerArray[i], "storage usage")

			// Get container info
			status, body = HTTPReq("/containers/json?all=1&size=1&limit=1&before=" + previousContainerID)
			if status != 200 {
				l.Error("Can't get docker container (", tmpContainerArray[i], ") info, status:", status)
			}

			// Parse returned json
			err = json.Unmarshal([]byte(body), &tmpDAPIContainer)
			if err != nil {
				l.Error("Parsing docker container (", tmpContainerArray[i], ") info:", err)
			}

			// Get container stats
			status, body = HTTPReq("/containers/" + tmpContainerArray[0].ID + "/stats?stream=0")
			if status != 200 {
				l.Error("Can't get docker container (", tmpContainerArray[0].ID, ") stats, status:", status)
			}

			// Parse container returned json
			err = json.Unmarshal([]byte(body), &tmpDAPIContainerS)
			if err != nil {
				l.Error("Parsing docker container (", tmpContainerArray[0].ID, ") stats:", err)
			}

			// Add values to map
			previousContainerID = tmpContainerArray[i].ID
			SetContainerSizeRootFs(previousContainerID, tmpDAPIContainer[0].SizeRootFs)
			SetContainerSizeRw(previousContainerID, tmpDAPIContainer[0].SizeRw)
			SetContainerMemoryUsed(previousContainerID, float64(tmpDAPIContainerS.MemoryStats.Usage))

			// Pause 1sec * StorageControllerInterval
			time.Sleep(time.Second * time.Duration(DGConfig.DockerGuard.StorageControllerInterval))
		}

		// Pause 1sec * StorageControllerPause
		l.Silly("End getting containers storage usage")
		time.Sleep(time.Second * time.Duration(DGConfig.DockerGuard.StorageControllerPause))
	}

}

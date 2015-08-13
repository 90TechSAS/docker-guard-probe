package core

import (
	"encoding/json"
	"time"

	dapi "./docker-api"
)

var (
	// Container's storage usage
	ContainerStorage map[string][]float64
)

/*
	Get containers' storage usage
*/
func RunStorageController() {
	var status int                                 // HTTP status returned
	var body string                                // HTTP body returned
	var err error                                  // Error handling
	var tmpContainerList []dapi.ContainerShort     // Temporary DGC Container list
	var tmpDAPIContainer []dapi.ContainerShortSize // Temporary Docker API Container
	var previousContainerID string                 // Previous Container ID

	for {
		var tmpContainerStorage map[string][]float64 = make(map[string][]float64) // Temporary container storage

		l.Silly("Start getting containers storage usage")

		// Refresh Container list
		RefreshContainerList()
		tmpContainerList = ContainerList

		// Get first container info
		status, body = HTTPReq("/containers/json?all=1&size=1&limit=1")
		if status != 200 {
			l.Error("Can't get docker container (", tmpContainerList[0].ID, ") info, status:", status)
		}

		// Parse first container returned json
		err = json.Unmarshal([]byte(body), &tmpDAPIContainer)
		if err != nil {
			l.Error("Parsing docker container (", tmpContainerList[0].ID, ") info:", err)
		}

		// If no container => stop
		if len(tmpContainerList) == 0 {
			time.Sleep(time.Second * time.Duration(DGConfig.DockerGuard.StorageControllerPause))
			continue
		}

		// Add values to map
		previousContainerID = tmpContainerList[0].ID
		tmpContainerStorage[previousContainerID] = make([]float64, 2)
		tmpContainerStorage[previousContainerID][0] = tmpDAPIContainer[0].SizeRootFs
		tmpContainerStorage[previousContainerID][1] = tmpDAPIContainer[0].SizeRw

		for i := 1; i < len(tmpContainerList); i++ {
			l.Silly("Get", tmpContainerList[i], "storage usage")

			// Get container info
			status, body = HTTPReq("/containers/json?all=1&size=1&limit=1&before=" + previousContainerID)
			if status != 200 {
				l.Error("Can't get docker container (", tmpContainerList[i], ") info, status:", status)
			}

			// Parse returned json
			err = json.Unmarshal([]byte(body), &tmpDAPIContainer)
			if err != nil {
				l.Error("Parsing docker container (", tmpContainerList[i], ") info:", err)
			}

			// Add values to map
			previousContainerID = tmpContainerList[i].ID
			tmpContainerStorage[previousContainerID] = make([]float64, 2)
			tmpContainerStorage[previousContainerID][0] = float64(tmpDAPIContainer[0].SizeRootFs)
			tmpContainerStorage[previousContainerID][1] = float64(tmpDAPIContainer[0].SizeRw)

			// Pause 1sec * StorageControllerInterval
			time.Sleep(time.Second * time.Duration(DGConfig.DockerGuard.StorageControllerInterval))
		}

		// Set ContainerStorage
		ContainerStorage = tmpContainerStorage

		// Display silly logs about storage
		l.Silly("New container storage list:")
		for i, j := range ContainerStorage {
			l.Silly("\t", i, ":", j[0], "/", j[1])
		}

		// Pause 1sec * StorageControllerPause
		l.Silly("End getting containers storage usage")
		time.Sleep(time.Second * time.Duration(DGConfig.DockerGuard.StorageControllerPause))
	}
}

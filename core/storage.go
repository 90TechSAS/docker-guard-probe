package core

import (
	"encoding/json"
	"time"

	dapi "./docker-api"

	"../utils"
)

/*
	Get containers' disk usage
*/
func StorageController() {
	var status int                              // HTTP status returned
	var body string                             // HTTP body returned
	var err error                               // Error handling
	var tmpContainerArray []dapi.ContainerShort // Temporary container array
	var sizeRW int64                            // Container's Size RW
	var sizeRootFs int64                        // Container's Size RootFs

	for {
		// Get container list
		l.Debug("StorageController: Get tmpContainerArray")
		status, body = HTTPReq("/containers/json?all=1")
		if status != 200 {
			l.Error("StorageController: Can't get container list, status:", status)
			time.Sleep(time.Second * 5)
			continue
		}

		// Parse returned json
		err = json.Unmarshal([]byte(body), &tmpContainerArray)
		if err != nil {
			l.Error("StorageController: Parsing container list error:", err)
			time.Sleep(time.Second * 5)
			continue
		}
		l.Debug("StorageController: Get tmpContainerArray OK")

		// If no container => stop
		if len(tmpContainerArray) == 0 {
			time.Sleep(time.Second * time.Duration(DGConfig.DockerGuard.StorageControllerPause))
			continue
		}

		for i := 0; i < len(tmpContainerArray); i++ {
			// Get container sizes
			sizeRW, err = utils.DirectorySize("/var/lib/docker/aufs/diff/" + tmpContainerArray[i].ID)
			if err != nil {
				l.Error("StorageController: Can't get container (", tmpContainerArray[i].ID, ") SizeRootFs:", err)
			}
			sizeRootFs, err = utils.DirectorySize("/var/lib/docker/aufs/mnt/" + tmpContainerArray[i].ID)
			if err != nil {
				l.Error("StorageController: Can't get container (", tmpContainerArray[i].ID, ") SizeRW:", err)
			}

			// Add values to map
			if status == 200 {
				l.Debug("StorageController: Add values to map")
				SetContainerSizeRootFs(tmpContainerArray[i].ID, float64(sizeRootFs))
				SetContainerSizeRw(tmpContainerArray[i].ID, float64(sizeRW))
				ContainerResetTime(tmpContainerArray[i].ID)
				l.Debug("StorageController: Add values to map OK")
			}

			// Pause 1sec * StorageControllerInterval
			time.Sleep(time.Second * time.Duration(DGConfig.DockerGuard.StorageControllerInterval))
		}

		// Pause 1sec * StorageControllerPause
		l.Silly("StorageController: End getting containers storage usage")
		time.Sleep(time.Second * time.Duration(DGConfig.DockerGuard.StorageControllerPause))
	}
}

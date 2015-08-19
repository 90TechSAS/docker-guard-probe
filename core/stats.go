package core

import (
	"encoding/json"
	"time"

	dapi "./docker-api"

	"../utils"
)

func StatsController() {
	var status int                              // HTTP status returned
	var body string                             // HTTP body returned
	var err error                               // Error handling
	var tmpContainerArray []dapi.ContainerShort // Temporary container array
	var tmpDAPIContainerS dapi.ContainerStats   // Temporary DAPI container stats
	var sizeRW int64                            // Container's Size RW
	var sizeRootFs int64                        // Container's Size RootFs

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
			l.Silly("Get", tmpContainerArray[i], "storage usage")

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

			// Get container sizes
			sizeRW, err = utils.DirectorySize("/var/lib/docker/aufs/diff/" + tmpContainerArray[i].ID)
			if err != nil {
				l.Error("StatsController: Can't get container (", tmpContainerArray[i].ID, ") SizeRootFs:", err)
			}
			sizeRootFs, err = utils.DirectorySize("/var/lib/docker/aufs/mnt/" + tmpContainerArray[i].ID)
			if err != nil {
				l.Error("StatsController: Can't get container (", tmpContainerArray[i].ID, ") SizeRW:", err)
			}

			// Add values to map
			if status == 200 {
				l.Debug("StatsController: Add values to map")
				SetContainerSizeRootFs(tmpContainerArray[i].ID, float64(sizeRootFs))
				SetContainerSizeRw(tmpContainerArray[i].ID, float64(sizeRW))
				SetContainerMemoryUsed(tmpContainerArray[i].ID, float64(tmpDAPIContainerS.MemoryStats.Usage))
				ContainerResetTime(tmpContainerArray[i].ID)
				l.Debug("StatsController: Add values to map OK")
			}

			// Pause 1sec * StorageControllerInterval
			time.Sleep(time.Second * time.Duration(DGConfig.DockerGuard.StorageControllerInterval))
		}

		// Pause 1sec * StorageControllerPause
		l.Silly("End getting containers storage usage")
		time.Sleep(time.Second * time.Duration(DGConfig.DockerGuard.StorageControllerPause))
	}

}

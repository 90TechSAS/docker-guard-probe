package core

import (
	"encoding/json"
	"strings"
	"sync"
	"time"

	dapi "./docker-api"
)

/*
	Previous stats
*/
type OldStats struct {
	Filled         bool
	NetRX          float64
	NetTX          float64
	CPUUsage       float64
	SystemCPUUsage float64
	Time           time.Time
}

/*
	Loop to get containers' stats
*/
func StatsController() {
	var status int                               // HTTP status returned
	var body string                              // HTTP body returned
	var err error                                // Error handling
	var tmpContainerArray []*dapi.ContainerShort // Temporary container array
	var oldStats map[string]*OldStats            // Previous stats
	var wg sync.WaitGroup                        // Waiting group for API client

	oldStats = make(map[string]*OldStats)

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

			// Skip getting stats if container stopped.
			if !strings.HasPrefix(tmpContainer.Status, "Up") {
				l.Debug("StatsController: Container", tmpContainer.ID, "is stopped. Get stats skipped.")
				continue
			}

			// Get asynchronously stats
			if oldStats[tmpContainer.ID] == nil {
				oldStats[tmpContainer.ID] = new(OldStats)
			}
			wg.Add(1)
			go GetStats(tmpContainer, oldStats[tmpContainer.ID], &wg)

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
func GetStats(container *dapi.ContainerShort, oldS *OldStats, wg *sync.WaitGroup) {
	var tmpDAPIContainerS dapi.ContainerStats // Temporary DAPI container stats
	var status int                            // HTTP status returned
	var body string                           // HTTP body returned
	var err error                             // Error handling

	defer wg.Done()

	// Get container stats
	status, body = HTTPReq("/containers/" + container.ID + "/stats?stream=0")
	if status != 200 {
		l.Error("StatsController: Can't get docker container (", container.ID, ") stats, status:", status)
		return
	}

	// Parse container returned json
	err = json.Unmarshal([]byte(body), &tmpDAPIContainerS)
	if err != nil {
		l.Error("StatsController: Parsing docker container (", container.ID, ") stats:", err, "\nJSON:", body)
		return
	}

	// Add values to map
	if status == 200 {
		var rxb, txb, cpuu float64
		if oldS.Filled {
			rxb = (float64(tmpDAPIContainerS.Network.RxBytes) - oldS.NetRX) / time.Since(oldS.Time).Seconds()
			txb = (float64(tmpDAPIContainerS.Network.TxBytes) - oldS.NetTX) / time.Since(oldS.Time).Seconds()
			delta1 := float64(tmpDAPIContainerS.CPUStats.CPUUsage.TotalUsage) - oldS.CPUUsage
			delta2 := float64(tmpDAPIContainerS.CPUStats.SystemCPUUsage) - oldS.SystemCPUUsage
			if delta1 > 0.0 && delta2 > 0.0 {
				cpuu = delta1 / delta2 * 100
			}
			if rxb < 0 {
				rxb = 0
			}
			if txb < 0 {
				txb = 0
			}
		} else {
			rxb = 0
			txb = 0
			cpuu = 0
		}

		oldS.NetRX = float64(tmpDAPIContainerS.Network.RxBytes)
		oldS.NetTX = float64(tmpDAPIContainerS.Network.TxBytes)
		oldS.CPUUsage = float64(tmpDAPIContainerS.CPUStats.CPUUsage.TotalUsage)
		oldS.SystemCPUUsage = float64(tmpDAPIContainerS.CPUStats.SystemCPUUsage)
		oldS.Time = time.Now()
		oldS.Filled = true

		l.Debug("StatsController: Add values to map")
		SetContainerMemoryUsed(container.ID, float64(tmpDAPIContainerS.MemoryStats.Usage))
		SetContainerNetBandwithRX(container.ID, rxb)
		SetContainerNetBandwithTX(container.ID, txb)
		SetContainerCPUUsage(container.ID, cpuu)
		ContainerResetTime(container.ID)
		l.Debug("StatsController: Add values to map OK")
	}
}

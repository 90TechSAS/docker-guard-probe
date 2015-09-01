package core

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
	"syscall"

	dapi "./docker-api"
	dguard "github.com/90TechSAS/libgo-docker-guard"

	"../utils"
)

/*
	Handle GET /dockerinfos
*/
func HTTPHandlerDockerinfos(w http.ResponseWriter, r *http.Request) {
	var returnStr string                    // Returned string
	var tmpDockerInfos dapi.DockerInfos     // Docker API infos
	var tmpDockerVersion dapi.DockerVersion // Docker API version
	var dockerInfos dguard.DockerInfos      // DGC Docker infos
	var status int                          // HTTP status returned
	var body string                         // HTTP body returned
	var err error                           // Error handling

	// Get docker info
	status, body = HTTPReq("/info")
	if status != 200 {
		l.Error("Can't get docker info, status:", status)
	}

	// Parse returned json
	err = json.Unmarshal([]byte(body), &tmpDockerInfos)
	if err != nil {
		l.Error("Parsing docker info:", err)
	}

	// Get docker version
	status, body = HTTPReq("/version")
	if status != 200 {
		l.Error("Can't get docker version, status:", status)
	}

	// Parse returned json
	err = json.Unmarshal([]byte(body), &tmpDockerVersion)
	if err != nil {
		l.Error("Parsing docker version:", err)
	}

	// Set dockerInfos
	dockerInfos.ID = tmpDockerInfos.ID
	dockerInfos.Name = tmpDockerInfos.Name
	dockerInfos.Containers = tmpDockerInfos.Containers
	dockerInfos.Images = tmpDockerInfos.Images
	dockerInfos.Driver = tmpDockerInfos.Driver
	dockerInfos.SystemTime = tmpDockerInfos.SystemTime
	dockerInfos.OperatingSystem = tmpDockerInfos.OperatingSystem
	dockerInfos.NCPU = tmpDockerInfos.NCPU
	dockerInfos.MemTotal = tmpDockerInfos.MemTotal

	dockerInfos.APIVersion = tmpDockerVersion.APIVersion
	dockerInfos.Arch = tmpDockerVersion.Arch
	dockerInfos.Experimental = tmpDockerVersion.Experimental
	dockerInfos.GitCommit = tmpDockerVersion.GitCommit
	dockerInfos.GoVersion = tmpDockerVersion.GoVersion
	dockerInfos.KernelVersion = tmpDockerVersion.KernelVersion
	dockerInfos.Os = tmpDockerVersion.Os
	dockerInfos.Version = tmpDockerVersion.Version

	// dockerInfos => json
	tmpJSON, _ := json.Marshal(dockerInfos)

	// Add json to the returned string
	returnStr = string(tmpJSON)

	fmt.Fprint(w, returnStr)
}

/*
	Handle GET /probeinfos
*/
func HTTPHandlerProbeinfos(w http.ResponseWriter, r *http.Request) {
	var returnStr string             // Returned string
	var probeInfos dguard.ProbeInfos // DGC Probe infos
	var stat syscall.Statfs_t        // Syscall to get disk usage
	var out []byte                   // Command output
	var err error                    // Error handling

	// Get load average
	out, err = exec.Command("sh", "-c", "uptime | awk '{printf \"%s%s%s\",$8,$9,$10}'").Output()
	if err != nil {
		l.Error("HTTPHandlerProbeinfos: get load avg:", err)
		http.Error(w, http.StatusText(500), 500)
		return
	}
	probeInfos.LoadAvg = string(out)

	// Get memory usage
	out, err = exec.Command("sh", "-c", "cat /proc/meminfo | grep MemTotal | awk '{printf \"%d\",$2}'").Output()
	if err != nil {
		l.Error("HTTPHandlerProbeinfos: get mem total:", err)
		http.Error(w, http.StatusText(500), 500)
		return
	}
	probeInfos.MemoryTotal, err = utils.S2F(string(out))
	if err != nil {
		l.Error("HTTPHandlerProbeinfos: get mem total S2I:", err)
		http.Error(w, http.StatusText(500), 500)
		return
	}
	out, err = exec.Command("sh", "-c", "cat /proc/meminfo | grep MemAvailable | awk '{printf \"%d\",$2}'").Output()
	if err != nil {
		l.Error("HTTPHandlerProbeinfos: get mem available:", err)
		http.Error(w, http.StatusText(500), 500)
		return
	}
	if err != nil {
		l.Error("HTTPHandlerProbeinfos: get mem available S2I:", err)
		http.Error(w, http.StatusText(500), 500)
		return
	}
	probeInfos.MemoryAvailable, err = utils.S2F(string(out))

	// Get disk usage
	syscall.Statfs("/", &stat)
	probeInfos.DiskTotal = float64(stat.Blocks * uint64(stat.Bsize))
	probeInfos.DiskAvailable = float64(stat.Bavail * uint64(stat.Bsize))

	// probeInfos => json
	tmpJSON, err := json.Marshal(probeInfos)
	if err != nil {
		l.Error("HTTPHandlerProbeinfos: marshal JSON:", err)
		http.Error(w, http.StatusText(500), 500)
	}

	// Add json to the returned string
	returnStr = string(tmpJSON)

	fmt.Fprint(w, returnStr)
}

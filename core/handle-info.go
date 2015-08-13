package core

import (
	"encoding/json"
	"fmt"
	"net/http"

	dapi "./docker-api"

	dguard "github.com/90TechSAS/libgo-docker-guard"
)

func HTTPHandlerInfo(w http.ResponseWriter, r *http.Request) {
	var returnStr string
	var tmpDockerInfo dapi.DockerInfo       // Docker API info
	var tmpDockerVersion dapi.DockerVersion // Docker API version
	var dockerInfo dguard.DockerInfo        // DGC Docker info
	var status int                          // HTTP status returned
	var body string                         // HTTP body returned
	var err error                           // Error handling

	// Get docker info
	status, body = HTTPReq("/info")
	if status != 200 {
		l.Error("Can't get docker info, status:", status)
	}

	// Parse returned json
	err = json.Unmarshal([]byte(body), &tmpDockerInfo)
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

	// Set dockerInfo
	dockerInfo.ID = tmpDockerInfo.ID
	dockerInfo.Name = tmpDockerInfo.Name
	dockerInfo.Containers = tmpDockerInfo.Containers
	dockerInfo.Images = tmpDockerInfo.Images
	dockerInfo.Driver = tmpDockerInfo.Driver
	dockerInfo.SystemTime = tmpDockerInfo.SystemTime
	dockerInfo.OperatingSystem = tmpDockerInfo.OperatingSystem
	dockerInfo.NCPU = tmpDockerInfo.NCPU
	dockerInfo.MemTotal = tmpDockerInfo.MemTotal

	dockerInfo.APIVersion = tmpDockerVersion.APIVersion
	dockerInfo.Arch = tmpDockerVersion.Arch
	dockerInfo.Experimental = tmpDockerVersion.Experimental
	dockerInfo.GitCommit = tmpDockerVersion.GitCommit
	dockerInfo.GoVersion = tmpDockerVersion.GoVersion
	dockerInfo.KernelVersion = tmpDockerVersion.KernelVersion
	dockerInfo.Os = tmpDockerVersion.Os
	dockerInfo.Version = tmpDockerVersion.Version

	// dockerInfo => json
	tmpJson, _ := json.Marshal(dockerInfo)

	// Add json to the returned string
	returnStr = string(tmpJson)

	fmt.Fprint(w, returnStr)
}

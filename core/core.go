package core

import (
	"encoding/json"

	dguard "github.com/90TechSAS/libgo-docker-guard"
)

/*
	Initialize Core
*/
func Init() {
	// API Client
	l.Verbose("Init api-client")
	InitAPIClient()
	l.Verbose("api-client OK")

	// Test Docker API
	l.Verbose("Test Docker API")
	TestDockerAPI()
	l.Verbose("Docker API OK")

	// Run HTTP Server
	l.Verbose("Run HTTP Server")
	RunHTTPServer()
}

/*
	Test Docker API connectivity
*/
func TestDockerAPI() {
	var status int                         // HTTP status returned
	var body string                        // HTTP body returned
	var err error                          // Error handling
	var dockerVersion dguard.DockerVersion // DockerVersion struct

	// Get /version on API
	status, body = HTTPReq("/version")
	if status != 200 {
		l.Critical("Can't get docker version, status:", status)
	}

	// Parse returned json
	err = json.Unmarshal([]byte(body), &dockerVersion)
	if err != nil {
		l.Critical("Parsing docker version error:", err)
	}

	// Display version infos
	l.Info("Docker API connection OK:")
	l.Info("\tAPIVersion:\t", dockerVersion.APIVersion)
	l.Info("\tArch:\t\t", dockerVersion.Arch)
	l.Info("\tExperimental:\t", dockerVersion.Experimental)
	l.Info("\tGitCommit:\t", dockerVersion.GitCommit)
	l.Info("\tGoVersion:\t", dockerVersion.GoVersion)
	l.Info("\tKernelVersion:\t", dockerVersion.KernelVersion)
	l.Info("\tOs:\t\t", dockerVersion.Os)
	l.Info("\tVersion:\t", dockerVersion.Version)
}

package core

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	dapi "./docker-api"
	dguard "github.com/90TechSAS/libgo-docker-guard"
)

/*
	Handle HTTP requests
*/
func HTTPHandler(w http.ResponseWriter, r *http.Request) {
	var URIItems []string // URI
	var returnStr string

	l.Debug("Request URI:", r.RequestURI)

	// Split URI
	URIItems = strings.Split(r.RequestURI, "/")
	if len(URIItems) < 2 {
		l.Error("Error: len(URIItems)<2")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Check first URI item
	switch URIItems[1] {
	case "info":
		fmt.Fprint(w, "/info TODO")
	case "version":
		fmt.Fprint(w, "/version TODO")
	case "list": // Get container list with basic info
		var tmpContainerList []dguard.Container // Temporary DGC Container list

		// Refresh Container list
		RefreshContainerList()

		// Browse container list
		for _, container := range ContainerList {
			var tmpDAPIContainer dapi.Container  // Temporary Docker API Container
			var tmpDGSContainer dguard.Container // Temporary DGS Container
			var status int                       // HTTP status returned
			var body string                      // HTTP body returned
			var err error                        // Error handling

			// Get container info
			status, body = HTTPReq("/containers/" + container.ID + "/json")
			if status != 200 {
				l.Error("Can't get docker container (", container.ID, ") info, status:", status)
			}

			// Parse returned json
			err = json.Unmarshal([]byte(body), &tmpDAPIContainer)
			if err != nil {
				l.Error("Parsing docker container (", container.ID, ") info:", err)
			}

			// Set tmpDGSContainer
			tmpDGSContainer.ID = tmpDAPIContainer.ID
			tmpDGSContainer.Hostname = tmpDAPIContainer.Config.Hostname
			tmpDGSContainer.Image = tmpDAPIContainer.Image
			tmpDGSContainer.IPAddress = tmpDAPIContainer.NetworkSettings.IPAddress
			tmpDGSContainer.MacAddress = tmpDAPIContainer.NetworkSettings.MacAddress
			tmpDGSContainer.State.Dead = tmpDAPIContainer.State.Dead
			tmpDGSContainer.State.Error = tmpDAPIContainer.State.Error
			tmpDGSContainer.State.ExitCode = tmpDAPIContainer.State.ExitCode
			tmpDGSContainer.State.FinishedAt = tmpDAPIContainer.State.FinishedAt
			tmpDGSContainer.State.OOMKilled = tmpDAPIContainer.State.OOMKilled
			tmpDGSContainer.State.Paused = tmpDAPIContainer.State.Paused
			tmpDGSContainer.State.Pid = tmpDAPIContainer.State.Pid
			tmpDGSContainer.State.Restarting = tmpDAPIContainer.State.Restarting
			tmpDGSContainer.State.Running = tmpDAPIContainer.State.Running
			tmpDGSContainer.State.StartedAt = tmpDAPIContainer.State.StartedAt

			// Add tmpDGSContainer to tmpContainerList
			tmpContainerList = append(tmpContainerList, tmpDGSContainer)
		}

		// tmpContainerList => json
		tmpJson, _ := json.Marshal(tmpContainerList)

		// Add json to the returned string
		returnStr = string(tmpJson)

	default:
		fmt.Fprint(w, "default TODO")
	}

	fmt.Fprint(w, returnStr)
}

/*
	Run HTTP Server
*/
func RunHTTPServer() {
	http.HandleFunc("/", HTTPHandler)
	http.ListenAndServe(DGConfig.DockerGuard.ListenInterface+":"+DGConfig.DockerGuard.ListenPort, nil)
}

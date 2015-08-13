package core

import (
	"fmt"
	"net/http"
	"strings"
)

/*
	Handle HTTP requests
*/
func HTTPHandler(w http.ResponseWriter, r *http.Request) {
	var URIItems []string // URI

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
	default:
		fmt.Fprint(w, "default TODO")
	}

}

/*
	Run HTTP Server
*/
func RunHTTPServer() {
	http.HandleFunc("/", HTTPHandler)
	http.ListenAndServe(DGConfig.DockerGuard.ListenInterface+":"+DGConfig.DockerGuard.ListenPort, nil)
}

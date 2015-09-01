package core

import (
	"net/http"

	"github.com/gorilla/mux"
)

/*
	Log HTTP requests' URI
*/
func HTTPURILogger(r *http.Request, rm *mux.RouteMatch) bool {
	l.Verbose("Request URI:", r.RequestURI)
	return true
}

/*
	Secure API access
*/
func HTTPSecureAPI(r *http.Request, rm *mux.RouteMatch) bool {
	auth, ok := r.Header["Auth"]

	if ok && len(auth) == 1 && auth[0] == DGConfig.DockerGuard.APIPassword {
		l.Debug("Auth OK from", r.RemoteAddr)
		return true
	}

	l.Warn("Failed auth from", r.RemoteAddr)
	return false
}

/*
	Run HTTP Server
*/
func RunHTTPServer() {
	r := mux.NewRouter()
	r1 := r.MatcherFunc(HTTPURILogger).MatcherFunc(HTTPSecureAPI).Subrouter()
	rGET := r1.Methods("GET").Subrouter()

	rGET.HandleFunc("/dockerinfos", HTTPHandlerDockerinfos)
	rGET.HandleFunc("/probeinfos", HTTPHandlerProbeinfos)
	rGET.HandleFunc("/list", HTTPHandlerList)
	http.Handle("/", r)

	http.ListenAndServe(DGConfig.DockerGuard.ListenInterface+":"+DGConfig.DockerGuard.ListenPort, r)
}

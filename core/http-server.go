package core

import (
	"net/http"

	"github.com/gorilla/mux"
)

/*
	Handle HTTP requests
*/
func HTTPMatcher(r *http.Request, rm *mux.RouteMatch) bool {
	l.Verbose("Request URI:", r.RequestURI)
	return true
}

/*
	Run HTTP Server
*/
func RunHTTPServer() {
	r := mux.NewRouter().MatcherFunc(HTTPMatcher).Subrouter()
	r_GET := r.Methods("GET").Subrouter()

	r_GET.HandleFunc("/info", HTTPHandlerInfo)
	r_GET.HandleFunc("/list", HTTPHandlerList)
	http.Handle("/", r)

	http.ListenAndServe(DGConfig.DockerGuard.ListenInterface+":"+DGConfig.DockerGuard.ListenPort, r)
}

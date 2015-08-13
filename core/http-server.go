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
	r0 := mux.NewRouter()
	r1 := r0.Methods("GET").Subrouter()
	r2 := r1.MatcherFunc(HTTPMatcher).Subrouter()

	r2.HandleFunc("/info", HTTPHandlerInfo)
	r2.HandleFunc("/list", HTTPHandlerList)
	http.Handle("/", r0)

	http.ListenAndServe(DGConfig.DockerGuard.ListenInterface+":"+DGConfig.DockerGuard.ListenPort, r0)
}

package core

import (
	"io/ioutil"
	"math/rand"
	"net"
	"net/http"
)

var (
	// HTTP stuff
	tr     *http.Transport
	client *http.Client
)

// Custom dial for Docker unix socket
func unixSocketDial(proto, addr string) (conn net.Conn, err error) {
	return net.Dial("unix", DGConfig.Docker.UnixSocketPath)
}

/*
	Initialize API Client
*/
func InitAPIClient() {
	tr = &http.Transport{
		Dial: unixSocketDial,
	}
	client = &http.Client{Transport: tr}

}

/*
	Do HTTP request on API
*/
func HTTPReq(path string) (int, string) {
	var resp *http.Response               // Docker API response
	var body []byte                       // Docker API response body
	var err error                         // Error handling
	var reqID int = rand.Intn(2000000000) // Request ID (for debugging)

	// HTTP Get request on the docker unix socket
	l.Silly("(", reqID, ") Get API:", path)
	resp, err = client.Get("http://docker" + path)
	if err != nil {
		l.Error("Error: http request:", err)
		return 400, ""
	}

	// Read the body
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		l.Error("(", reqID, ") Error: http response body:", err)
		return 400, ""
	}

	l.Silly("(", reqID, ")Docker API response body:", "\n"+string(body))

	// Return HTTP status code + body
	return resp.StatusCode, string(body)
}

package core

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/*
	Handle GET /list
*/
func HTTPHandlerList(w http.ResponseWriter, r *http.Request) {
	var returnStr string

	// tmpContainerList => json
	tmpJSON, _ := json.Marshal(ContainerList)

	// Add json to the returned string
	returnStr = string(tmpJSON)

	fmt.Fprint(w, returnStr)
}

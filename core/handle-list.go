package core

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func HTTPHandlerList(w http.ResponseWriter, r *http.Request) {
	var returnStr string

	// tmpContainerList => json
	tmpJson, _ := json.Marshal(ContainerList)

	// Add json to the returned string
	returnStr = string(tmpJson)

	fmt.Fprint(w, returnStr)
}

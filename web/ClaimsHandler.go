package web

import (
	"claims-service/fileprocessor"
	"claims-service/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Accepts a file path to be processed
// Expected the pod's to have the file storage mounted to be consumed
// ToDo - Extend to work with Object Store
func ClaimsFileHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	var rbody models.RequestObj
	err = json.Unmarshal(body, &rbody)
	if err != nil {
		fmt.Fprintf(w, "Failed to read file path from request body.\n")
		return
	}

	// Call claimes file processor and read contents
	fileprocessor.ReadAndPushClaims(rbody.ReqBody)

	fmt.Fprintf(w, "File Processed\n")
}

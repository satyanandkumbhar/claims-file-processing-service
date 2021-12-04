package main

import (
	"claims-service/web"
	"fmt"
	"log"
	"net/http"
)

func main() {

	// Handle claims file processing request from Azure Function
	http.HandleFunc("/", web.ClaimsFileHandler)
	fmt.Println("Server started at port 8089")
	log.Fatal(http.ListenAndServe(":8089", nil))
}

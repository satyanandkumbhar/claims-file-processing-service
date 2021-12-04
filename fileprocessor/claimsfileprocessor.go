package fileprocessor

import (
	"claims-service/kafkaprocessor"
	"claims-service/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

//Read the JSON format claims and push the claims to Kafka topic
func ReadAndPushClaims(filepath string) {
	log.Println(filepath)

	// Open our jsonFile
	jsonFile, err := os.Open(filepath)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	//fmt.Println("Successfully Opened", filepath)
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var claims models.ClaimRequests

	// unmarshal claims data
	json.Unmarshal(byteValue, &claims)

	//Send the claims to Kafka topic
	kafkaprocessor.ProcessClaims(&claims)
}

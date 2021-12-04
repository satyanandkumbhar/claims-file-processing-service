package kafkaprocessor

import (
	"claims-service/models"
	"context"
	"encoding/json"
	"fmt"
	"time"

	kafka "github.com/segmentio/kafka-go"
)

// Get's the struct of claims and sends them to Kafka topic
// Init's the Kafka connection and closes by end of processing
func ProcessClaims(claims *models.ClaimRequests) {
	//Get the writer
	writer := getKafkaWriter()
	//process indivuals writes
	for i := 0; i < len(claims.ClaimRequests); i++ {
		serialisedBytes, err := json.Marshal(claims.ClaimRequests[i])
		serialisedClaimValue := string(serialisedBytes)
		if err != nil {
			panic("Failed to serialise current claime")
		}
		sendClaim(&claims.ClaimRequests[i], serialisedClaimValue, writer)
	}
	//close connection
	closeKafkaConnection(writer)
}

func sendClaim(claim *models.Claim, claimData string, writer *kafka.Writer) {
	fmt.Println("claim record", claimData)

	key := fmt.Sprintf("Key-%d", claim.Id)
	msg := kafka.Message{
		Key:   []byte(key),
		Value: []byte(claimData),
	}
	err := writer.WriteMessages(context.Background(), msg)
	if err != nil {
		fmt.Println(err)
	}
	time.Sleep(1 * time.Second)

}

func getKafkaWriter() *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP("localhost:9092"),
		Topic:    "claims",
		Balancer: &kafka.LeastBytes{},
	}
}

func closeKafkaConnection(kafkaConnection *kafka.Writer) {
	kafkaConnection.Close()
}

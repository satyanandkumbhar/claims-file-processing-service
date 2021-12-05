# claims-file-processing-service
Service responsible to read the json formatted claims
Process the data and send the claims to kafka topic to be consumed by downstream systems

## Prerequisites
the code currently expects
* Kafka running at localhost on post 9092
* Having a topic 'claims'

## How to run
* Run the main.go, this will start a http server on port 8089
* Make a request to the only path available at '/'
* Sample request 
> curl -H "Content-Type: application/json" --data '{"reqBody":"c:\\temp\\claims.json"}' http://localhost:8089/

## Further enhancements
* Add unit tests
* Make the Kafka URL and topic as application arguments
* Add functioality to read file from Object Store
* Better error handling
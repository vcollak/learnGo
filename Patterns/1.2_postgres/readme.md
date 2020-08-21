
## Build the image

To build the image, execute: `docker build -t eg_postgresql .`

## Run the container
Now run the container: `docker run --rm -P --name pg_test eg_postgresql`

## Find the docker TCP port. 
The docker image assigns a port it exposes. to find the port, execute `docker ps`. Find the TCP port and configure models/uses.go port variable with it. 

## Run the program 
`go run main.go`
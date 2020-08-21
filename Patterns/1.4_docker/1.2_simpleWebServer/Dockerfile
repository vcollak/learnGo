
FROM golang:latest

#maintainer info
LABEL maintainer="Bobbyimir Smith <Bobby@smith.net>"

#copy everyting from working directory to docker container 
WORKDIR $GOPATH/src/learnGo/docker/1.2_docker
COPY . .

#get dependencies
RUN go get -d -v ./...

#build the app
RUN GOOS=linux go build -a -o /go/bin/testServer .

CMD ["/go/bin/testServer"] 
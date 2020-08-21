

#############################
# STAGE 1 - build the GO app
#############################

FROM golang:latest as builder

#maintainer info
LABEL maintainer="Bobbyimir Peters <Bobby@peters.net>"

#copy everyting from working directory to docker container 
WORKDIR $GOPATH/src/learnGo/docker
COPY . .

#get dependencies
RUN go get -d -v ./...

#build the app
RUN GOOS=linux go build -a -o /go/bin/dockerTest .

###############################################
# STAGE 2 - smaller image 
###############################################

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /go/bin/dockerTest .

CMD ["./dockerTest"] 
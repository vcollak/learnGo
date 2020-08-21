# Various useful Docker commands

### To build the docker image run: 
`docker build -t dockertest . `  

### To show all docker containers: 
`docker image ls`

### To show all containers:
`docker container ls`

### To run the container called dockertest : 
`docker run dockertest`

### To execute a "pwd" command inside of a running docker: 
`docker exec pwd`

### delete all docker images and force any that give errors
`docker rmi $(docker images -a -q) -f`

### Help with other commands:
https://gist.github.com/bradtraversy/89fad226dc058a41b596d586022a9bd3
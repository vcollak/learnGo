

This will create a much smaller image because we're first building the app
and then creating alpine based container

### Build the container
`docker image build -t testserver .`

### Start the container
Run in a foreground: `docker container run -it -p 8080:8080 testserver`

Run in a background: `docker container run -d -p 8080:8080 testserver`
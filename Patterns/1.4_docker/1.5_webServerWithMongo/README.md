

# Background
This is based on [How to Keep Your Development Environment Clean](https://outcrawl.com/clean-development-environment-docker/) article.
It starts mongoDB as well as golang based web server to which users can POST data and read data. Data is persisted in the Mongo DB


# Start the containers
docker-compose up

# Add some data
curl -i -X POST -H "Content-Type: application/json" -d "{\"text\":\"hello there\"}" http://localhost:8080/posts

# Check data
http://localhost:8080/posts

# Connect to DB
You can connect to the Mongo DB on the port 27100
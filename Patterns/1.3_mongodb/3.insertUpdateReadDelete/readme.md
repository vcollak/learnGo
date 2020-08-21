To start the mongo DB container, use the `docker-compose up`
After the container us running, you can run the program using `go run main.go`


Instead of using mgo, this uses a new MongoDB supported driver `go.mongodb.org/mongo-driver/mongo`



Code from: https://github.com/ruanbekker/code-examples/blob/master/mongodb/golang/examples.go
Based on article: https://blog.ruanbekker.com/blog/2019/04/17/mongodb-examples-with-golang/
More Mongo info: https://github.com/mongodb/mongo-go-driver
More on the new driver: https://www.compose.com/articles/mongodb-and-go-moving-on-from-mgo/
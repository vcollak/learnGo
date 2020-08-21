# REST based server using Gorilla Mux

Execute: `go run main.go`


## Testing

### Index Page (home)

`curl --location --request GET localhost:8080`


### Accessing /todos route

`curl --location --request GET localhost:8080/todos`


### /todos with a parameter

`curl --location --request GET localhost:8080/todos/1`




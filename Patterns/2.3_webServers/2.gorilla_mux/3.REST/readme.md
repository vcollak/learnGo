# REST based server using Gorilla Mux

Execute: `go run main.go`


## Testing

### Index Page (home)

`curl --location --request GET localhost:9000/`


### Get user information by ID (1)

`curl --location --request GET localhost:9000/user/1`


### Create user

`curl --location --request POST localhost:9000/user`

### Update user
`curl --location --request PUT localhost:9000/user/1`

### Delete User
`curl --location --request DELETE localhost:9000/user/1`



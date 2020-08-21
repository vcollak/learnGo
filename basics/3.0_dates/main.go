package main

import (
	"fmt"
	"time"
)

//Token holds the expiration
type Token struct {
	ExpireDate time.Time
}

func isExpired(expireDate time.Time) bool {

	if time.Now().After(expireDate) {
		return true
	}

	return false

}

func main() {

	//expired
	expiresIn := 10

	//not expired
	//expiresIn := 20

	expireDate := time.Now().Add(time.Second * time.Duration(expiresIn))

	t := Token{
		ExpireDate: expireDate,
	}

	fmt.Printf("Waiting %d seconds\n", expiresIn)
	time.Sleep(time.Second * 15)

	if isExpired(t.ExpireDate) {
		fmt.Println("Expired")
	} else {
		fmt.Println("Not Expired")
	}

}

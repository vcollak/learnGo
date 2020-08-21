package main

import (
	"fmt"
	"learnGo/basics/2.6_errors/9.errorPackage/apperrors"
)

func main() {

	err := apperrors.New("error message", "101")
	if err != nil {
		fmt.Println(err)
	}

}

package main

import (
	"fmt"

	"github.com/pkg/errors"
)

var ErrGetAge = errors.New("GetAge: unable to get age")

func GetAge() error {
	return errors.Wrap(ErrGetAge, "Get Age failed")
}

func GetName() error {

	if err := GetAge(); err != nil {
		return errors.Wrap(err, "Unable to get Name")
	}

	return nil

}

func main() {

	if err := GetName(); err != nil {

		//error
		fmt.Println(err)

		//error cause (the first thing that resulted in error)
		fmt.Println(errors.Cause(err))

		if errors.Cause(err) == ErrGetAge {

			//entire error
			fmt.Printf("Get Age problem, %v\n", err)

			//stack trace
			fmt.Printf("%+v\n", err)
		}

	}

}

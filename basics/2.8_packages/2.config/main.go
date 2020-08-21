package main

import (
	"LearnGo/basics/2.8_packages/2.config/config"
	"LearnGo/basics/2.8_packages/2.config/users"
	"fmt"
)

func main() {

	//////////////////////////////////
	// Configuration
	//////////////////////////////////

	//get using simple variables in the config package
	path := config.Path
	version := config.Version
	host := config.Hostname
	port := config.Port

	fmt.Println("Configured Variables:")
	fmt.Println(path)
	fmt.Println(version)
	fmt.Println(host)
	fmt.Println(port)

	//getting the System struct from the config
	//using this approach we can have multiple config
	//types in one package
	fmt.Println()
	fmt.Println("Struct config")
	c := config.New()
	fmt.Println(c.Path)
	fmt.Println(c.Version)
	fmt.Println(c.Hostname)
	fmt.Println(c.Port)

	//////////////////////////////////
	// New User object
	//////////////////////////////////

	fmt.Println()
	u := users.New()
	u.FullName = "Bobbyimir Peters"
	u.Username = "vpeters"
	u.Password = "somePassword"
	fmt.Println(u)

	u1 := users.New()
	u1.FullName = "Liz Peters"
	u1.Username = "apeters"
	u1.Password = "somePassword1"
	fmt.Println(u1)
}

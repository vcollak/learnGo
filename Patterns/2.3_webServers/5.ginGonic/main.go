/*
Serve static content with Gin alongside of having handlers
*/

package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	//define static route
	router.Static("/static", "./static")

	//redirect anything going / to static
	router.GET("/", func(c *gin.Context) {
		c.Request.URL.Path = "/static"
		router.HandleContext(c)
	})

	//ping route with a json response
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	//hello world handler
	router.GET("/helloWorld", helloWorldHandler)

	// Listen and serve on 0.0.0.0:8080
	router.Run(":8080")
}


//Hello World handler
func helloWorldHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World!",
	})
}

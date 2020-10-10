package main

// Gin is a web framework written in Golang: https://github.com/gin-gonic/gin
// Gin doc: https://gin-gonic.com/docs/
// Go Gin example: https://github.com/eddycjy/go-gin-example
// Go Gin real world example: https://github.com/gothinkster/golang-gin-realworld-example-app

import (
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
			"now":     time.Now(),
			"a":       "a",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}

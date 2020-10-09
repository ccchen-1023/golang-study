package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

type employee struct {
	// filed name must start with a capital letter
	Name     string `json:"name"`
	EmpNo    int    `json:"emp_no"`
	Birthday int64  `json:"birthday"`
}

func main() {
	engine := gin.Default()
	engine.POST("/json", handleRequest)
	engine.Run()

	/*
		curl --location --request POST 'http://localhost:8080/json' \
		--header 'Content-Type: application/json;charset=utf8' \
		--data-raw '{
		    "name": "sunnyboy2",
		    "emp_no": 123,
		    "birthday": 1602236765
		}'
	*/
}

func handleRequest(ctx *gin.Context) {

	// If `GET`, only `Form` binding engine (`query`) used.
	// If `POST`, first checks the `content-type` for `JSON` or `XML`, then uses `Form` (`form-data`).
	// See more at https://github.com/gin-gonic/gin/blob/master/binding/binding.go#L48

	var emp employee
	if ctx.ShouldBind(&emp) == nil {
		log.Println(emp)
	}
	save(emp)
	ctx.JSON(200, emp)
}

func save(emp employee) {
	// save to ddb
}

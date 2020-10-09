package main

import "github.com/gin-gonic/gin"

func main() {

	engine := gin.Default()

	engine.GET("/form", formStruct)

	engine.Run()

	//http://localhost:8080/form?a=hello&b=world
}

func formStruct(ctx *gin.Context) {
	var b structB
	ctx.Bind(&b)
	ctx.JSON(200, b)
}

type structA struct {
	FieldA string `form:"a"`
}

type structB struct {
	FieldA structA
	FieldB string `form:"b"`
}

package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	engine := gin.Default()
	engine.GET("json", json)
	engine.Run(":5566")

	// http://localhost:5566/json
}

func json(ctx *gin.Context) {
	data := map[string]interface{}{
		"language": "Go 語言",
		"tag":      "<br/>",
	}
	ctx.JSON(200, data)
	//c.AsciiJSON(http.StatusOK, data)

}

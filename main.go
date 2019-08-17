package main

import (
	"StudyGin/app"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		app.RespSuccess(c, nil	)
	})
	r.Run(":3003") // listen and serve on 0.0.0.0:8080
}
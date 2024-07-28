package main

import (
	"suspectRecall/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Static("/public", "./public")
	r.NoRoute(func(ctx *gin.Context) {
		ctx.File("./public/index.html")
	})
	router.InitializeRoutes(r)

	r.Run(":3001")
}

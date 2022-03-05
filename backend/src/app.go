package main

// TODO: convert this into an API, define endpoints

import (
	"backend/src/controller"
	dao "backend/src/database"
	"github.com/gin-gonic/gin"
)

func main() {
	dao.SetupDb()
	router := gin.Default()
	router.GET("/", controller.LandingPage)
	router.GET("/product/:id", controller.GetProduct)
	router.POST("/product", controller.FilteredProducts)
	err := router.Run(":5001")
	if err != nil {
		return
	}
}

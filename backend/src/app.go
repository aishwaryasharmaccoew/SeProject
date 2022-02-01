package main

// TODO: convert this into an API, define endpoints

import (
	"backend/src/controller"
	dao "backend/src/database"
	"github.com/gin-gonic/gin"
)

//func main() {
//	log.Println("Inside the app")
//	dao.SetupDb()
//
//	var query = []string{"screw", "large", "scissor"}
//	controller.FilteredProducts(query)
//}

func main() {
	dao.SetupDb()
	router := gin.Default()
	router.POST("/filter", controller.FilteredProducts)
	err := router.Run(":5001")
	if err != nil {
		return
	}
}

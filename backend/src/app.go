package main

// TODO: convert this into an API, define endpoints

import (
	"backend/src/controller"
	dao "backend/src/database"
	"github.com/gin-gonic/gin"
	"log"
)

//func main() {
//	log.Println("Inside the app")
//	dao.SetupDb()
//
//	var query = []string{"screw", "large", "scissor"}
//	controller.FilteredProducts(query)
//}

type FilterQuery struct {
	Filters []string `json:"filters,omitempty"`
}

func main() {
	dao.SetupDb()
	router := gin.Default()
	router.POST("/filter", func(c *gin.Context) {
		var filterQuery FilterQuery
		err := c.BindJSON(&filterQuery)
		log.Println(filterQuery)
		if err != nil {
			c.JSON(404, "Error while parsing the input JSON")
			log.Println(err.Error())
		}
		c.JSON(200, controller.FilteredProducts(filterQuery.Filters)) // Your custom response here
	})

	err := router.Run(":5001")
	if err != nil {
		log.Fatalln(err.Error())
	}
}

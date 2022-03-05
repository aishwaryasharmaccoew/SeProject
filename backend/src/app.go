package main

// TODO: convert this into an API, define endpoints

import (
	"backend/src/controller"
	dao "backend/src/database"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/handlers"
)

func main() {
	dao.SetupDb()
	router := gin.Default()
	credentials := handlers.AllowCredentials()
	methods := handlers.AllowedMethods([]string{"POST", "GET", "OPTIONS"})

	origins := handlers.AllowedOrigins([]string{"http://localhost:5001", "http://localhost:4200"})

	router.GET("/", controller.LandingPage)
	router.GET("/product/:id", controller.GetProduct)
	router.POST("/product/:id", controller.FilteredProducts)
	err := router.Run(":5001")
	if err != nil {
		return
	}
	log.Fatal(http.ListenAndServe(":5001", handlers.CORS(credentials, methods, origins)(router)))
}

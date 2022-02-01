package main

import (
	controller "backend/src/controller"
	"backend/src/database"
	"log"
)

// TODO: convert this into an API, define endpoints

func main() {
	log.Println("Inside the app")
	dao.SetupDb()

	var query = []string{"screw", "large", "scissor"}
	controller.FilteredProducts(query)
}

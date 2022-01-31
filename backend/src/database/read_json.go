package main

import (
	"backend/src/database/models"
	"bufio"
	"encoding/json"
	"log"
	"os"
)

func check(e error) {
	if e != nil {
		log.Fatalln(e.Error())
	}
}

func Read(path string) []models.Product {
	dat, err := os.Open(path)
	check(err)
	defer dat.Close()
	log.Println("File opened for reading")

	var product_list []models.Product

	scanner := bufio.NewScanner(dat)
	for scanner.Scan() {
		var product models.Product
		err = json.Unmarshal(scanner.Bytes(), &product)
		check(err)
		//log.Println(product.Name)
		product_list = append(product_list, product)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err.Error())
	}

	return product_list
}

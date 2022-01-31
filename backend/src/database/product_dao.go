package main

import (
	"backend/src/database/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"strings"
)

func main() {
	var file_path = "/Users/nimishbajaj/GolandProjects/SeProject/backend/src/database/files/vict_products_detailed.jsonl"
	var product_list = Read(file_path)
	//log.Println(product_list)

	db, err := gorm.Open(sqlite.Open("findmyknife.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Error opening the database")
	}

	err2 := db.AutoMigrate(&models.ProductSQL{})
	if err2 != nil {
		log.Fatal(err2.Error())
	}

	addData(db, product_list)

	printData(db)
}

func listToString(items []string) string {
	return strings.Join(items, ",")
}

func toSql(product models.Product) models.ProductSQL {
	var productsql models.ProductSQL
	productsql.Id = product.Id
	productsql.Name = product.Name
	productsql.Price = product.Price
	productsql.Category = product.Category
	productsql.ImageUrl = product.ImageUrl
	productsql.Product_url = product.Product_url

	productsql.Colors = listToString(product.Colors)
	productsql.Key_features = listToString(product.Key_features)
	productsql.Tools = listToString(product.Tools)

	return productsql
}

func addData(db *gorm.DB, product_list []models.Product) {
	for _, product := range product_list {
		var productSql = toSql(product)
		db.Create(productSql)
	}
}

func printData(db *gorm.DB) {
	var productSQL models.ProductSQL
	rows, err := db.Model(&models.ProductSQL{}).Rows()
	if err != nil {
		log.Fatal(err.Error())
	}

	for rows.Next() {
		err := db.ScanRows(rows, &productSQL)
		if err != nil {
			return
		}
		log.Println(productSQL)
	}
}

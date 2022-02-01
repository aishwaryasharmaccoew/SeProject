package dao

import (
	"backend/src/database/models"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"os"
	"strings"
)

var DB *gorm.DB

func connectDB() {
	db, err := gorm.Open(sqlite.Open("findmyknife.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Error opening the database")
	}
	DB = db
}

// TODO: change the path to relative path
func insertRowsFromJSON() {
	var file_path = "/Users/nimishbajaj/GolandProjects/SeProject/backend/src/database/files/vict_products_detailed.jsonl"
	var product_list = Read(file_path)
	err2 := DB.AutoMigrate(&models.ProductSQL{})
	if err2 != nil {
		log.Fatal(err2.Error())
	}
	addData(DB, product_list)
	//printData(DB)
}

func SetupDb() {
	db_file_path := "findmyknife.db"
	_, err := os.OpenFile(db_file_path, os.O_RDWR, 0644)
	if err == nil {
		e := os.Remove(db_file_path)
		check(e)
	}
	connectDB()
	insertRowsFromJSON()
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

// TODO: implement pagination
func QueryTable(filters []string) []models.ProductSQL {
	var productsSql []models.ProductSQL
	tx := DB.Where("Tools LIKE ?", fmt.Sprintf("%%%s%%", filters[0])).Find(&productsSql)
	for i := 1; i < len(filters); i++ {
		var filter = filters[i]
		tx.Where("Tools LIKE ?", fmt.Sprintf("%%%s%%", filter)).Find(&productsSql)
	}
	log.Println(fmt.Sprintf("Number of results obtained %d", len(productsSql)))
	return productsSql
}

package dao

import (
	"backend/src/database/models"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func connectDB() {
	db, err := gorm.Open(sqlite.Open("findmyknife.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Error opening the database")
	}
	DB = db
}

func insertRowsFromJSON() {
	file_path := filepath.Join(".", "database/files/vict_products_detailed.jsonl")
	var product_list = Read(file_path)
	err2 := DB.AutoMigrate(&models.ProductSQL{})
	if err2 != nil {
		log.Fatal(err2.Error())
	}
	addData(DB, product_list)

	err3 := DB.AutoMigrate(&models.Customer{})
	if err3 != nil {
		log.Fatal(err3.Error())
	}

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

func AllQueryTable(filters []string) []models.ProductSQL {
	var productsSql []models.ProductSQL
	tx := DB.Find(&productsSql)
	for i := 0; i < len(filters); i++ {
		var filter = filters[i]
		tx.Where("Tools LIKE ?", fmt.Sprintf("%%%s%%", filter)).Find(&productsSql)
	}
	log.Println(fmt.Sprintf("Number of results obtained %d", len(productsSql)))
	return productsSql
}
func QueryTable(filters []string, numResults int, pageNum int) []models.ProductSQL {
	var productsSql []models.ProductSQL
	tx := DB.Find(&productsSql)
	for i := 0; i < len(filters); i++ {
		var filter = filters[i]
		tx.Where("Tools LIKE ?", fmt.Sprintf("%%%s%%", filter)).Find(&productsSql)
	}
	tx.Limit(numResults).Offset(numResults * pageNum).Find(&productsSql)
	log.Println(fmt.Sprintf("Number of results obtained %d", len(productsSql)))
	return productsSql
}

func GetProduct(id string) models.ProductSQL {
	var productSQL models.ProductSQL
	DB.Where("Id = ?", id).First(&productSQL)
	return productSQL
}

func GetUser(uname string, pass string) models.Customer {
	var user models.Customer
	err := DB.Where("user_name = ? AND password = ?", uname, pass).First(&user)
	if err != nil {
		log.Println(err)
	}
	log.Println(fmt.Sprintf("Got user %s", user))
	return user
}

func PersistUser(user models.Customer) {
	DB.Create(user)
	log.Println(fmt.Sprintf("%s successfully added", user.UserName))
}

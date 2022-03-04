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
	file_path, _ := filepath.Abs("../src/database/files/vict_products_detailed.jsonl")
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
	tx := DB.Find(&productsSql)
	for i := 0; i < len(filters); i++ {
		var filter = filters[i]
		tx.Where("Tools LIKE ?", fmt.Sprintf("%%%s%%", filter)).Find(&productsSql)
	}
	log.Println(fmt.Sprintf("Number of results obtained %d", len(productsSql)))
	return productsSql
}
func PaginatedQueryTable(filters []string, pageID string) *models.ProductList {
	// hardcoding pageId to check API
	//pageID := "B-0.8413.M3"
	pageSize := 5
	productsSql := &models.ProductList{}
	tx := DB.Find(&productsSql.Items)
	for i := 0; i < len(filters); i++ {
		var filter = filters[i]
		tx = tx.Where("Tools LIKE ?", fmt.Sprintf("%%%s%%", filter)).Find(&productsSql.Items)
	}
	tx.Where("Id >= ?", pageID).Order("Id").Limit(pageSize + 1).Find(&productsSql.Items)
	if len(productsSql.Items) == pageSize+1 {
		productsSql.NextPageID = productsSql.Items[len(productsSql.Items)-1].Id
		productsSql.Items = productsSql.Items[:pageSize] // this shortens the slice by 1
	}
	log.Println(fmt.Sprintf("Number of results obtained %d", len(productsSql.Items)))
	log.Println(fmt.Sprintf("Id obtained %s", pageID))
	return productsSql
}

func GetProduct(id string) models.ProductSQL {
	var productSQL models.ProductSQL
	DB.Where("Id = ?", id).First(&productSQL)
	return productSQL
}
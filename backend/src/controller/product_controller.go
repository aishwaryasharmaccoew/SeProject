package controller

import (
	dao "backend/src/database"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

type FilterQuery struct {
	Filters []string `json:"filters,omitempty"`
}

func FilteredProducts(c *gin.Context) {
	var filterQuery FilterQuery
	err := c.BindJSON(&filterQuery)
	pageid := c.Param("id")

	log.Println(filterQuery)
	if err != nil {
		c.JSON(404, "Error while parsing the input JSON")
		log.Println(err.Error())
	}
	query_result := dao.PaginatedQueryTable(filterQuery.Filters, pageid)
	c.JSON(200, query_result)
}

func LandingPage(c *gin.Context) {
	query_result := dao.QueryTable([]string{})
	c.JSON(200, query_result)
}

func GetProduct(c *gin.Context) {
	id := c.Param("id")
	log.Println(fmt.Sprintf("Id %s", id))
	query_result := dao.GetProduct(id)
	c.JSON(200, query_result)
}

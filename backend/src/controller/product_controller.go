package controller

import (
	dao "backend/src/database"
	"fmt"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FilterQuery struct {
	Filters []string `json:"filters,omitempty"`
}

func getPaginationInfo(c *gin.Context) (int, int) {
	numResults, err := strconv.Atoi(c.Query("numResults"))
	if err != nil {
		log.Println(err.Error())
		return -1, -1
	}

	pageNum, err := strconv.Atoi(c.Query("pageNum"))
	if err != nil {
		log.Println(err.Error())
		return -1, -1
	}

	return numResults, pageNum
}

func FilteredProducts(c *gin.Context) {
	var filterQuery FilterQuery
	err := c.BindJSON(&filterQuery)
	if err != nil {
		log.Println(err.Error())
		c.JSON(404, "Error while parsing the input JSON")
		return
	}

	numResults, pageNum := getPaginationInfo(c)

	if numResults == -1 {
		c.JSON(404, "Error while parsing the input JSON")
		return
	}

	log.Println(filterQuery)
	query_result := dao.QueryTable(filterQuery.Filters, numResults, pageNum)
	c.JSON(200, query_result)
}

func LandingPage(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")

	numResults, pageNum := getPaginationInfo(c)
	if numResults == -1 {
		c.JSON(404, "Error while parsing the input JSON")
		return
	}

	query_result := dao.QueryTable([]string{}, numResults, pageNum)
	c.JSON(200, query_result)
}

func GetProduct(c *gin.Context) {
	id := c.Param("id")
	log.Println(fmt.Sprintf("Id %s", id))
	query_result := dao.GetProduct(id)
	c.JSON(200, query_result)
}

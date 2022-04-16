package controller

import (
	dao "backend/src/database"
	"backend/src/database/models"
	"fmt"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FilterQuery struct {
	Filters []string `json:"filters,omitempty"`
}

type LoginUser struct {
	UserName string `json:"username"`
	Password string `json:"password"`
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
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
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
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
	id := c.Param("id")
	log.Println(fmt.Sprintf("Id %s", id))
	query_result := dao.GetProduct(id)
	c.JSON(200, query_result)
}

func Login(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")

	var user LoginUser

	if err := c.BindJSON(&user); err != nil {
		c.JSON(404, "Error while parsing the input JSON")
		log.Println(err.Error())
	}
	query_result := dao.GetUser(user.UserName, user.Password)

	if query_result.UserName == "" {
		c.JSON(404, "User not found")
	} else {
		c.JSON(200, query_result)
	}

}

func SignUp(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")

	//jsonFromUI := []byte(`{ "LastName":"sharma","UserName": "abc", "Password":"test","FirstName":"aishwarya","Email":"abc@gmail.com"}`)
	var user models.Customer

	if err := c.BindJSON(&user); err != nil {
		c.JSON(404, "Error while parsing the input JSON")
		log.Println(err.Error())
	}

	fmt.Println(user.Email)

	dao.PersistUser(user)
	c.JSON(200, "User succesfully Created!")

}

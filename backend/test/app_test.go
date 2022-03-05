package test

import (
	dao "backend/src/database"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDataBaseSetup(t *testing.T) {
	dao.SetupDb()
	req, _ := http.NewRequest("GET", "/product", nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	router := gin.Default()
	router.ServeHTTP(rr, req)

	return rr
}

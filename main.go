package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/people", getPeople)

	router.Run("localhost:8080")
}

type person struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Credit int    `json:"credit"`
}

var people = []person{
	{ID: "1", Name: "Suzie", Email: "***REMOVED***", Credit: 0},
	{ID: "2", Name: "Camille", Email: "***REMOVED***", Credit: 0},
	{ID: "3", Name: "Mateusz", Email: "***REMOVED***", Credit: 0},
	{ID: "4", Name: "Steve", Email: "***REMOVED***", Credit: 0},
	{ID: "5", Name: "Kenny", Email: "***REMOVED***", Credit: 0},
	{ID: "6", Name: "Brian", Email: "***REMOVED***", Credit: 0},
}

func getPeople(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, people)
}

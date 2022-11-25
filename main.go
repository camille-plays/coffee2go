package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/people", getPeople)
	router.POST("/people", postPeople)
	router.GET("/people/:id", getPersonByID)

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

func postPeople(c *gin.Context) {
	var newPerson person

	// Call BindJSON to bind the received JSON to
	// newPerson.
	if err := c.BindJSON(&newPerson); err != nil {
		return
	}

	// Add the new person to the slice.
	people = append(people, newPerson)
	c.IndentedJSON(http.StatusCreated, newPerson)
}

func getPersonByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of people, looking for
	// the person whose ID value matches the parameter.
	for _, a := range people {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "person not found"})
}

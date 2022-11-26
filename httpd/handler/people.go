package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetPeople(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, people)
}

func GetPersonByID(c *gin.Context) {
	id := c.Param("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		return
	}

	// Loop over the list of people, looking for
	// the person whose ID value matches the parameter.
	for _, a := range people {
		if a.ID == intId {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "person not found"})
}

func PostPeople(c *gin.Context) {
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

var people = []person{
	{ID: 1, Name: "Suzie", Email: "***REMOVED***", Credit: 0},
	{ID: 2, Name: "Camille", Email: "***REMOVED***", Credit: 0},
	{ID: 3, Name: "Mateusz", Email: "***REMOVED***", Credit: 0},
	{ID: 4, Name: "Steve", Email: "***REMOVED***", Credit: 0},
	{ID: 5, Name: "Kenny", Email: "***REMOVED***", Credit: 0},
	{ID: 6, Name: "Brian", Email: "***REMOVED***", Credit: 0},
}

type person struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Credit int    `json:"credit"`
}

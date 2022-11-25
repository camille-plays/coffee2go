package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/people", getPeople)
	router.POST("/people", postPeople)
	router.GET("/people/:id", getPersonByID)
	router.PUT("/people/:id/incrementCredit", incrementCredit)
	router.PUT("/people/:id/decrementCredit", decrementCredit)

	router.Run("localhost:8080")
}

type person struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Credit int    `json:"credit"`
}

type incrementCreditCmd struct {
	NbCredit int `json:"credit"`
}

var people = []person{
	{ID: 1, Name: "Suzie", Email: "***REMOVED***", Credit: 0},
	{ID: 2, Name: "Camille", Email: "***REMOVED***", Credit: 0},
	{ID: 3, Name: "Mateusz", Email: "***REMOVED***", Credit: 0},
	{ID: 4, Name: "Steve", Email: "***REMOVED***", Credit: 0},
	{ID: 5, Name: "Kenny", Email: "***REMOVED***", Credit: 0},
	{ID: 6, Name: "Brian", Email: "***REMOVED***", Credit: 0},
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

// we increment the credit of the person refered by an arbitrary number passed in the body
func incrementCredit(c *gin.Context) {
	id := c.Param("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		return
	}
	var credit incrementCreditCmd

	if err := c.BindJSON(&credit); err != nil {
		return
	}

	for _, a := range people {
		if a.ID == intId {
			a.Credit += credit.NbCredit
			people[intId-1] = a
			c.IndentedJSON(http.StatusCreated, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "person not found"})

}

// we decrement the credit of the person refered by 1
func decrementCredit(c *gin.Context) {
	id := c.Param("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		return
	}
	for _, a := range people {
		if a.ID == intId {
			a.Credit -= 1
			people[intId-1] = a
			c.IndentedJSON(http.StatusCreated, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "person not found"})
}

package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type incrementCreditCmd struct {
	NbCredit int `json:"credit"`
}

// we increment the credit of the person refered by an arbitrary number passed in the body
func IncrementCredit(c *gin.Context) {
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
func DecrementCredit(c *gin.Context) {
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

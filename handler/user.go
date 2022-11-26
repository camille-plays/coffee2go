package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/camille-plays/coffee2go/models"
)

func GetUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, models.Users)
}

func GetUserByID(c *gin.Context) {

	// Make sure provided ID is a UUID
	_, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	// TO REMOVE!
	// Loop over the list of people, looking for
	// the person whose ID value matches the parameter.
	for _, a := range models.Users {
		if c.Param("id") == a.ID {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.Status(http.StatusNotFound)
	return
}

func CreateUser(c *gin.Context) {
	var userRequest CreateUserRequest

	// Call BindJSON to bind the received JSON to
	// newPerson.
	if err := c.BindJSON(&userRequest); err != nil {
		return
	}

	user := userFromRequest(userRequest)

	// Add the new person to the slice.
	models.Users = append(models.Users, &user)
	c.IndentedJSON(http.StatusCreated, user)
}

func userFromRequest(r CreateUserRequest) models.User {
	return models.User{
		Email: r.Email,
		Name:  r.Name,
		ID:    uuid.New().String(),
	}
}

package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/camille-plays/coffee2go/dao"
)

func (h Handler) GetUsers(c *gin.Context) {
	users := h.DB.GetUsers()

	c.IndentedJSON(http.StatusOK, users)
}

func (h Handler) GetUserByID(c *gin.Context) {
	// Make sure provided ID is a UUID
	_, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	user := h.DB.GetUser(c.Param("id"))
	if user != nil {
		c.IndentedJSON(http.StatusOK, user)
		return
	}

	c.Status(http.StatusNotFound)
	return
}

func (h Handler) CreateUser(c *gin.Context) {
	var userRequest CreateUserRequest

	if err := c.BindJSON(&userRequest); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	if err := h.validateCreateUserRequest(userRequest); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	user := userFromRequest(userRequest)

	h.DB.CreateUser(&user)
	c.IndentedJSON(http.StatusCreated, user)
}

func userFromRequest(r CreateUserRequest) dao.User {
	return dao.User{
		Email: r.Email,
		Name:  r.Name,
		ID:    uuid.New().String(),
	}
}

func (h Handler) DeleteUser(c *gin.Context) {
	var userRequest DeleteUserRequest

	if err := c.BindJSON(&userRequest); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	if err := h.validateDeleteUserRequest(userRequest); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	h.DB.DeleteUser(h.DB.GetUser(userRequest.ID))
	c.IndentedJSON(http.StatusOK, h.DB.GetUsers())

}

func (h Handler) validateCreateUserRequest(userRequest CreateUserRequest) error {

	if userRequest.Name == "" {
		return fmt.Errorf("Please provide a valid name")
	}

	if userRequest.Email == "" {
		return fmt.Errorf("please provide a valid email")
	}

	for _, u := range h.DB.GetUsers() {
		if userRequest.Name == u.Name {
			return fmt.Errorf("username already exists")
		}
		if userRequest.Email == u.Email {
			return fmt.Errorf("email already exists")
		}
	}

	return nil
}

func (h Handler) validateDeleteUserRequest(userRequest DeleteUserRequest) error {
	var userFound bool
	for _, u := range h.DB.GetUsers() {
		if userRequest.ID == u.ID {
			userFound = true
		}
	}
	if !userFound {
		return fmt.Errorf("user not found")
	}
	return nil
}

package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/camille-plays/coffee2go/dao"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestGetUsers(t *testing.T) {
	router := gin.Default()
	h := NewTestHandler()
	router.GET("/users", h.GetUsers)

	req, _ := http.NewRequest("GET", "/users", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	var users []dao.User

	err := json.NewDecoder(w.Body).Decode(&users)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, w.Code)
	require.Equal(t, users, h.DB.GetUsers())
}

func TestCreateUser(t *testing.T) {
	router := gin.Default()
	h := NewTestHandler()
	router.POST("/user", h.CreateUser)

	jsonBody := []byte(`{"name": "Foo", "email": "Foo@gmail.com"}`)
	bodyReader := bytes.NewReader(jsonBody)

	req, _ := http.NewRequest("POST", "/user", bodyReader)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	var newUser dao.User

	err := json.NewDecoder(w.Body).Decode(&newUser)
	require.NoError(t, err)
	require.Equal(t, http.StatusCreated, w.Code)
	require.Equal(t, newUser.Name, "Foo")
	require.Equal(t, newUser.Email, "Foo@gmail.com")
}

func TestCreateUserInvalidRequest(t *testing.T) {
	router := gin.Default()
	h := NewTestHandler()
	router.POST("/user", h.CreateUser)

	jsonBody := []byte(`{"name": "Foo"}`)
	bodyReader := bytes.NewReader(jsonBody)

	req, _ := http.NewRequest("POST", "/user", bodyReader)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	require.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetUserById(t *testing.T) {
	router := gin.Default()
	h := NewTestHandler()
	router.GET("/user/:id", h.GetUserByID)

	validId := h.DB.GetUsers()[0].ID

	req, _ := http.NewRequest("GET", "/user/"+validId, nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	var firstUser dao.User

	err := json.NewDecoder(w.Body).Decode(&firstUser)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, w.Code)
	require.Equal(t, firstUser.ID, validId)
}

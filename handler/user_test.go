package handler

import (
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

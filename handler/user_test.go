package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/camille-plays/coffee2go/dao"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	tests := []struct {
		name               string
		id                 string
		store              *dao.MockStore
		expectedUser       dao.User
		expectedStatusCode int
	}{
		{
			name: "user exists",
			id:   "96840db2-3676-4399-847e-82e9d2667457",
			store: &dao.MockStore{
				Users: []*dao.User{
					{
						ID:     "96840db2-3676-4399-847e-82e9d2667457",
						Name:   "John",
						Email:  "john@email.com",
						Credit: 10,
					},
				},
			},
			expectedUser: dao.User{
				ID:     "96840db2-3676-4399-847e-82e9d2667457",
				Name:   "John",
				Email:  "john@email.com",
				Credit: 10,
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name: "user doesn't exist",
			id:   uuid.NewString(),
			store: &dao.MockStore{
				Users: []*dao.User{
					{
						ID:     "96840db2-3676-4399-847e-82e9d2667457",
						Name:   "John",
						Email:  "john@email.com",
						Credit: 10,
					},
				},
			},
			expectedStatusCode: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			router := gin.Default()
			h := Handler{DB: tt.store}

			router.GET("/user/:id", h.GetUserByID)
			req, _ := http.NewRequest("GET", "/user/"+tt.id, nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			var user dao.User
			require.Equal(t, tt.expectedStatusCode, w.Code)

			if tt.expectedStatusCode == http.StatusOK {
				err := json.NewDecoder(w.Body).Decode(&user)
				require.NoError(t, err)
				require.Equal(t, user, tt.expectedUser)
			}
		})
	}

}

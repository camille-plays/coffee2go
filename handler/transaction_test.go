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

func TestGetTransactions(t *testing.T) {
	router := gin.Default()
	h := NewTestHandler()
	router.GET("/transactions", h.GetTransactions)

	req, _ := http.NewRequest("GET", "/transactions", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	var transactions []dao.Transaction

	err := json.NewDecoder(w.Body).Decode(&transactions)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, w.Code)
	require.Equal(t, transactions, h.DB.GetTransactions())
}

// func TestCreateTransaction(t *testing.T) {
// 	router := gin.Default()
// 	h := NewTestHandler()
// 	router.POST("/transaction", h.CreateTransaction)

// 	owner := h.DB.GetUsers()[0]
// 	ownerCredit := owner.Credit

// 	recipient := h.DB.GetUsers()[1]
// 	recipientCredit := recipient.Credit

// 	var recipients []string
// 	recipients = append(recipients, owner.ID, recipient.ID)

// 	jsonBody := []byte(`{"owner": owner.ID, "recipients": recipients}`)
// 	bodyReader := bytes.NewReader(jsonBody)

// 	req, _ := http.NewRequest("POST", "/transaction", bodyReader)
// 	w := httptest.NewRecorder()
// 	router.ServeHTTP(w, req)

// 	var newTransaction dao.Transaction

// 	err := json.NewDecoder(w.Body).Decode(&newTransaction)
// 	require.NoError(t, err)
// 	require.Equal(t, http.StatusCreated, w.Code)
// 	require.Equal(t, newTransaction.Owner, ownerCredit+1)
// 	require.Equal(t, newTransaction.Recipients[1], recipientCredit-1)
// }

func TestCreateTransactionInvalidRequest(t *testing.T) {
	router := gin.Default()
	h := NewTestHandler()
	router.POST("/transaction", h.CreateTransaction)

	jsonBody := []byte(`{"owner": "bla"`)
	bodyReader := bytes.NewReader(jsonBody)

	req, _ := http.NewRequest("POST", "/transaction", bodyReader)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	require.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetTransactionById(t *testing.T) {
	router := gin.Default()
	h := NewTestHandler()
	router.GET("/transaction/:id", h.GetTransactionById)

	validId := h.DB.GetTransactions()[0].ID

	req, _ := http.NewRequest("GET", "/transaction/"+validId, nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	var firstTransaction dao.Transaction

	err := json.NewDecoder(w.Body).Decode(&firstTransaction)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, w.Code)
	require.Equal(t, firstTransaction.ID, validId)
}

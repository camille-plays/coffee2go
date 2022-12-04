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

func TestCreateTransaction(t *testing.T) {
	tests := []struct {
		name                string
		transaction         TransactionRequest
		store               *dao.MockStore
		expectedOwnerCredit int
		expectedStatusCode  int
		expectedError       string
	}{
		{
			name: "create transaction success",
			transaction: TransactionRequest{
				Owner: "96840db2-3676-4399-847e-82e9d2667457",
				Recipients: []string{
					"96840db2-3676-4399-847e-82e9d2667457",
					"b1af6aba-9ec1-4f7b-ad78-d8e4496d9cbe",
				},
			},
			store: &dao.MockStore{
				Users: []*dao.User{
					{
						ID:     "96840db2-3676-4399-847e-82e9d2667457",
						Name:   "John",
						Email:  "john@email.com",
						Credit: 10,
					},
					{
						ID:     "b1af6aba-9ec1-4f7b-ad78-d8e4496d9cbe",
						Name:   "Jane",
						Email:  "Jane@email.com",
						Credit: 10,
					},
				},
			},
			expectedOwnerCredit: 11,
			expectedStatusCode:  http.StatusCreated,
			expectedError:       "owner or recipient does not exist",
		},

		{
			name:        "create transaction failed",
			transaction: TransactionRequest{},
			store: &dao.MockStore{
				Users: []*dao.User{
					{
						ID:     "96840db2-3676-4399-847e-82e9d2667457",
						Name:   "John",
						Email:  "john@email.com",
						Credit: 10,
					},
					{
						ID:     "b1af6aba-9ec1-4f7b-ad78-d8e4496d9cbe",
						Name:   "Jane",
						Email:  "Jane@email.com",
						Credit: 10,
					},
				},
			},
			expectedOwnerCredit: 10,
			expectedStatusCode:  http.StatusBadRequest,
			expectedError:       "please provide an owner",
		},
		{
			name: "create transaction failed when owner isnt found",
			transaction: TransactionRequest{
				Owner: uuid.NewString(),
				Recipients: []string{
					"96840db2-3676-4399-847e-82e9d2667457",
					"b1af6aba-9ec1-4f7b-ad78-d8e4496d9cbe",
				},
			},
			store: &dao.MockStore{
				Users: []*dao.User{
					{
						ID:     "96840db2-3676-4399-847e-82e9d2667457",
						Name:   "John",
						Email:  "john@email.com",
						Credit: 10,
					},
					{
						ID:     "b1af6aba-9ec1-4f7b-ad78-d8e4496d9cbe",
						Name:   "Jane",
						Email:  "Jane@email.com",
						Credit: 10,
					},
				},
			},
			expectedOwnerCredit: 10,
			expectedStatusCode:  http.StatusBadRequest,
			expectedError:       "owner or recipient does not exist",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			router := gin.Default()
			h := Handler{DB: tt.store}

			router.POST("/transaction", h.CreateTransaction)
			body, _ := json.Marshal(tt.transaction)
			req, _ := http.NewRequest("POST", "/transaction", bytes.NewReader(body))
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			require.Equal(t, tt.expectedStatusCode, w.Code)

			if tt.expectedStatusCode == http.StatusOK {
				var newTransaction dao.Transaction
				err := json.NewDecoder(w.Body).Decode(&newTransaction)
				require.NoError(t, err)
				require.Equal(t, h.DB.GetUser(newTransaction.Owner).Credit, tt.expectedOwnerCredit)
			}
			if tt.expectedStatusCode == http.StatusBadRequest {
				var err ErrorResponse
				require.NoError(t, json.NewDecoder(w.Body).Decode(&err))
				require.Equal(t, tt.expectedError, err.Error)
			}
		})
	}
}

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

// func TestGetTransactionById(t *testing.T) {
// 	router := gin.Default()
// 	h := NewTestHandler()
// 	var transactionID string = uuid.New().String()
// 	h.DB.CreateTransaction()
// 	router.GET("/transaction/:id", h.GetTransactionById)

// 	validId := h.DB.GetTransactions()[0].ID

// 	req, _ := http.NewRequest("GET", "/transaction/"+validId, nil)
// 	w := httptest.NewRecorder()
// 	router.ServeHTTP(w, req)

// 	var firstTransaction dao.Transaction

// 	err := json.NewDecoder(w.Body).Decode(&firstTransaction)
// 	require.NoError(t, err)
// 	require.Equal(t, http.StatusOK, w.Code)
// 	require.Equal(t, firstTransaction.ID, validId)
// }

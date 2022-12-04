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

func TestGetTransactionById(t *testing.T) {

	tests := []struct {
		name                string
		id                  string
		store               *dao.MockStore
		expectedTransaction dao.Transaction
		expectedStatusCode  int
		expectedError       string
	}{
		{
			name: "transaction exists",
			id:   "a35d3f60-1e40-4a33-94dd-99e58121c29e",
			store: &dao.MockStore{
				Transactions: []dao.Transaction{
					{
						ID:    "a35d3f60-1e40-4a33-94dd-99e58121c29e",
						Owner: "96840db2-3676-4399-847e-82e9d2667457",
						Recipients: []string{
							"96840db2-3676-4399-847e-82e9d2667457",
							"b1af6aba-9ec1-4f7b-ad78-d8e4496d9cbe",
						},
						Timestamp: 1,
					},
				},
			},
			expectedTransaction: dao.Transaction{
				ID:    "a35d3f60-1e40-4a33-94dd-99e58121c29e",
				Owner: "96840db2-3676-4399-847e-82e9d2667457",
				Recipients: []string{
					"96840db2-3676-4399-847e-82e9d2667457",
					"b1af6aba-9ec1-4f7b-ad78-d8e4496d9cbe",
				},
				Timestamp: 1,
			},
			expectedStatusCode: http.StatusOK,
		},

		{
			name: "transaction doesn't exist",
			id:   uuid.NewString(),
			store: &dao.MockStore{
				Transactions: []dao.Transaction{
					{
						ID:    "a35d3f60-1e40-4a33-94dd-99e58121c29e",
						Owner: "96840db2-3676-4399-847e-82e9d2667457",
						Recipients: []string{
							"96840db2-3676-4399-847e-82e9d2667457",
							"b1af6aba-9ec1-4f7b-ad78-d8e4496d9cbe",
						},
						Timestamp: 1,
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

			router.GET("/transaction/:id", h.GetTransactionById)
			req, _ := http.NewRequest("GET", "/transaction/"+tt.id, nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			require.Equal(t, tt.expectedStatusCode, w.Code)

			if tt.expectedStatusCode == http.StatusOK {
				var transaction dao.Transaction
				err := json.NewDecoder(w.Body).Decode(&transaction)
				require.NoError(t, err)
				require.Equal(t, transaction, tt.expectedTransaction)
			}
		})
	}
}

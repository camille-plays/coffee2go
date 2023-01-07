package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/camille-plays/coffee2go/dao"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h Handler) CreateTransaction(c *gin.Context) {
	var transactionRequest TransactionRequest

	if err := c.BindJSON(&transactionRequest); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	if err := h.validateTransaction(transactionRequest); err != nil {
		c.IndentedJSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	transaction := transactionFromRequest(transactionRequest)

	h.DB.CreateTransaction(&transaction)
	c.IndentedJSON(http.StatusCreated, transaction)
}

func (h Handler) GetTransactions(c *gin.Context) {
	transactions := h.DB.GetTransactions()
	c.IndentedJSON(http.StatusOK, transactions)
}

func (h Handler) GetTransactionById(c *gin.Context) {
	// Make sure provided ID is a UUID
	_, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	transaction := h.DB.GetTransaction(c.Param("id"))
	if transaction != nil {
		c.IndentedJSON(http.StatusOK, transaction)
		return
	}

	c.Status(http.StatusNotFound)
}

func transactionFromRequest(t TransactionRequest) dao.Transaction {
	return dao.Transaction{
		Owner:      t.Owner,
		Recipients: t.Recipients,
		ID:         uuid.New().String(),
		Timestamp:  int(time.Now().Unix()),
	}
}

func (h Handler) validateTransaction(tr TransactionRequest) error {
	if tr.Owner == "" {
		return fmt.Errorf("please provide an owner")
	}

	if tr.Recipients == nil || len(tr.Recipients) == 0 {
		return fmt.Errorf("please provide a list of recipients")
	}

	for _, u := range append(tr.Recipients, tr.Owner) {
		if h.DB.GetUser(u) == nil {
			return fmt.Errorf("owner or recipient does not exist")
		}
	}
	return nil
}

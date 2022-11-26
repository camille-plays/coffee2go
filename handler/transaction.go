package handler

import (
	"fmt"
	"net/http"

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

	if err := validateTransaction(transactionRequest); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	transaction := transactionFromRequest(transactionRequest)

	h.DB.CreateTransaction(transaction)
	c.Status(http.StatusOK)
	return
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
	return
}

func transactionFromRequest(t TransactionRequest) dao.Transaction {
	return dao.Transaction{
		Owner:      t.Owner,
		Recipients: t.Recipients,
		ID:         uuid.New().String(),
	}
}

func validateTransaction(transactionRequest TransactionRequest) error {

	if transactionRequest.Owner == "" {
		return fmt.Errorf("Please provide an owner")
	}

	if transactionRequest.Recipients == nil {
		return fmt.Errorf("please provide a list of recipients")
	}
	return nil
}

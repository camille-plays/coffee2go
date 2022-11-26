package handler

import (
	"net/http"

	"github.com/camille-plays/coffee2go/models"
	"github.com/gin-gonic/gin"
)

func CreateTransaction(c *gin.Context) {
	var transactionRequest TransactionRequest

	if err := c.BindJSON(&transactionRequest); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	// increment credit of owner
	for _, a := range models.Users {
		if transactionRequest.Owner == a.ID {
			a.Credit += len(transactionRequest.Recipients)
		}
	}

	// decrement credits of recipients
	for _, r := range transactionRequest.Recipients {
		for _, u := range models.Users {
			if r == u.ID {
				u.Credit -= 1
			}
		}
	}
	c.Status(http.StatusOK)
	return
}

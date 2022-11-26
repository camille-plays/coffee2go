package main

import (
	"github.com/camille-plays/coffee2go/dao"
	"github.com/camille-plays/coffee2go/handler"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func main() {

	router := gin.Default()
	h := handler.NewTestHandler()

	h.DB.CreateTransaction(&dao.Transaction{
		ID:         uuid.NewString(),
		Owner:      "96840db2-3676-4399-847e-82e9d2667457",
		Recipients: []string{"b1af6aba-9ec1-4f7b-ad78-d8e4496d9cbe"},
	})

	router.GET("/users", h.GetUsers)
	router.GET("/user/:id", h.GetUserByID)
	router.POST("/user", h.CreateUser)
	router.GET("transactions", h.GetTransactions)
	router.GET("transaction/:id", h.GetTransactionById)
	router.POST("transaction", h.CreateTransaction)

	router.Run("localhost:8080")
}

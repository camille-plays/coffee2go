package main

import (
	"github.com/gin-gonic/gin"

	"github.com/camille-plays/coffee2go/handler"
)

func main() {

	router := gin.Default()
	h := handler.NewTestHandler()

	router.GET("/users", h.GetUsers)
	router.GET("/user/:id", h.GetUserByID)
	router.POST("/user", h.CreateUser)
	router.GET("transactions", h.GetTransactions)
	router.GET("transaction/:id", h.GetTransactionById)
	router.POST("transaction", h.CreateTransaction)

	router.Run("localhost:8080")
}

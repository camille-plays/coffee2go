package main

import (
	"github.com/camille-plays/coffee2go/handler"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	h := handler.NewLocalHandler()

	router.GET("/users", h.GetUsers)
	router.GET("/user/:id", h.GetUserByID)
	router.POST("/user", h.CreateUser)
	router.GET("transactions", h.GetTransactions)
	router.GET("transaction/:id", h.GetTransactionById)
	router.POST("transaction", h.CreateTransaction)

	router.Run("localhost:8080")
}

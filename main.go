package main

import (
	"github.com/gin-gonic/gin"

	"github.com/camille-plays/coffee2go/handler"
)

func main() {
	router := gin.Default()
	router.GET("/users", handler.GetUsers)
	router.GET("/user/:id", handler.GetUserByID)
	router.POST("/user", handler.CreateUser)
	router.POST("/transaction", handler.CreateTransaction)

	router.Run("localhost:8080")
}

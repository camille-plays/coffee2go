package main

import (
	"github.com/gin-gonic/gin"

	"github.com/camille-plays/coffee2go/httpd/handler"
)

func main() {
	router := gin.Default()
	router.GET("/people", handler.GetPeople)
	router.POST("/people", handler.PostPeople)
	router.GET("/people/:id", handler.GetPersonByID)
	router.PUT("/people/:id/incrementCredit", handler.IncrementCredit)
	router.PUT("/people/:id/decrementCredit", handler.DecrementCredit)

	router.Run("localhost:8080")
}

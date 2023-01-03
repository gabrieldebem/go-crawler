package main

import (
	"github.com/gabrieldebem/go-crawler/packages/contracts"
	"github.com/gabrieldebem/go-crawler/packages/handlers"
	"github.com/gin-gonic/gin"
)

var crawler contracts.ICrawler

func main() {
	r := gin.Default()

	api := r.Group("/api")
	{
		api.GET("/products", handlers.GetProducts)
	}

	r.Run()
}

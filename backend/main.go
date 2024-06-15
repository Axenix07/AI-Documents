package main

import (
	"backend/document"
	"backend/summary"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/documents", document.GetAllDocuments)
	router.POST("/documents", document.PostDocument)

	router.GET("/summaries", summary.GetAllSummaries)
	router.POST("/summaries", summary.PostSummary)

	router.Run("localhost:8080")
}

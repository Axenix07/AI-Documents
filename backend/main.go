package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

type document struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Summary string `json:"summary"`
	Document string `json:"document"`
}

var documents = []document{
    {ID: "1", Name: "Document 1", Summary: "Initial Summary", Document: "The First Document Body"},
    {ID: "2", Name: "Document 2", Summary: "Initial Summary", Document: "The Second Document Body"},
}

func main() {
	router := gin.Default()
	router.GET("/documents", getDocuments)  
	router.POST("/documents", postDocuments)

	router.Run("localhost:8080")
}

func getDocuments(c *gin.Context){
	c.IndentedJSON(http.StatusOK, documents)
}

func postDocuments(c *gin.Context) {
	var newDocument document

	if err := c.BindJSON(&newDocument); err != nil {
		return
	}

	documents = append(documents, newDocument)
	c.IndentedJSON(http.StatusCreated, newDocument)
}
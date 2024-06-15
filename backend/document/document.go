package document

import (
	"backend/summary"
	"net/http"

	"github.com/gin-gonic/gin"
)

type document struct {
	ID       string          `json:"id"`
	Name     string          `json:"name"`
	Summary  summary.Summary `json:"summary"`
	Document string          `json:"document"`
}

func GetAllDocuments(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, documents)
}

func PostDocument(c *gin.Context) {
	var newDocument document

	if err := c.BindJSON(&newDocument); err != nil {
		return
	}

	documents = append(documents, newDocument)
	c.IndentedJSON(http.StatusCreated, newDocument)
}

var documents = []document{
	{ID: "1", Name: "Document 1", Summary: "1", Document: "The First Document Body"},
	{ID: "2", Name: "Document 2", Summary: "2", Document: "The Second Document Body"},
}

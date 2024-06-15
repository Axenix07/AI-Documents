package summary

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Summary struct {
	Id         string `json:"id"`
	Summary    string `json:"summary"`
	DocumentID string `json:"documentID"`
}

func GetAllSummaries(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, summaries)
}

func PostSummary(c *gin.Context) {

}

var summaries = []Summary{
	{Summary: "The First Summary", DocumentID: "1"},
	{Summary: "The Second Summary", DocumentID: "2"},
}

package analyzer

import (
	api "GOTASK/api/repository"
	"GOTASK/services"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func AnalyzeText(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.String(http.StatusBadRequest, "error reading file: %v", err)
			return
		}

		data, _ := file.Open()
		defer data.Close()
		fileData, _ := io.ReadAll(data)
		text := string(fileData)

		result, timeTaken := services.TextAnalysis(text)

		err = api.Result(db, file.Filename, result, timeTaken.String())
		if err != nil {
			c.String(http.StatusInternalServerError, "error saving to database: %v", err)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"Filename":       file.Filename,
			"Words":          result.Words,
			"Lines":          result.Lines,
			"Spaces":         result.Spaces,
			"Punctuation":    result.Punc,
			"Characters":     result.Characters,
			"Sentences":      result.Sentences,
			"Vowels":         result.Vowels,
			"Consonants":     result.Consonants,
			"Digits":         result.Digits,
			"Paragraphs":     result.Para,
			"Execution time": timeTaken.String(),
		})
	}
}

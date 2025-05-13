package handler

import (
	"GOTASK/combo"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func AnalyzeText(db *sqlx.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("Reading File")

		file, err := ctx.FormFile("sample")
		if err != nil {
			ctx.String(http.StatusBadRequest, "error Reading File: %v", err)
		}
		file2, _ := file.Open()
		defer file2.Close()
		FileData, _ := io.ReadAll(file2)

		Str := string(FileData)

		startTime2 := time.Now()
		ch := make(chan []int)
		go combo.Combo(Str, ch)
		value := <-ch
		timeTaken2 := time.Since(startTime2)
		_, err = db.Exec(
			`INSERT INTO gotask(
		    filename, 
		file_size, 
		word_count, 
		line_count, 
		sentence_count, 
		paragraph_count, 
		punctuation_count, 
		special_char_count, 
		vowel_count, 
		consonant_count, 
		digit_count,
		combine_time
			)VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12)`,
			file.Filename,
			file.Size,
			value[0],
			value[1],
			value[2],
			value[3],
			value[4],
			value[5],
			value[6],
			value[7],
			value[8],
			timeTaken2.String(),
		)

		if err != nil {
			ctx.String(http.StatusInternalServerError, "Error Inserting Data: %v", err)
			return
		}
		ctx.JSON(200, gin.H{
			"Filename":      file.Filename,
			"File Size":     file.Size,
			"Word Count ":   value[0],
			"Lines ":        value[1],
			"sentences":     value[2],
			"paragraphs":    value[3],
			"punctuations":  value[4],
			"specialChar":   value[5],
			"vowelscount":   value[6],
			"consonants":    value[7],
			"digits":        value[8],
			"Time Taking: ": timeTaken2,
		})

	}
}

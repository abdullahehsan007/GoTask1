package main

import (
	"GOTASK/chunks"
	"GOTASK/combo"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {

	fmt.Println("Reading File")
	router := gin.Default()
	router.POST("/analyze", func(ctx *gin.Context) {
		file, err := ctx.FormFile("sample")
		if err != nil {
			ctx.String(http.StatusBadRequest, "error Reading File: %v", err)
		}
		file2, _ := file.Open()
		defer file2.Close()
		FileData, _ := io.ReadAll(file2)

		Str := string(FileData)
		// startTime := time.Now()
		// var wg sync.WaitGroup
		// wg.Add(9)
		// go separate.WordCounter(Str, &wg)
		// go separate.LinesCounter(Str, &wg)
		// go separate.SentenceCounter(Str, &wg)
		// go separate.ParasCounter(Str, &wg)
		// go separate.PuncCounter(Str, &wg)
		// go separate.SpecialCounter(Str, &wg)
		// go separate.VowelsCounter(Str, &wg)
		// go separate.ConsonantsCounter(Str, &wg)
		// go separate.DigitsCounter(Str, &wg)
		// wg.Wait()
		// timeTaken := time.Since(startTime)
		// fmt.Println("Time Taken After Using Goroutine: ", timeTaken)
		// fmt.Printf("\n")

		startTime2 := time.Now()
		ch := make(chan []int)
		// //var names = [9]string{"Word Count", "Lines", "Sentences", "Paragraphs", "Punctuations", "Special Characters", "Vowels", "Consonants", "Digits"}
		go combo.Combo(Str, ch)
		value := <-ch
		timeTaken2 := time.Since(startTime2)
		// ctx.JSON(200, gin.H{
		// 	"Filename":     file.Filename,
		// 	"File Size":    file.Size,
		// 	"Word Count ":  value[0],
		// 	"Lines ":       value[1],
		// 	"sentences":    value[2],
		// 	"paragraphs":   value[3],
		// 	"punctuations": value[4],
		// 	"specialChar":  value[5],
		// 	"vowelscount":  value[6],
		// 	"consonants":   value[7],
		// 	"digits":       value[8],
		// 	"Time Taking: " : timeTaken2,
		// })

		// for i := 0; i < 9; i++ {

		// 	fmt.Println(names[i], value[i])

		// }

		//	fmt.Println("Time Taken After Using Channels & Goroutines: ", timeTaken2)

		startTime3 := time.Now()
		result := chunks.Chunks(Str)
		timeTaken3 := time.Since(startTime3)
		ctx.JSON(200, gin.H{
			"Filename":                                file.Filename,
			"Using Combination Function Word Count ":  value[0],
			"Using Combination Function Lines ":       value[1],
			"Using Combination Function sentences":    value[2],
			"Using Combination Function paragraphs":   value[3],
			"Using Combination Function punctuations": value[4],
			"Using Combination Function specialChar":  value[5],
			"Using Combination Function vowelscount":  value[6],
			"Using Combination Function consonants":   value[7],
			"Using Combination Function digits":       value[8],
			"Time Taken for Combination Function: ":   timeTaken2.String(),
			"Chunk Result":                            result,
			"Time Taken for Chunks ":                  timeTaken3.String(),
		})

	})
	router.Run(":8080")
}

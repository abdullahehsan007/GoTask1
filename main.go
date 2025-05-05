package main

import (
	"GOTASK/combo"
	"GOTASK/separate"
	"fmt"
	"os"
	"sync"
	"time"
)

func main() {
	startTime := time.Now()

	fmt.Println("Reading File")
	textfile := "Test.txt"
	FileData, err := os.ReadFile(textfile)
	if err != nil {
		panic(err)
	}
	Str := string(FileData)
	var wg sync.WaitGroup
	wg.Add(9)
	go separate.WordCounter(Str, &wg)
	go separate.LinesCounter(Str, &wg)
	go separate.SentenceCounter(Str, &wg)
	go separate.ParasCounter(Str, &wg)
	go separate.PuncCounter(Str, &wg)
	go separate.SpecialCounter(Str, &wg)
	go separate.VowelsCounter(Str, &wg)
	go separate.ConsonantsCounter(Str, &wg)
	go separate.DigitsCounter(Str, &wg)
	wg.Wait()
	timeTaken := time.Since(startTime)

	fmt.Println("Time Taken After Using Goroutine: ", timeTaken)
	startTime2 := time.Now()

	ch := make(chan int, 10)
	var names = [10]string{"Word Count", "Spaces", "Lines", "Sentences", "Paragraphs", "Punctuations", "Special Characters", "Vowels", "Consonants", "Digits"}
	go combo.Combo(Str, ch)
	for i := 0; i < 9; i++ {
		value := <-ch
		fmt.Println(names[i], value)

	}
	timeTaken2 := time.Since(startTime2)
	fmt.Println("Time Taken After Using Channels & Goroutines: ", timeTaken2)
}

package chunks

import (
	"GOTASK/combo"
	"fmt"
)

func Chunks(fileData string) {
	chunksize := len(fileData) / 4
	ch := make(chan []int)
	var names = [10]string{"Word Count", "Lines", "Sentences", "Paragraphs", "Punctuations", "Special Characters", "Vowels", "Consonants", "Digits"}

	chunk := []string{
		fileData[:chunksize],
		fileData[chunksize : chunksize*2],
		fileData[chunksize*2 : chunksize*3],
		fileData[chunksize*3:],
	}
	for i := 0; i < 4; i++ {
		
		go combo.Combo(chunk[i], ch)
		fmt.Printf("\n")
		fmt.Println("Chunk No:", i)
		value := <-ch
		for j := 0; j < 9; j++ {
			fmt.Println(names[j], value[j])

		}
	}

}

package chunks

import (
	"GOTASK/combo"
)

type chunkresult struct {
	ChunkNumber       int
	wordCount         int
	Lines             int
	Sentences         int
	Paragraphs        int
	Punctuations      int
	Special_Character int
	Vowels            int
	Consonants        int
	Digits            int
}

func Chunks(fileData string) []chunkresult {
	chunksize := len(fileData) / 4
	ch := make(chan []int)
	//var names = [10]string{"Word Count", "Lines", "Sentences", "Paragraphs", "Punctuations", "Special Characters", "Vowels", "Consonants", "Digits"}

	chunk := []string{
		fileData[:chunksize],
		fileData[chunksize : chunksize*2],
		fileData[chunksize*2 : chunksize*3],
		fileData[chunksize*3:],
	}
	var result []chunkresult
	for i := 0; i < 4; i++ {

		go combo.Combo(chunk[i], ch)
		value := <-ch
		result = append(result, chunkresult{
			ChunkNumber:       i + 1,
			wordCount:         value[0],
			Lines:             value[1],
			Sentences:         value[2],
			Paragraphs:        value[3],
			Punctuations:      value[4],
			Special_Character: value[5],
			Vowels:            value[6],
			Consonants:        value[7],
			Digits:            value[8],
		})

	}
	return result

}

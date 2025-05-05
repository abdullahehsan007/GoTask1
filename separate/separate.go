package separate

import (
	"fmt"
	"strings"
	"sync"
)

func WordCounter(fileData string, wg *sync.WaitGroup) int {
	var wordCount int = 0
	for i := 0; i < len(fileData); i++ {
		switch fileData[i] {
		case ' ', '.':
			{
				wordCount++
			}
		}

	}
	fmt.Println("Words:", wordCount)
	defer wg.Done()

	return int(wordCount)
}
func LinesCounter(fileData string, wg *sync.WaitGroup) int {
	var lines int = 1
	for i := 0; i < len(fileData); i++ {
		switch fileData[i] {
		case '\n':
			{
				lines++
			}
		}

	}
	defer wg.Done()

	fmt.Println("Lines", lines)
	return int(lines)
}
func SentenceCounter(fileData string, wg *sync.WaitGroup) int {
	var sentences int = 0
	for i := 0; i < len(fileData); i++ {
		if fileData[i] == '.' {
			if i+1 < len(fileData) && fileData[i+1] == ' ' {
				sentences++
			}

		}

	}
	defer wg.Done()

	fmt.Println("Sentences", sentences)
	return int(sentences)
}
func ParasCounter(fileData string, wg *sync.WaitGroup) {
	var Paragraphs int = 1
	for i := 0; i < len(fileData); i++ {
		if fileData[i] == '\n' {
			// if i+1<len(fileData) && fileData[i+1] == ' ' {
			if i+1 < len(fileData) && fileData[i+1] == '\n' {
				Paragraphs++
			}
		}
	}

	//}
	defer wg.Done()

	fmt.Println("Paragraphs", Paragraphs)
	//return int(Paragraphs)
}

func PuncCounter(fileData string, wg *sync.WaitGroup) int {
	var punctuations int = 0
	for i := 0; i < len(fileData); i++ {
		switch fileData[i] {
		case ',', ';', '"', '/', ':', '.', '\'', '?', '`':
			{
				punctuations++
			}
		}
	}
	defer wg.Done()

	fmt.Println("Punctuations:", punctuations)
	return int(punctuations)
}
func SpecialCounter(fileData string, wg *sync.WaitGroup) int {
	var specialChar int = 0
	for i := 0; i < len(fileData); i++ {
		switch fileData[i] {
		case '!', '@', '#', '$', '%', '^', '&', '*',
			'(', ')', '[', ']', '{', '}', '+', '=', '-',
			'_', '\\', '|', '<', '>', '?', '~', '`':
			{
				specialChar++
			}
		}
	}
	defer wg.Done()

	fmt.Println("Special Characters:", specialChar)
	return int(specialChar)
}
func VowelsCounter(fileData string, wg *sync.WaitGroup) int {
	var vowelscount int = 0
	for i := 0; i < len(fileData); i++ {
		switch fileData[i] {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			{
				vowelscount++
			}
		}
	}
	defer wg.Done()

	fmt.Println("Vowels:", vowelscount)
	return int(vowelscount)
}
func ConsonantsCounter(fileData string, wg *sync.WaitGroup) int {
	var consonants int = 0
	for i := 0; i < len(fileData); i++ {
		upper := strings.ToUpper(string(fileData[i]))[0]
		switch upper {
		case 'B', 'C', 'D', 'F', 'G', 'H', 'J', 'K', 'L', 'M',
			'N', 'P', 'Q', 'R', 'S', 'T', 'V', 'W', 'X', 'Y', 'Z':
			consonants++
		}
	}
	defer wg.Done()

	fmt.Println("Consonants:", consonants)
	return int(consonants)
}

func DigitsCounter(fileData string, wg *sync.WaitGroup) int {
	var digits int = 0
	for i := 0; i < len(fileData); i++ {
		switch fileData[i] {
		case '1', '2', '3', '4', '5', '6', '7', '8', '9', '0':
			{
				digits++
			}
		}
	}
	defer wg.Done()
	fmt.Println("Digits:", digits)
	return int(digits)
}

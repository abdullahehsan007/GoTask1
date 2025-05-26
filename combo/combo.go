package combo

type ResultStruct struct {
	Para       int
	Words      int
	Spaces     int
	Lines      int
	Sentences  int
	Punc       int
	Characters int
	Digits     int
	Vowels     int
	Consonants int
}

func Combo (fileData string, resultchann chan<- ResultStruct) {
	para := 1
	words := 0
	spaces := 0
	lines := 1
	senten := 0
	punc := 0
	character := 0
	digit := 0
	vowels := 0
	consonants := 0

	for i := 0; i < len(fileData); i++ {
		ch := fileData[i]
		switch ch {
		case ' ':
			{
				words++
				spaces++

			}
		case '\n':
			{
				lines++
				if i+1 < len(fileData) && fileData[i+1] == '\n' {
					para++
				}
			}
		case '.', ',', ';', '"', '/', ':', '\'', '?', '`':
			if i+1 < len(fileData) && fileData[i+1] == ' ' {
				senten++

			}
			punc++
		case '!', '@', '#', '$', '%', '^', '&', '*',
			'(', ')', '[', ']', '{', '}', '+', '=', '-',
			'_', '\\', '|', '<', '>', '~':
			{
				character++
			}
		case 'A', 'E', 'I', 'O', 'U', 'a', 'e', 'i', 'o', 'u':
			{
				vowels++
			}
		case 'B', 'C', 'D', 'F', 'G', 'H', 'J', 'K', 'L', 'M',
			'N', 'P', 'Q', 'R', 'S', 'T', 'V', 'W', 'X', 'Y', 'Z', 'b', 'c', 'd', 'f', 'g', 'h', 'j', 'k', 'l', 'm', 'n', 'p', 'q', 'r', 's', 't', 'v', 'w', 'x', 'y', 'z':
			{
				consonants++
			}

		case '1', '2', '3', '4', '5', '6', '7', '8', '9', '0':
			{
				digit++
			}
		}
	}

	result := ResultStruct{
		Para:       para,
		Words:      words,
		Spaces:     spaces,
		Lines:      lines,
		Sentences:  senten,
		Punc:       punc,
		Characters: character,
		Digits:     digit,
		Vowels:     vowels,
		Consonants: consonants,
	}
	resultchann <- result

	close(resultchann)
}

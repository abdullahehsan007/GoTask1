package combo

func Combo(fileData string, ch chan<- int) {
	var wordCount int = 0
	var lines int = 1
	var sentences int = 0
	var paragraphs int = 1
	var vowelscount int = 0
	var punctuations int = 0
	var specialChar int = 0
	var consonants int = 0
	var digits int = 0
	for i := 0; i < len(fileData); i++ {
		switch fileData[i] {
		case ' ':
			{
				wordCount++

			}
		case '\n':
			{
				lines++
				if i+1 < len(fileData) && fileData[i+1] == '\n' {
					paragraphs++
				}
			}
		case '.', ',', ';', '"', '/', ':', '\'', '?', '`':
			if i+1 < len(fileData) && fileData[i+1] == ' ' {
				sentences++

			}

			punctuations++
		case '!', '@', '#', '$', '%', '^', '&', '*',
			'(', ')', '[', ']', '{', '}', '+', '=', '-',
			'_', '\\', '|', '<', '>', '~':
			{
				specialChar++
			}
		case 'A', 'E', 'I', 'O', 'U', 'a', 'e', 'i', 'o', 'u':
			{
				vowelscount++
			}
		case 'B', 'C', 'D', 'F', 'G', 'H', 'J', 'K', 'L', 'M',
			'N', 'P', 'Q', 'R', 'S', 'T', 'V', 'W', 'X', 'Y', 'Z', 'b', 'c', 'd', 'f', 'g', 'h', 'j', 'k', 'l', 'm', 'n', 'p', 'q', 'r', 's', 't', 'v', 'w', 'x', 'y', 'z':
			{
				consonants++
			}

		case '1', '2', '3', '4', '5', '6', '7', '8', '9', '0':
			{
				digits++
			}
		}
	}
	ch <- wordCount
	ch <- lines
	ch <- sentences
	ch <- paragraphs
	ch <- punctuations
	ch <- specialChar
	ch <- vowelscount
	ch <- consonants
	ch <- digits
	defer close(ch)

}

package asciiart

// Function PrintArt() converts word into art form
func PrintArt(asciiMap map[rune][]string, input string) string {
	store := [][]string{}

	for _, chr := range input {
		store = append(store, asciiMap[chr])
	}

	word := ""
	for i := 0; i < 8; i++ {
		for j := 0; j < len(store); j++ {
			word += store[j][i]
		}
		if i != 7 {
			word += "\n"
		}
	}
	return word
}

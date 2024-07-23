package asciiart

import (
	"strings"
)

// Function WordDistributer() splits words into chunks to be converted into ascii-art
func WordDistributor(input, bannerTemplate string) string {
	asciiMap := ReadBanner(bannerTemplate)

	artString := ""

	words := strings.Split(input, "\n")
	for i, word := range words {
		if word == "" {
			artString += "\n"
		} else {
			art := PrintArt(asciiMap, word)
			artString += art
			if i != len(words)-1 && words[len(words)-1] != "" {
				artString += "\n"
			}
		}
	}
	return artString
}

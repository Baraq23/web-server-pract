package asciiart

import (
	"fmt"
	"strings"
)


func Start(input, fileFlag string) string {
	asciiMap := ReadBanner(fileFlag)
	artString := ""
	words := strings.Split(input, "\n")
	fmt.Println(words)

	for _, word := range words {
		if word == "" {
			artString+="\n"
		} else {
			art := PrintArt(asciiMap, word)
			artString += art
		}
	}
	fmt.Println("artString:")
	fmt.Println(artString)
	// This return value is used for testing purposes.
	return artString
}

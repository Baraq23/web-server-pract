package asciiart

import "fmt"

// this function checks if the input given by the user is within the the banner files
func StrValidate(str string) bool {
	for _, v := range str {
		if v >= 32 && v <= 126 {
			continue
		} else {
			fmt.Printf("Unrecognised character: %v\n", v)
			return false
		}
	}
	return true
}

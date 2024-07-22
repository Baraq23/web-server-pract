package asciiart

import (
	"fmt"
	"os"
)

// This function checks if banner file is corrupt by checking its filesize.
func ValidateBanner(banner string) {
	standardSize := 6623
	shadowSize := 7463
	thinkertoySize := 5558
	file, err := os.Open(banner)
	checkErr(err)
	defer file.Close()
	finfo, err := file.Stat()
	checkErr(err)
	filesize := finfo.Size()
	if filesize != int64(standardSize) && filesize != int64(shadowSize) && filesize != int64(thinkertoySize) {
		fmt.Printf("Error: %v is corrupt or empty, please provide a new %v file\n", banner, banner)
		os.Exit(0)
	}
}

func checkErr(err error) {
	if err != nil {
		fmt.Println("Error: Could not ", err)
		os.Exit(0)
	}
}

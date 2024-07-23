package asciiart

import (
	"fmt"
	"os"
)

// Function ValidateBanner() checks if banner file is corrupt by checking its filesize.
func ValidateBanner(banner string) {
	//Banner template sizes in bytes
	standardSize := 6623
	shadowSize := 7463
	thinkertoySize := 5558

	file, err := os.Open(banner)
	checkErr(err)
	defer file.Close()
	fileInfo, err := file.Stat()
	checkErr(err)
	filesize := fileInfo.Size()

	if filesize != int64(standardSize) && filesize != int64(shadowSize) && filesize != int64(thinkertoySize) {
		fmt.Printf("Error: %v is corrupt or empty, refer to README.MD document to download a new %v file\n", banner, banner)
		os.Exit(1)
	}
}

func checkErr(err error) {
	if err != nil {
		fmt.Println("Error: Could not ", err)
		os.Exit(1)
	}
}

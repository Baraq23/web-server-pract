package asciiart

import (
	"testing"
	asciiart "asciiart/functions"
)
	

func TestStart(t *testing.T) {
	type test struct {
		name   string
		input  string
		output string
	}

	tests := []test{
		{"Testing for letter 'A'", "A", "           \n    /\\     \n   /  \\    \n  / /\\ \\   \n / ____ \\  \n/_/    \\_\\ \n           \n           "},
		{"Testing for number '7'", "7", "        \n _____  \n|___  | \n   / /  \n  / /   \n /_/    \n        \n        "},
		{"Testing for other characters '#'", "#", "   _  _    \n _| || |_  \n|_  __  _| \n _| || |_  \n|_  __  _| \n  |_||_|   \n           \n           "},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := asciiart.Start(tc.input, "standard.txt"); got != tc.output {
				t.Errorf("Expected:\n %v \n Got:\n %v \n", tc.output, got)
			}
		})
	}
}

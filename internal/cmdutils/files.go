package cmdutils

import (
	"fmt"
	"os"
)

func ReadFile(path string) []byte {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Failed to read json file at " + path)
		os.Exit(1)
	}
	return data
}

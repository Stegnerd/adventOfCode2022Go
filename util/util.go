package util

import (
	"fmt"
	"log"
	"os"
)

func ReadFile(fileName string) *os.File {
	file, err := os.Open(fmt.Sprintf("data/%s", fileName))
	if err != nil {
		log.Fatal(err)
	}

	return file
}

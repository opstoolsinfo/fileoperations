package filesysystem

import (
	"log"
	"os"
)

func OpenFile(filename string)  (fp *os.File){
	fp, err := os.Open(filename)
	if err != nil {
		log.Fatal("Error opening the file", err)
	}

	return fp
}
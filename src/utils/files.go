package utils

import (
	"log"
	"os"
)

func DeferCloseFile(f *os.File) {
	err := f.Close()
	if err != nil {
		log.Fatal(err)
	}
}

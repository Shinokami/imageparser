package reader

import (
	"log"
	"os"
)

func ReadFile(fileName string) []byte {
	fileData, err := os.ReadFile(fileName)

	if err != nil {
		log.Fatalf("Error when reading file: %s", err)
	}
	return fileData
}

func CreateFile(fileData []byte, fileName string) {
	err := os.WriteFile(fileName, fileData, 0644)
	if err != nil {
		log.Fatalf("Error when creating file: %s", err)
	}
}

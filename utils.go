package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func saveDataToFile(data []LinkInfo, filePath string) error {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		file, err := os.Create(filePath)
		if err != nil {
			return fmt.Errorf("error creating file %s: %w", filePath, err)
		}
		defer file.Close()
	}

	jsonData, jsonErr := json.MarshalIndent(data, "", "  ")
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	writeErr := os.WriteFile(filePath, jsonData, 0644)

	if writeErr != nil {
		return fmt.Errorf("error writing to file %s: %w", filePath, writeErr)
	}

	fmt.Printf("Data saved to file %s\n", filePath)
	return nil
}

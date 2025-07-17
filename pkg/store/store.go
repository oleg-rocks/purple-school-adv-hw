package store

import (
	"encoding/json"
	"fmt"
	"os"
)

type VerifyData struct {
	Email string `json:"email"`
	Hash string `json:"hash"`
}

const dataFile = "data.json"

func StoreVerifyData(data VerifyData) error {
	var storedData []VerifyData
	bytes, err := os.ReadFile(dataFile)
	if err == nil {
		_ = json.Unmarshal(bytes, &storedData)
	}

	storedData = append(storedData, data)

	file, err := os.Create(dataFile)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(storedData); err != nil {
		return fmt.Errorf("failed encoding JSON: %w", err)
	}
	return nil
}

func HasHash(hash string) (bool, error) {
	bytes, err := os.ReadFile(dataFile)
	if err != nil {
		return false, fmt.Errorf("failed to read file")
	}

	var storedData []VerifyData
	err = json.Unmarshal(bytes, &storedData)
	if err != nil {
		return false, fmt.Errorf("failed to decode from JSON")
	}

	for _, item := range storedData {
		if item.Hash == hash {
			return true, nil
		}
	}
	return false, fmt.Errorf("failed find hash")
}
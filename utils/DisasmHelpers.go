package utils

import "os"

func ReadFromFile(fileName string) ([]uint8, error) {
	byteCode, err := os.ReadFile(fileName)
	if err != nil {
		return []uint8{}, err
	}

	parsedCode := []uint8(byteCode)

	return parsedCode, nil
}

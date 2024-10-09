package fileoperations

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)


func GetFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return 1000, errors.New("Failed to read file.")
	}

	valueText := string(data)
	value, err := strconv.ParseFloat(valueText, 64)
	if err != nil {
		return 1000, errors.New("Failed to parse stored value.")
	}
	return value, nil

}

func WriteToFile(value float64, fileName string) {
	// Balance'ı string'e çevirirken format belirliyoruz
	valueText := fmt.Sprintf("%.2f", value)
	err := os.WriteFile(fileName, []byte(valueText), 0644) // Hata kontrolü ekliyoruz
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	fmt.Println("Value successfully written to file.")
}

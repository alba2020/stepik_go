package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
)

func readCsvFile(filePath string) ([][]string, error) {
    f, err := os.Open(filePath)
    if err != nil {
		return nil, err
    }
    defer f.Close()

    csvReader := csv.NewReader(f)
    records, err := csvReader.ReadAll()
    if err != nil {
		return nil, err
	}

    return records, nil
}

func walkFunc(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err // Если по какой-то причине мы получили ошибку, проигнорируем эту итерацию
	}

	if info.IsDir() {
		return nil // Проигнорируем директории
	}

	// fmt.Printf("Name: %s\tSize: %d byte\tPath: %s\n", info.Name(), info.Size(), path)

    records, err := readCsvFile(path)
	if err != nil {
		fmt.Printf("%v\n", err)
		return err
	}

	if len(records) > 1 {
		fmt.Println("found")
		fmt.Println(path)

		fmt.Println(records[4][2])
	}

	return nil
}

func main() {
	const root = "./task" // Файлы моей программы находятся в другой директории

	if err := filepath.Walk(root, walkFunc); err != nil {
		fmt.Printf("Какая-то ошибка: %v\n", err)
	}
}

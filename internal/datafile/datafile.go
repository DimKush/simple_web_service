package datafile

import (
	"bufio"
	"os"
)

func GetStringsFromFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	var collectionFile []string
	for scanner.Scan() {
		str := scanner.Text()
		collectionFile = append(collectionFile, str)
	}

	return collectionFile, nil
}

package util

import (
	"bufio"
	"os"
	"path/filepath"
)

//GetFiles get all file path of a give root directory
func GetFiles(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

//GetWords Get all words in file
func GetWords(filePath string) ([]string, error) {
	var words []string
	file, err := os.Open(filePath)

	defer file.Close()
	if err != nil {
		return words, err
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	return words, err
}

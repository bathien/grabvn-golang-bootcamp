package main

import (
	"fmt"
	"log"
	"sync"

	"./mapper"
	"./reducer"
	"./util"
)

func main() {
	textDirectory := "./sample_files"
	// Get the list string path of all files in the directory.
	files, err := util.GetFiles(textDirectory)
	if err != nil {
		log.Fatal("Cannot get file from directory: ", textDirectory)
	}
	lists := make(chan map[string]int)
	// Ensure all send operation is done.
	var wg sync.WaitGroup
	// Ensure the final value after Reducer is obtained.
	finalValue := make(chan map[string]int)

	// Mapping
	wg.Add(len(files))
	for _, file := range files {
		go func(file string) {
			defer wg.Done()
			listWords, err := mapper.Map(file)
			if err != nil {
				log.Fatal("Cannot get words from file: ", file)
			}
			lists <- listWords
		}(file)
	}
	// Reduction
	go reducer.Reducer(lists, finalValue)

	wg.Wait()
	close(lists)
	for wordFreq := range finalValue {
		for word, freq := range wordFreq {
			fmt.Printf("%s appear %d times\n", word, freq)
		}
		close(finalValue)
	}
}

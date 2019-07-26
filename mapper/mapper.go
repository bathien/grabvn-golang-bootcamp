package mapper

import (
	"../util"
)

//Map count frequenly word in file
func Map(fileName string) (map[string]int, error) {
	list := map[string]int{}
	// Get the list of all words in the file.
	words, err := util.GetWords(fileName)
	if err != nil {
		return list, err
	}
	for _, word := range words {
		list[word]++
	}
	return list, err
}

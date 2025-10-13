package internal

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
)

var wordSelect string

func ChooseWord(filePath string) string { // taking a random word in the 3 files
	//+ language + difficultySelected}
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
	}
	words := strings.Fields(string(data))
	wordSelect = words[rand.Intn(len(words))]
	return wordSelect
}

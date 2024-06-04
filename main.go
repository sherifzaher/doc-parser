package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	var filename string
	var wordToSearch string

	flag.StringVar(&filename, "filename", "", "file to search in it.")
	flag.StringVar(&wordToSearch, "word", "", "insert a word to search.")
	flag.Parse()

	if filename == "" {
		fmt.Println("Please insert the filename to search within")
		return
	}

	if wordToSearch == "" {
		fmt.Println("Please insert the word to search")
		return
	}

	fmt.Printf("Searching for %s in file %s\n", wordToSearch, filename)

	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening file %s: %s\n", filename, err)
	}
	defer file.Close()

	linesHaveTheWord, numberOfOccurrences := searchWithinFile(file, wordToSearch)
	if numberOfOccurrences == 0 {
		fmt.Printf("No occurrences found for %s in file %s\n", wordToSearch, filename)
		return
	}

	for line, text := range linesHaveTheWord {
		fmt.Printf("Line %d: %s\n", line, text)
	}

	fmt.Println("Number of occurrences:", numberOfOccurrences)
}

func searchWithinFile(file io.Reader, wordToSearch string) (map[int]string, int) {
	var linesHaveTheWord = make(map[int]string)
	var lineNumber = 0
	var numberOfOccurrences = 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lineNumber++
		line := scanner.Text()
		if strings.Contains(strings.ToLower(line), strings.ToLower(wordToSearch)) {
			numberOfOccurrences++
			linesHaveTheWord[lineNumber] = line
		}
	}

	return linesHaveTheWord, numberOfOccurrences
}

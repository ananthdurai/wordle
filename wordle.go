package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

const OutputDir = "/Users/ananthpackkildurai/data"

// A little Wordle utility to save my time. The program can filter any occurrence of a given letter or
// filter a letter based on its position
func main() {
	fmt.Println("Wordle is cooking")
	argsLength := len(os.Args)

	if argsLength < 3 {
		fmt.Println("The input file name and the letter to remove required")
		return
	}

	file := os.Args[1]
	wordList := parseWords(file)
	character := os.Args[2]
	var filterList []string
	if len(os.Args) > 3 {
		// Position is zero indexed
		position, _ := strconv.Atoi(os.Args[3])
		filterList = MatchLetterWithPosition(wordList, character, position)
	} else {
		filterList = MatchLetterWithAny(wordList, character)
	}
	Write(filterList)
	fmt.Println("Done")
}

func Write(words []string) {

	file, err := os.OpenFile(OutputFileWithTimestamp(), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	writer := bufio.NewWriter(file)

	for _, data := range words {
		_, _ = writer.WriteString(data + "\n")
	}
	err = writer.Flush()
	if err != nil {
		log.Fatal(err)
	}
	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func OutputFileWithTimestamp() string {
	return fmt.Sprintf("%s/%d.txt", OutputDir, time.Now().Unix())
}

func MatchLetterWithPosition(wordList []string, letter string, position int) []string {
	var filterList []string
	for _, word := range wordList {
		if strings.Index(word, letter) != position {
			filterList = append(filterList, word)
		}
	}
	return filterList
}

func MatchLetterWithAny(wordList []string, letter string) []string {
	var filterList []string
	for _, word := range wordList {
		if strings.Index(word, letter) < 0 {
			filterList = append(filterList, word)
		}
	}
	return filterList
}

func parseWords(file string) []string {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(f)

	scanner := bufio.NewScanner(f)

	var wordList []string

	for scanner.Scan() {
		words := strings.Fields(scanner.Text())
		wordList = append(wordList, words...)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return wordList
}

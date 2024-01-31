package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

func main() {

	// Prompt the user to enter the filename
	fmt.Println("Welcome!!! : ")
	fmt.Println("Please enter the name of your file to proceed: ")
	var filename string
	fmt.Scanln(&filename)

	// file reader
	inputFile, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal("help file error:", err)
	}

	// splitting per new line
	lines := strings.Split(string(inputFile), "\n")

	outMap := make(map[string]int)

	for _, spliceString := range lines {
		// calculating word count
		outMap = countWord(spliceString, outMap)
	}

	// printing final output
	reportResults(outMap)

}

/*
this function is taking the string as the input it also has a map which is keeping the
record of count and also it return each word count
*/
func countWord(lineValue string, wordCount map[string]int) (wordCounts map[string]int) {

	words := strings.Split(removePunctuation(lineValue), " ")

	for _, word := range words {
		_, exists := wordCount[word] // this
		if exists {
			wordCount[word] += 1
		} else {
			wordCount[word] = 1
		}
	}

	return wordCount
}

/*
removing the punctuations
*/
func removePunctuation(inputString string) (outputString string) {
	reg, err := regexp.Compile("[^a-z A-Z 0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	outputString = reg.ReplaceAllString(inputString, "")
	return outputString
}

/*
printing the map
*/
func reportResults(wordCount map[string]int) {
	for k, v := range wordCount {
		fmt.Println(k, v)
	}

}

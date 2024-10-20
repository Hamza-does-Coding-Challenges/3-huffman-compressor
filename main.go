package main

import (
	"fmt"
	"os"
	"strings"
)

// Step 1
func getCharacterCount(fileContent []byte) map[string]int {
	characterMap := make(map[string]int)

	stringContent := string(fileContent)

	stringContent = strings.ReplaceAll(stringContent, "\n", "")
	stringContent = strings.ReplaceAll(stringContent, "\r", "")
	stringContent = strings.ReplaceAll(stringContent, "\t", "")
	stringContent = strings.ReplaceAll(stringContent, " ", "")
	
	charactersArray := strings.Split(stringContent, "")

	for i := 0; i < len(charactersArray); i++ {
		character := charactersArray[i]

		_, ok := characterMap[character]

		if (ok) {
			characterMap[character]++
		} else {
			characterMap[character] = 1
		}
	}

	val := characterMap["X"]
	fmt.Println("X count: ", val)

	return characterMap
}

func readFile(filename string) ([]byte, error) {
	file, err := os.ReadFile(filename)

	if (err != nil) {
		fmt.Println("Error reading file: ", err);

		return nil, err
	}

	return file, nil;
}



func main() {

	// Args[0] is main.go, 2nd arg has to be the file to be read
	if (len(os.Args) != 2) {
		fmt.Println("Please provide a single file to read")
		return
	}

	file := os.Args[1]

	fileContent, err := readFile(file)

	if (err != nil) {
		fmt.Println("Error reading file: ", err)
		return
	}

	characterCountMap := getCharacterCount(fileContent)

	fmt.Println("Character count: ", characterCountMap)

}
package functions

import (
	"fmt"
	"os"
	"strings"
)

func ReadAsciiBanner(filname string) map[rune][]string {
	file, err := os.ReadFile(filname)
	if err != nil {
		fmt.Println("Error when opening the file!")
		return nil
	}
	str := string(file)
	input := strings.Split(str, "\n\n")
	if len(input[0]) != 8 {
		input[0] = input[0][1:]
	}

	asciiMap := make(map[rune][]string)
	char := 32
	for _, value := range input {
		asciiMap[rune(char)] = strings.Split(value, "\n")
		char++
	}
	return asciiMap
}

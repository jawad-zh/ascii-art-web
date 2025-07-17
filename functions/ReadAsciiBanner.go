package functions

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func ReadAsciiBanner(filname string) (map[rune][]string, error) {
	file, err := os.ReadFile(filname)
	if len(file) == 0 {
		fmt.Println("Error in Banner file!")
		return nil, errors.New("error in Banner file")
	}
	if err != nil {
		fmt.Println("Error when opening the file!")
		return nil, err
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
	return asciiMap, nil
}

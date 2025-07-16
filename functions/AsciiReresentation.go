package functions

import (
	"fmt"
	"strings"
)

func AsciiRepresentation(str string, asciiMap map[rune][]string) string {
	stro := ""
	words := strings.Split(str, "\\n")
	fmt.Println(words)
	if words[0] == "" {
		words = words[1:]
	}
	slice := [][]string{}
	for _, word := range words {
		for _, char := range word {
			for key, value := range asciiMap {
				if char == key {
					slice = append(slice, value)
				}
			}
		}
		stro += Print(slice)
		slice = nil
	}
	return stro
}

func Print(slice [][]string) string {
	str := ""
	for i := 0; i <= 7; i++ {
		lineSlice := []string{}
		for _, char := range slice {
			lineSlice = append(lineSlice, char[i])
		}
		strPrint := strings.Join(lineSlice, " ")
		
			str += "\n"
		
		str += strPrint
	}
	return str
}

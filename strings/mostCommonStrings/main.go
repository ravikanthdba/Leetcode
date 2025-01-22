package main

import (
	"strings"
	"unicode"
	"math"
	"fmt"
)

func contains(bannedWords []string, word string) bool {
    for _, value := range bannedWords {
        if strings.ToLower(value) == strings.ToLower(word) {
            return true
        }
    }
    return false
}

func removePunctuation(input string) string {
    result := []rune{} // Using rune slice to handle Unicode characters
    for i, char := range input {
        // Check if the character is alphanumeric or a space
        if unicode.IsLetter(char) || unicode.IsDigit(char) || unicode.IsSpace(char) {
            result = append(result, char)
        } else if char == '.' && i > 0 && i < len(input)-1 { // Allow decimal points in numbers
            // Check if the previous and next characters are digits
            if unicode.IsDigit(rune(input[i-1])) && unicode.IsDigit(rune(input[i+1])) {
                result = append(result, char)
            }
        }
    }
    return string(result) // Convert rune slice back to string
}

func mostCommonWord(paragraph string, banned []string) string {
    if len(paragraph) == 0 {
        return ""
    }

    var commonWords = make(map[string]int)
	fmt.Println("para without punctuation", removePunctuation(paragraph))

    wordsList := strings.Split(paragraph, " ")
	fmt.Println("words list = ", wordsList)

    for _, value := range wordsList {
        punctuationRemovedWord := strings.Trim(removePunctuation(value), "")
		fmt.Println(value, punctuationRemovedWord)
        if !contains(banned, punctuationRemovedWord) {
                commonWords[strings.ToLower(punctuationRemovedWord)]++
        }
    }
	fmt.Println(commonWords)
	


    maxWords := math.MinInt
    var word string
    for key, value := range commonWords {
        if value > maxWords {
            maxWords = value
            word = key
        }
    }

    return word
}

func main() {
	fmt.Println(mostCommonWord("a, a, a, a, b,b,b,c, c", []string{"a"}))
}
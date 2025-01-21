package main

import (
	"fmt"
)

func PhoneNumberMnemonics(phoneNumber string) []string {
	// Write your code here.
	if phoneNumber == "" {
		return []string{}
	}

	var numberCharacterMap = make(map[string][]string)

	numberCharacterMap["0"] = []string{"0"}
	numberCharacterMap["1"] = []string{"1"}
	numberCharacterMap["2"] = []string{"a", "b", "c"}
	numberCharacterMap["3"] = []string{"d", "e", "f"}
	numberCharacterMap["4"] = []string{"g", "h", "i"}
	numberCharacterMap["5"] = []string{"j", "k", "l"}
	numberCharacterMap["6"] = []string{"m", "n", "o"}
	numberCharacterMap["7"] = []string{"p", "q", "r", "s"}
	numberCharacterMap["8"] = []string{"t", "u", "v"}
	numberCharacterMap["9"] = []string{"w", "x", "y", "z"}

	if len(phoneNumber) == 1 {
		return numberCharacterMap[phoneNumber]
	}

	firstNumber := string(phoneNumber[0])
	remainingNumber := phoneNumber[1:]
	lettersFirstNumber := numberCharacterMap[firstNumber]
	subMnemonic := PhoneNumberMnemonics(remainingNumber)

	var result []string
	for _, letter := range lettersFirstNumber {
		for _, mnemonic := range subMnemonic {
			result = append(result, letter+mnemonic)
		}
	}

	return result
}

func main() {
	fmt.Println(PhoneNumberMnemonics("1905"))
}

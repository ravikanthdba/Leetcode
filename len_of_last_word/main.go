package main

import (
	"fmt"
	"strings"
)

func lengthOfLastWord(s string) int {
	if len(s) == 0 {
		return 0
	}

	list := strings.Split(s, " ")
	fmt.Println(list, len(list))

	var listPostProcessing []string

	for i := 0; i < len(list); i++ {
		if list[i] != "" {
			listPostProcessing = append(listPostProcessing, list[i])
			fmt.Println(listPostProcessing, len(listPostProcessing))
		}
	}

	return len(listPostProcessing[len(listPostProcessing)-1])

}

func main() {
	s := "   fly me   to   the moon  "
	println(lengthOfLastWord(s))
}

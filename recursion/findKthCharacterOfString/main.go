package main

import "fmt"

// Recursive helper function to find the k-th character
func kthCharacter(k int) byte {
	return findKthChar(k, 1, 'a') // Initial word "a" of length 1
}

// Recursive helper function to find the k-th character
func findKthChar(k, currentLength int, currentChar byte) byte {
	if k == 1 {
		return currentChar // Base case: The first character is the result
	}

	// Calculate the length of the word after the operation (doubles each time)
	newLength := currentLength * 2

	// If k is in the first half, recurse into the first half (which is the same as previous)
	if k <= newLength/2 {
		return findKthChar(k, currentLength, currentChar+1) // Move to next character
	}

	// If k is in the second half, recurse into the second half (adjust the index)
	return findKthChar(k-newLength/2, currentLength, currentChar+1)
}

func main() {
	fmt.Println(string(kthCharacter(5)))
}

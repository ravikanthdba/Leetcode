package main

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 || len(strs) == 1 {
		return ""
	}

	var lengthStringMap = make(map[int]string)

	for _, value := range strs {
		lengthStringMap[len(value)] = value
	}

	leastWord := 0

	for key := range lengthStringMap {
		if key < leastWord {
			leastWord = key
		} else {
			continue
		}
	}

	return result
}

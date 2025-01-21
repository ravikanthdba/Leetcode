package main

import (
	"fmt"
	"strings"
)

func reverseStr(s string, k int) string {
	if len(s) == 0 {
		return ""
	}

	if len(s) == 1 {
		return s
	}

	if len(s) == 2 {
		if k >= len(s) {
			return string(s[1]) + string(s[0])
		}
		return s
	}

	if k == 0 || k == 1 {
		return s
	}

	if k > len(s) {
		return s
	}

	strList := strings.Split(s, "")

	if len(strList) <= k {
		left := 0
		right := len(strList) - 1
		for left < right {
			strList[left], strList[right] = strList[right], strList[left]
			left++
			right--
		}

	} else if len(strList) > k && len(strList) <= (k*2) {
		left := 0
		right := k - 1
		for left < right {
			strList[left], strList[right] = strList[right], strList[left]
			left++
			right--
		}
	} else if len(strList) > (k * 2) {
		left := 0
		right := k - 1
		for left < right {
			strList[left], strList[right] = strList[right], strList[left]
			left++
			right--
		}

		i := (k * 2) - 1
		for i < len(strList) {
			left := i
			right := len(strList) - 1
			for left < right {
				if left%2 == 0 {
					strList[left], strList[right] = strList[right], strList[left]
					left++
					right--
				} else {
					left++
					right--
				}
			}
			i *= 2
		}
	}

	return strings.Join(strList, "")
}

func main() {
	fmt.Println(reverseStr("abcdefg", 2))
	fmt.Println(reverseStr("abcd", 2))
	fmt.Println(reverseStr("hyzqyljrnigxvdtneasepfahmtyhlohwxmkqcdfehybknvdmfrfvtbsovjbdhevlfxpdaovjgunjqlimjkfnqcqnajmebeddqsgl", 39))

}

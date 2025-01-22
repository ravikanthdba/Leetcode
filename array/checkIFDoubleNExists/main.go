package main

import (
	"fmt"
	"sort"
)

func checkIfExist(arr []int) bool {
    if len(arr) == 0 || len(arr) == 1 {
        return false
    }

    if len(arr) == 2 {
        if (arr[0] * 2) == arr[1] {
            return true
        }
        return false
    }

    sort.Ints(arr)

    ptr := 0
    var i int
    for ptr < len(arr) - 1 {
        i = ptr + 1
        for arr[i] <= arr[ptr] * 2 && i <= len(arr)-1 {
            if arr[ptr] * 2 == arr[i] {
                return true
            } else {
                if i == len(arr) - 1 {
                    return false
                } else {
                    i++
                }
            }
        }
        ptr++
    }

    return false
}

func main() {
	fmt.Println(checkIfExist([]int{3,1,7,11}))
}
package main

import (
	"fmt"
	"math"
)

type Array struct {
	array  [20]int
	size   int
	length int
}

func (a *Array) Add(element int) {
	a.array[a.length] = element
	a.length++
}

func (a *Array) Display() {
	for i := 0; i < a.length; i++ {
		fmt.Println(a.array[i])
	}
}

func (a *Array) Insert(position, element int) {
	for i := a.length - 1; i > position; i-- {
		a.array[i] = a.array[i-1]
	}
	a.array[position] = element
	a.length++
}

func (a *Array) Delete(position int) {
	if a.length == 0 || position > a.length {
		fmt.Println("nothing to delete")
	}

	x := a.array[position]
	for i := position; i < a.length; i++ {
		a.array[i] = a.array[i+1]
	}
	a.length--

	fmt.Println("deleted element is: ", x)
}

func (a *Array) Search(element int) {
	if a.length == 0 {
		fmt.Println("array is empty")
		return
	}

	if a.length == 1 {
		if a.array[0] == element {
			fmt.Println("element: ", element, " found at position: ", 1)
		} else {
			fmt.Println("element not found in the list")
		}
		return
	}

	for key, value := range a.array {
		if value == element {
			fmt.Println("element found at position: ", key)
			return
		}
	}

	fmt.Println("element not found")
}

func (a *Array) SearchWithMoveAtHead(element int) {
	if a.length == 0 {
		fmt.Println("array is empty")
		return
	}

	if a.length == 1 {
		if a.array[0] == element {
			fmt.Println("element: ", element, " found at position: ", 1)
		} else {
			fmt.Println("element not found in the list")
		}
		return
	}

	for key, value := range a.array {
		if value == element {
			fmt.Println("element found at position: ", key, "swapping")
			a.array[0], a.array[key] = a.array[key], a.array[0]
			return
		}
	}

	fmt.Println("element not found")
}

func (a *Array) SearchWithBinarySearch(element int) {
	if a.length == 0 {
		fmt.Println("array is empty")
		return
	}

	if a.length == 1 {
		if a.array[0] == element {
			fmt.Println("element: ", element, " found at position: ", 1)
		} else {
			fmt.Println("element not found in the list")
		}
		return
	}

	low := 0
	high := a.length

	for low <= high {
		mid := (low + high) / 2
		if a.array[mid] == element {
			fmt.Println("element: ", element, " returned at position: ", mid)
			return
		} else if element > a.array[mid] {
			low = mid + 1
		} else if element < a.array[mid] {
			high = mid - 1
		}
	}

	fmt.Println("element: ", element, " not found in the array")
}

func (a *Array) Get(position int) {
	if a.length == 0 {
		fmt.Println("array is empty")
		return
	}

	if position > a.length {
		fmt.Println("invalid position")
		return
	}

	fmt.Println("element at position: ", position, " is: ", a.array[position])
}

func (a *Array) Set(position, element int) {
	if a.length == 0 {
		fmt.Println("array is empty")
		return
	}

	if position > a.length {
		fmt.Println("invalid position")
		return
	}

	a.array[position] = element

	fmt.Println("element: ", element, " set at position: ", position, " is: ", a.array[position])
}

func (a *Array) Avg() {
	if a.length == 0 {
		fmt.Println("array is empty")
		return
	}

	if a.length == 1 {
		fmt.Println(a.array[0])
		return
	}

	sum := 0
	for _, value := range a.array {
		sum += value
	}

	fmt.Println("Average is: ", float64(sum)/float64(a.length))
}

func (a *Array) Max() {
	if a.length == 0 {
		fmt.Println("array is empty")
		return
	}

	if a.length == 1 {
		fmt.Println(a.array[0])
		return
	}

	maxValue := math.MinInt
	for _, value := range a.array {
		if value > maxValue {
			maxValue = value
		}
	}

	fmt.Println("Max element in the array is: ", maxValue)
}

func (a *Array) ReverseArrayUsingLinear() {
	var b [20]int
	j := 0
	for i := a.length - 1; i >= 0; i-- {
		b[j] = a.array[i]
		j++
	}

	for i := 0; i < a.length; i++ {
		a.array[i] = b[i]
	}

	fmt.Println("array reversed using linear approach")
}

func (a *Array) ReverseArrayUsingTwoPointerApproach() {
	if a.length == 0 || a.length == 1 {
		fmt.Println("nothing to reverse")
		return
	}

	if a.length == 2 {
		a.array[0], a.array[1] = a.array[1], a.array[0]
		return
	}

	p := 0
	q := a.length - 1

	for p < q {
		a.array[p], a.array[q] = a.array[q], a.array[p]
		p++
		q--
	}

	fmt.Println("array reversed using two pointer approach")
}

func (a *Array) LeftShift() {
	if a.length == 0 || a.length == 1 {
		fmt.Println("nothing to reverse")
		return
	}

	if a.length == 2 {
		a.array[0], a.array[1] = a.array[1], a.array[0]
		return
	}

	for i := 0; i < a.length; i++ {
		a.array[i] = a.array[i+1]
	}
	//a.length--

	fmt.Println("array left shifting completed")
}

func (a *Array) LeftRotate() {
	if a.length == 0 || a.length == 1 {
		fmt.Println("nothing to reverse")
		return
	}

	if a.length == 2 {
		a.array[0], a.array[1] = a.array[1], a.array[0]
		return
	}

	element := a.array[0]
	var i int
	for i = 0; i < a.length-1; i++ {
		a.array[i] = a.array[i+1]
	}
	a.array[i] = element
	fmt.Println("array left rotating completed")
}

func (a *Array) isSorted() {
	if a.length == 0 || a.length == 1 {
		fmt.Println("can't check if its sorted")
		return
	}

	if a.length == 2 {
		if a.array[0] < a.array[1] {
			fmt.Println("sorted")
		} else {
			fmt.Println("not sorted")
		}
		return
	}

	p := 0
	for p < a.length-1 {
		if a.array[p] > a.array[p+1] {
			fmt.Println("not sorted")
			return
		}
		p++
	}

	fmt.Println("sorted")
	return
}

func main() {
	var a Array
	//for i := 0; i < 10; i++ {
	//	a.Add(rand.Intn(10))
	//}
	////a.Display()
	////fmt.Println("====================")
	////a.Insert(5, 20)
	////a.Display()
	//
	//a.Insert(10, 20)
	//a.Display()
	//fmt.Println("====================")
	//fmt.Println("====================")
	//a.Delete(5)
	//a.Display()
	//a.Delete(9)
	//a.Display()
	//
	//a.SearchWithMoveAtHead(1)
	//a.Display()
	//a.Search(1)

	a.length = 10
	a.array = [20]int{8, 9, 10, 15, 20, 22, 30, 31, 60, 100}
	a.SearchWithBinarySearch(20)
	a.SearchWithBinarySearch(100)
	a.SearchWithBinarySearch(110)
	a.Get(5)
	a.Display()
	a.Set(5, 25)
	a.Display()
	a.Avg()
	a.Max()
	fmt.Println("**************************************")

	a.Display()
	a.ReverseArrayUsingLinear()
	a.Display()

	fmt.Println("**************************************")
	fmt.Println("**************************************")

	a.Display()
	a.ReverseArrayUsingTwoPointerApproach()
	a.Display()

	fmt.Println("**************************************")
	fmt.Println("**************************************")

	a.Display()
	a.LeftShift()
	a.Display()

	fmt.Println("**************************************")
	fmt.Println("**************************************")

	a.Display()
	a.LeftRotate()
	a.Display()

	fmt.Println("**************************************")
	fmt.Println("**************************************")

	a.isSorted()
	a.array = [20]int{8, 9, 10, 15, 20, 22, 30, 31, 60, 100}
	a.isSorted()
}

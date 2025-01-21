package main

import "fmt"

type MovingAverage struct {
	size   int
	front  int
	rear   int
	values [][]int
}

func Constructor(size int) MovingAverage {
	if size <= 0 {
		return MovingAverage{}
	}

	var array [][]int
	return MovingAverage{
		size:   size,
		front:  -1,
		rear:   -1,
		values: array,
	}
}

func (this *MovingAverage) Next(val int) float64 {
	var avg float64

	if len(this.values) < this.size {
		this.rear++
		this.values = append(this.values, []int{val})
	} else {
		this.values = this.values[1:]
		this.values = append(this.values, []int{val})
	}

	if this.front == -1 && this.rear == 0 {
		return float64(this.values[this.rear][0])
	}

	this.front = -1
	for this.front < this.rear {
		this.front++
		element := this.values[this.front][0]
		avg = avg + float64(element)
	}

	return avg / float64(len(this.values))
}

func main() {
	movingAverage := Constructor(3)
	fmt.Println(movingAverage.Next(1))
	fmt.Println(movingAverage.Next(10))
	fmt.Println(movingAverage.Next(3))
	fmt.Println(movingAverage.Next(5))
}

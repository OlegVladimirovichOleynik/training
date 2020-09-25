package main

import (
	"fmt"
	"golang-book/chapter11/math"
)

func main() {
	xs := []float64{1,2,3,4,5}
	avg := math.Average(xs)
	fmt.Println("Среднее значение", avg)
	mini := math.Min(xs)
	fmt.Println("Минимальное значение", mini)
	maxi := math.Max(xs)
	fmt.Println("Макимальное значени", maxi)
}

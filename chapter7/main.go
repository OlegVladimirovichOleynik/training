package main

import (
	"fmt"
)

func main() {
	xs := []int64{98, 93, 77, 82, 83}
	fmt.Println(sum(xs))
	chetNechet(5)
	chetNechet(6)
	fmt.Println("максимальный элемент", findMax(12, 11, 5, 122, 3, 5, 6))
	var nextOdd = makeOddGenerator()
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
	fmt.Println(fib(3))
	fmt.Println(factorial(4))
}

func sum(arr []int64) int64 {
	var result int64
	for _, i := range arr {
		result += i
	}
	return result
}

func chetNechet(x int64) bool {
	if x%2 == 0 {
		fmt.Println("значение х ", x, "четное")
		return true
	}
	fmt.Println("значение х ", x, "нечетное")
	return false
}

func findMax(arrgs ...int) int {
	var maxElement = 0
	for _, i := range arrgs {
		if maxElement < i {
			maxElement = i
		}
	}
	return maxElement
}

func makeOddGenerator() func() int64 {
	i := int64(1)
	return func() (ret int64) {
		ret = i
		i += 2
		return
	}
}

func fib(n int) []int {
	result := make([]int, n)
	for i := 0; i < n; i++ {
		if i < 2 {
			result[i] = i + 1
		} else {
			result[i] = result[i-1] + result[i-2]
		}
	}
	return result
}

func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

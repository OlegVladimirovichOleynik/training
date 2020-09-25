package main

import "fmt"

func main() {
	x := 1
	y := 2
	fmt.Println("x =", x, "y =", y)
	swap(&x, &y)
	fmt.Println("x =", x, "y =", y)
	z := 1.5
	square(&z)
	fmt.Println(z)
}

func swap(x, y *int) {
	*x, *y = *y, *x
}

func square(z *float64) {
	*z = *z * *z
}

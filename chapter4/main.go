package main

import "fmt"

func main() {
    fmt.Print("Enter yor temperature: ")
    var input float64
    fmt.Scanf("%f", &input)

    output := (input - 32) * 5/9

    fmt.Println(output)
}

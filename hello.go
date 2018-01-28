package main

import "fmt"

func Sqrt(x float64) float64 {
    guess := 1.0
    for i := 0; i < 10; i++ {
        guess -= (guess*guess - x) / (2*guess)
    }
    return guess
}

func main() {
    fmt.Printf("Hello, world. \n");
    for i := 1; i <= 10; i++ {
        for j := 1; j <= i; j ++ {
            fmt.Printf("%v %v ", j, Sqrt(float64(j)))
        }
        fmt.Println()
    }
}

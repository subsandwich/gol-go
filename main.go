package main

import "fmt"

func main() {

	arr := [100][100]bool{{true, true, true}}
	fmt.Println(arr)
	fmt.Println(StandardGameOfLifeRules(arr, 1, 99))
    fmt.Println("Hello, world.")
}
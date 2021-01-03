package main

import (
	"fmt"
)

func main() {
	fmt.Println("gony main started")

	zeros := Zeros(5,3)
	PrintMatrix(zeros, "zeros")


	fmt.Println("gony main finished")
}


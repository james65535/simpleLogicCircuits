package main

import (
	"github.com/james65535/simpleLogicCircuits/arith"
	"fmt"
)

func main() {
	num1 := []int{0, 0, 0, 1}
	num2 := []int{0, 1, 1, 1}
	total, err := arith.ByteAdder(num1, num2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Final byte: %d\n", total)
	}
}

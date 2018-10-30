package arith

import (
	"github.com/james65535/simpleLogicCircuits/logicgates"
	"fmt"
)

func HAdder(a, b int) (int, int) {
	return logicgates.Xor(a, b), logicgates.And(a, b)
}

func Adder(a, b, c int) (int, int) {
	s1, c1 := HAdder(a, b)
	s2, c2 := HAdder(c, s1)
	return s2, logicgates.Or(c2, c1)
}


func ByteAdder(a []int, b []int) ([]int, error) {
	carry := 0
	for i := len(a) - 1; i >= 0; i-- {
		a[i], carry = Adder(a[i], b[i], carry)
	}
	if carry != 0 {
		err := fmt.Errorf("Attempted to Exceed Max Byte")
		return nil, err
	} else {
		return a, nil
	}
}

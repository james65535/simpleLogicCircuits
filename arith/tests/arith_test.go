package arith_tests

import (
	"github.com/james65535/simpleLogicCircuits/arith"
	"testing"
)

func TestHAdder(t *testing.T) {
	tables := []struct {
		a int
		b int
		s int
		c int
	}{
		{0, 0, 0, 0},
		{0, 1, 1, 0},
		{1, 0, 1, 0},
		{1, 1, 0, 1},
	}

	for _, table := range tables {
		sum, carry := arith.HAdder(table.a, table.b)
		if sum != table.s || carry != table.c {
			t.Errorf("HAdder of (%d+%d) was incorrect, got: sum %d and carry %d, want: %d and %d.", table.a, table.b, sum, carry, table.s, table.c)
		}
		if carry != 0 {
		}
	}
}

func TestAdder(t *testing.T) {
	tables := []struct {
		a    int
		b    int
		cin  int
		cout int
		s    int
	}{
		{0, 0, 0, 0, 0},
		{0, 0, 1, 0, 1},
		{0, 1, 0, 0, 1},
		{0, 1, 1, 1, 0},
		{1, 0, 0, 0, 1},
		{1, 0, 1, 1, 0},
		{1, 1, 0, 1, 0},
		{1, 1, 1, 1, 1},
	}

	for _, table := range tables {
		sum, carry := arith.Adder(table.a, table.b, table.cin)
		if sum != table.s || carry != table.cout {
			t.Errorf("Adder of (%d+%d,%d) was incorrect, got: carry %d and sum %d, want: %d and %d.", table.a, table.b, table.cin, carry, sum, table.cout, table.s)
		}
		if carry != 0 {
		}
	}

}


func TestByteAdder(t *testing.T) {
	tables := []struct {
		a []int
		b []int
		s []int
	}{
		{[]int{0, 0, 0, 0}, []int{0, 0, 0, 0}, []int{0, 0, 0, 0}},
		{[]int{0, 0, 0, 1}, []int{0, 0, 0, 0}, []int{0, 0, 0, 1}},
		{[]int{0, 0, 0, 1}, []int{0, 0, 0, 1}, []int{0, 0, 1, 0}},
	}

	for _, table := range tables {
		total, err := arith.ByteAdder(table.a, table.b)
		for i, _ := range total {
			if total[i] != table.s[i] && err == nil {
				t.Errorf("ByteAdder of (%d,%d) was incorrect, got: %d, want: %d.", table.a, table.b, total, table.s)
				return
			}
		}
	}
}

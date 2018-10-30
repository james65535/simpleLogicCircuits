package tests

import (
	"github.com/james65535/simpleLogicCircuits/logicgates"
	"testing"
)

func TestNand(t *testing.T) {
	tables := []struct {
		a int
		b int
		q int
	}{
		{0, 0, 1},
		{0, 1, 1},
		{1, 0, 1},
		{1, 1, 0},
	}

	for _, table := range tables {
		total := logicgates.Nand(table.a, table.b)
		if total != table.q {
			t.Errorf("Nand of (%d,%d) was incorrect, got: %d, want: %d.", table.a, table.b, total, table.q)
		}
	}
}

func TestAnd(t *testing.T) {
	tables := []struct {
		a int
		b int
		q int
	}{
		{0, 0, 0},
		{0, 1, 0},
		{1, 0, 0},
		{1, 1, 1},
	}

	for _, table := range tables {
		total := logicgates.And(table.a, table.b)
		if total != table.q {
			t.Errorf("And of (%d,%d) was incorrect, got: %d, want: %d.", table.a, table.b, total, table.q)
		}
	}
}

func TestNot(t *testing.T) {
	tables := []struct {
		a int
		q int
	}{
		{0, 1},
		{1, 0},
	}

	for _, table := range tables {
		total := logicgates.Not(table.a)
		if total != table.q {
			t.Errorf("Not of (%d) was incorrect, got: %d, want: %d.", table.a, total, table.q)
		}
	}
}

func TestOr(t *testing.T) {
	tables := []struct {
		a int
		b int
		q int
	}{
		{0, 0, 0},
		{0, 1, 1},
		{1, 0, 1},
		{1, 1, 1},
	}

	for _, table := range tables {
		total := logicgates.Or(table.a, table.b)
		if total != table.q {
			t.Errorf("Or of (%d,%d) was incorrect, got: %d, want: %d.", table.a, table.b, total, table.q)
		}
	}
}

func TestNor(t *testing.T) {
	tables := []struct {
		a int
		b int
		q int
	}{
		{0, 0, 1},
		{0, 1, 0},
		{1, 0, 0},
		{1, 1, 0},
	}

	for _, table := range tables {
		total := logicgates.Nor(table.a, table.b)
		if total != table.q {
			t.Errorf(" nor of (%d,%d) was incorrect, got: %d, want: %d.", table.a, table.b, total, table.q)
		}
	}
}

func TestXor(t *testing.T) {
	tables := []struct {
		a int
		b int
		q int
	}{
		{0, 0, 0},
		{0, 1, 1},
		{1, 0, 1},
		{1, 1, 0},
	}

	for _, table := range tables {
		total := logicgates.Xor(table.a, table.b)
		if total != table.q {
			t.Errorf("xor of (%d,%d) was incorrect, got: %d, want: %d.", table.a, table.b, total, table.q)
		}
	}
}
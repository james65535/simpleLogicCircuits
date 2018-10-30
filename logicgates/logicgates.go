package logicgates

func Nand(a, b int) int {
	if a == 1 && b == 1 {
		return 0
	} else {
		return 1
	}
}

func And(a, b int) int {
	ab := Nand(a, b)
	return Nand(ab, ab)
}

func Not(a int) int {
	return Nand(a, a)
}

func Or(a, b int) int {
	return Nand(Nand(a, a), Nand(b, b))
}

func Nor(a, b int) int {
	return Not(Or(a, b))
}

func Xor(a, b int) int {
	if a != b {
		return 1
	} else {
		return 0
	}
}
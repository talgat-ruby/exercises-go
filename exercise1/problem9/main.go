package main

// Will do without using special characters &, |

func bitwiseAND(x, y int) int {
	a := binaryINT(x)
	b := binaryINT(y)
	//a, b := x, y
	res := 0
	mult := 1
	for a > 0 || b > 0 {
		modA := a % 10
		modB := b % 10
		if modA == 1 && modB == 1 {
			res += mult
		}
		mult *= 10
		a /= 10
		b /= 10
	}
	return binaryToDecimal(res)
}

func bitwiseOR(x, y int) int {
	a := binaryINT(x)
	b := binaryINT(y)
	//a, b := x, y
	res := 0
	mult := 1
	for a > 0 || b > 0 {
		modA := a % 10
		modB := b % 10
		if modA == 1 || modB == 1 {
			res += mult
		}
		mult *= 10
		a /= 10
		b /= 10
	}
	return binaryToDecimal(res)
}

func bitwiseXOR(x, y int) int {
	a := binaryINT(x)
	b := binaryINT(y)
	//a, b := x, y
	res := 0
	mult := 1
	for a > 0 || b > 0 {
		modA := a % 10
		modB := b % 10
		if modA != modB {
			res += mult
		}
		mult *= 10
		a /= 10
		b /= 10
	}
	return binaryToDecimal(res)
}

// will use function from problem2 (changed for int)
func binaryINT(n int) int {
	if n == 0 {
		return 0
	}

	res := 0
	mult := 1
	a := n

	for a > 0 {
		mod := a % 2
		res += mod * mult
		mult *= 10
		a /= 2
	}
	return res
}

func binaryToDecimal(n int) int {
	res := 0
	mult := 1

	for n > 0 {
		bit := n % 10
		res += bit * mult
		mult *= 2
		n /= 10
	}
	return res
}

/*
func bitwiseAND(a, b int) int {
	return a & b
}

func bitwiseOR(a, b int) int {
	return a | b
}

func bitwiseXOR(a, b int) int {
	return a ^ b
}
*/

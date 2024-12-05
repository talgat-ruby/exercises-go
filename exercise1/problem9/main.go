package main

func bitwiseAND(a int, b int) int {
	return a & b
}

func bitwiseOR(a int, b int) int {
	return a | b
}

func bitwiseXOR(a int, b int) int {
	return a ^ b
}

// func bitwise(a int) string {
// 	var result string
// 	for i := 0; i < 8; i++ {
// 		if a > 0 {
// 			result = fmt.Sprintf("%v", a%2) + result
// 			a = a / 2
// 		} else {
// 			result = "0" + result
// 		}
// 	}
// 	return result
// }

package main

func binary(n int) string {
	if n == 0 {
		return "0"
	}

	var result []byte
	for n > 0 {
		result = append(result, byte('0'+n%2))
		n /= 2
	}

	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return string(result)
}

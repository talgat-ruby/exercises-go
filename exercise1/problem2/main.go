package main

func binary(num int) string {
	if num == 0 {
		return "0"
	}
	out := make([]rune, 0)
	for i := 31; i >= 0; i-- {
		diff := num - 1<<i
		if diff < 0 && len(out) == 0 {
			continue
		} else if diff < 0 {
			out = append(out, '0')
		} else {
			out = append(out, '1')
			num = diff
		}
	}
	return string(out)
}

package main

func countVowels(s string) int {
	count:=0
	for _,el:= range s {
		if el=='a' || el=='A' || 
			el=='e' || el=='E' || 
			el=='i' || el=='I' || 
			el=='o' || el=='O' ||
			el=='U' || el=='u' {
				count++
		}
	}
	return count
}

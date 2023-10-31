package main

import "fmt"

func main() {

	a := 1
	b := 1
	found := false

	// 플래그 사용해 for문 한번에 빠져 나오기
	for ; a <= 9; a++ {
		for b = 1; b <= 9; b++ {
			if a*b == 45 {
				found = true
				break
			}
		}
		if found {
			break
		}
	}
	fmt.Printf("%d * %d = %d\n", a, b, a*b)
}

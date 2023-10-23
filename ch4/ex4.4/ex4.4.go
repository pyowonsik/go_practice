package main

import "fmt"

func main() {

	// 타입 변환
	aa := 3
	var bb float64 = 3.5

	var cc int = int(bb)
	dd := float64(aa * cc)
	var ee int64 = 7
	ff := int64(dd) * ee
	var gg int = int(bb * 3)
	var hh int = int(bb) * 3

	fmt.Println("타입 변환")
	fmt.Println(gg, hh, ff)
}

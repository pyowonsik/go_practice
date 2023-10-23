package main

import "fmt"

func main() {
	// Scan 함수와 Scanln 함수의 차이점은 Scan 함수는 공백이나 개행을 기준으로 변환을 하지만
	// Scanln은 공백만을 기준으로 변환하며 개행을 만나면 입력 완료로 인지하여 변환 후 결과를 반환합니다.
	// -> '개행시 입력완료로 인지'

	var a int
	var b int
	n, err := fmt.Scanln(&a, &b) // a,b 주소를 입력으로 넘겨야함
	if err != nil {
		fmt.Println(n, err)
	} else {
		fmt.Println(n, a, b)
	}
}

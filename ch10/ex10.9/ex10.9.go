package main

import "fmt"

func main() {
	a := 3
	switch a {
	case 1:
		fmt.Println("a == 1")
	case 2:
		fmt.Println("a == 2")
	case 3:
		fmt.Println("a == 3")
		fallthrough
	case 4:
		fmt.Println("a == 4")
		fallthrough
	default:
		fmt.Println("a > 4")
	}
}


switch 비굣값{
case 값1:	// 비굣값과 값1이 같을 때 수행합니다.
	문장
case 값2:	// 비굣값과 값2이 같을 때 수행합니다.
	문장
default:	// 만족하는 case가 없을 때 수행합니다.
	문장	// default 구문은 생략 가능합니다.
}


switch 비굣값{
	case 값1:	// 비굣값과 값1이 같을 때 수행합니다.
		문장
	case 값2:	// 비굣값과 값2이 같을 때 수행합니다.
		문장
	default:	// 만족하는 case가 없을 때 수행합니다.
		문장	// default 구문은 생략 가능합니다.
	}
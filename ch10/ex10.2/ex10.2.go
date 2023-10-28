package main

import "fmt"

func main() {

	day := 3

	switch day {
	case 1:
		fmt.Println("첫째날 입니다.\n오늘은 팀미팅이 있습니다.")
	case 2:
		fmt.Println("둘째날 입니다.\n새로운 팀원 면접이있습니다.")
	case 3:
		fmt.Println("셋째날 입니다.\n설계안을 확정하는 날입니다.")
	case 4:
		fmt.Println("넷째날 입니다.\n예산을 확정하는 날입니다.")
	case 5:
		fmt.Println("다섯째날 입니다.\n최종 계약하는 날입니다.")
	default:
		fmt.Println("프로젝트를 진행하세요")
	}
}

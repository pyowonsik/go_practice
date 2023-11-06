package main

import "fmt"

type Data struct {
	value int
	data  [200]int
}

func ChnageData(arg *Data) {
	arg.value = 999
	arg.data[100] = 999
}

func main() {

	var data Data

	// 주소값이 아니 그냥 데이터를 던져주면 함수 내부의 데이터만 변경되고
	// 리턴되는 값이 없기때문에 main에서는 데이터 변경이 되지 않는다.
	// 1608 바이트 복사

	// 하지만 주소값을 던져주게 되면
	// 함수 내부의 Data의 주소가 가리키는 값이 바뀌기 때문에 main()의 값도 바뀐다.
	// 8 바이트만 복사

	// &data  == arg
	// data ==  *arg
	ChnageData(&data)

	fmt.Printf("value : %d\n", data.value)
	fmt.Printf("data : %d\n", data.data[100])

	// var data Data
	// var p *Data = &data

	// var p *Data = &Data{}

	// p := &Data{}

	// var p := new(Data) -> 내부 필드값 초기화 불가 .

	// 인스턴스 생성이 아닌 가리키는것 : 인스턴스는 1개
	// var p1 *Data = &Data{}
	// var p2 *Data = p1
	// var p3 *Data = p1

}

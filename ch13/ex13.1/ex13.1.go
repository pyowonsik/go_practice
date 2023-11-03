package main

import "fmt"

type House struct {
	Adress string
	Size   int
	Price  float64
	Type   string
}

func main() {

	var house House = House{"서울시 서초구 ....", 8, 6.5, "오피스텔"}

	// house.Adress = "서울시 서초구 ...."
	// house.Size = 8
	// house.Price = 6.5
	// house.Type = "오피스텔"

	fmt.Println("주소 : ", house.Adress)
	fmt.Printf("크기 : %d평\n", house.Size)
	fmt.Println("가격 : ", house.Price)
	fmt.Println("타입 : ", house.Type)

}

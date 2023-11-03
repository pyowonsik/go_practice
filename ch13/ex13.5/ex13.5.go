package main

import "fmt"

type Student struct {
	Age   int
	No    int
	Scroe float64
}

func PrintStudent(student Student) {

	fmt.Printf("나이 : %d 번호 : %d 점수 : %.2f\n", student.Age, student.No, student.Scroe)

}

func main() {

	var student = Student{15, 23, 88.2}

	student2 := student

	PrintStudent(student2)

}

package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func updatePerson(p *Person, name string, age int) {
	p.Name = name
	p.Age = age
}

func main() {

	var person1 = Person{"Alice", 30}

	person2 := &Person{"Bob", 25}

	person2.Age = 26

	fmt.Printf("person1 : %s , Age : %d\n", person1.Name, person1.Age)
	fmt.Printf("person2 : %s , Age : %d\n", person2.Name, person2.Age)

	// 인스턴스는 하나 이기때문에 주소가 가리키고 있는 값이 변경되면 모두 변경이 되는것 !
	var person3 = &person1

	updatePerson(person3, "Pyo", 25)
	fmt.Printf("person1 : %s , Age : %d\n", person1.Name, person1.Age)
	fmt.Printf("person3 : %s , Age : %d\n", person3.Name, person3.Age)
	fmt.Printf("person2 : %s , Age : %d\n", person2.Name, person2.Age)
}

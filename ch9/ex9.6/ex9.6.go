package main

import "fmt"

func getMyAge() (int, bool) {
	return 33, true
}

func main() {

	// UploadFile() -> 멀티반환 함수
	// if filename,success := UploadFile(); success{
	// 	fmt.Println("Upload success",filename)
	// }else{
	// 	fmt.Println("Failed to upload")
	// }

	if age, ok := getMyAge(); ok && age < 20 {
		fmt.Println("You are young", age)
	} else if ok && age < 30 {
		fmt.Println("Nice age", age)
	} else if ok {
		fmt.Println("You are beautiful", age)
	} else {
		fmt.Println("Error")
	}
	// fmt.Println("Your age is", age) 소멸

}

package something

import "fmt"

// function의 시작이 소문자일 경우 private
func sayBye() {
	fmt.Println("Bye")
}

// function의 시작이 대문자일 경우 package로 사용이 가능해짐
func SayHello() {
	fmt.Println("Hello")
}

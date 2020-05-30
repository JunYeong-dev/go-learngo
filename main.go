package main

import (
	"fmt"

	"github.com/JunYeong-dev/learngo/something"
)

// Node.js, python과 다르게 특정 function을 찾게 되는데, 그게 바로 function main - func main()
// 컴파일러가 main package와 main function을 먼저 찾고 실행시킴
func main() {
	var x string = "Hello World"
	// Println과 같이 왜 함수의 시작이 대문자일까?
	fmt.Println(x)
	// 이처럼 시작이 대문자인 것들만 package로 사용이 가능
	something.SayHello()
}

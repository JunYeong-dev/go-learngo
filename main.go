package main

import (
	"fmt"
	"strings"

	"github.com/JunYeong-dev/learngo/something"
)

// Node.js, python과 다르게 특정 function을 찾게 되는데, 그게 바로 function main - func main()
// 컴파일러가 main package와 main function을 먼저 찾고 실행시킴
func main() {
	// Println과 같이 왜 함수의 시작이 대문자일까?
	fmt.Println("Hello World")
	// 이처럼 시작이 대문자인 것들만 package로 사용이 가능
	something.SayHello()

	// 상수(Constant)와 변수(Variables)
	// Go는 type언어 이기 때문에 type이 무엇인지 설정해 줘야 함(Java나 C언어와 같이)
	const name string = "nick"
	// 상수이기 때문에 변경 불가능
	// name = "judy" (x)

	// 변수기 때문에 변경 가능
	var a string = "nick"
	// a := "nick" <- 축약형; 이 경우에는 Go가 알아서 type을 찾아줌
	// 축약형은 오로지 function안에서만 가능하고, 변수에만 적용 가능
	a = "judy"
	fmt.Println(a)

	// 함수 사용
	fmt.Println(multiply(2, 3))

	// return값이 2개이기 때문에 2개의 변수로 받아줌
	// 만약 한개의 return값만 받고 싶은 경우
	// totalLength, _ := lenAndUpper("nick")
	// 위와 같이 _ (언더바)를 사용하면 컴파일러가 무시하게 됨
	totalLength, upperName := lenAndUpper("nick")
	fmt.Println(totalLength, upperName)

	repeatMe("nick", "judy", "sindy", "nana")
}

// Java와 Python등과 다르게 매개변수와 return값의 type을 설정해 줘야함
// 매개변수가 같은 type일 경우에는 (a, b int) 와 같이 쓸 수 있음
func multiply(a int, b int) int {
	return a * b
}

// Go의 함수가 다른 언어와 다른 점은 여러 개의 return값을 가질 수 있음
func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

// 다수의 매개변수를 받기 위해서 type앞에 ...을 붙여주면 됨
func repeatMe(words ...string) {
	fmt.Println(words)
}

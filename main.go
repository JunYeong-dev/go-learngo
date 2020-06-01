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

	fmt.Println(lenAndLower("JUDY"))

	fmt.Println(superAdd(1, 2, 3, 4, 5, 6))

	fmt.Println(canIDrink(18))
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

// 'naked' return
func lenAndLower(name string) (length int, lowercase string) {
	// 'defer' : function이 끝난 후에 실행되는 코드
	defer fmt.Println("Finish")
	length = len(name)
	lowercase = strings.ToLower(name)
	return
}

// range를 사용한 loop
func superAdd(numbers ...int) int {
	total := 0
	// 첫번째 값은 index지만 이 함수에서는 사용하지 않기 때문에 _ 로 ignore처리
	// range는 array에 loop를 적용할 수 있도록 해줌
	for _, number := range numbers {
		total += number
	}

	// 흔히 알고 있는 for문 형식으로도 loop를 만들 수 있음
	// for i := 0; i < len(numbers); i++ {
	// 	total += numbers[i]
	// }

	return total
}

// 조건문
func canIDrink(age int) bool {
	// 기본적인 조건문
	// if age < 18 {
	// 	return false
	// } else {
	// 	return true
	// }

	// Go에서는 if안에서 변수를 정의하고 사용할 수 있음
	// 물론 일반적 사용방법으로 변수를 따로 설정할 수 있지만
	// 이 조건문에서만 사용하기 때문에 코드를 읽는 사람이 이해하기 더 좋음
	// else가 없는 이유는 단순히 조금 더 깔끔한 코드를 위해(물론 써도 상관 없음)
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	}
	return true
}

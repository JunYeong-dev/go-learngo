package main

import "fmt"

// Node.js, python과 다르게 특정 function을 찾게 되는데, 그게 바로 function main - func main()
// 컴파일러가 main package와 main function을 먼저 찾고 실행시킴
func main() {
	var x string = "Hello World"
	fmt.Println(x)
}

// variables.go
package main

import "fmt"

func main() {
	// 명시적 타입 선언
	var a int = 10
	var b float64 = 20.5
	var c string = "Hello, Go!"

	// 짧은 선언 방식
	d := 42    // 자동으로 int 타입으로 추론
	e := "Go!" // 자동으로 string 타입으로 추론

	// 출력
	fmt.Println(a, b, c, d, e)
}

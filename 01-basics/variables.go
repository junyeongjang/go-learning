// variables.go
package basics

import "fmt"

// 변수 선언 및 타입 추론
func DeclareVariables() {
	var a int = 10
	var b float64 = 20.5
	var c string = "Hello, Go!"

	// 짧은 선언 방식
	d := 42
	e := "Go!"

	fmt.Println(a, b, c, d, e)
}

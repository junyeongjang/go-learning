// conditionals.go
package basics

import "fmt"

// if-else 조건문 예제
func ConditionalExample() {
	age := 25
	if age >= 18 {
		fmt.Println("Adult")
	} else {
		fmt.Println("Minor")
	}
}

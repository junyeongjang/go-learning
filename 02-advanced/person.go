// 02-advanced/person.go
package advanced

import "fmt"

// Person 구조체 정의
type Person struct {
	Name string
	Age  int
}

// Person 구조체에 대한 메서드 정의
func (p Person) Greet() string {
	return "Hello, my name is " + p.Name + " and I am " + fmt.Sprint(p.Age) + " years old."
}

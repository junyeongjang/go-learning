// 02-advanced/interfaces.go
package advanced

import "fmt"

// Speaker 인터페이스 정의
type Speaker interface {
	Speak() string
}

// Person 구조체는 person.go에서 정의되었으므로, 여기서 재정의할 필요 없음
func (p Person) Speak() string {
	return "Hello, my name is " + p.Name
}

// Speaker 인터페이스를 만족하는 함수
func Introduce(speaker Speaker) {
	fmt.Println(speaker.Speak())
}

func InterfaceExample() {
	p := Person{Name: "Bob"}
	Introduce(p) // Person 타입이 Speaker 인터페이스를 만족하므로 호출 가능
}

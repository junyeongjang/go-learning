// 02-advanced/structs.go
package advanced

import "fmt"

// Person 구조체는 person.go에서 정의되었으므로, 여기서 재정의할 필요 없음

func StructExample() {
	p := Person{Name: "Alice", Age: 30}
	fmt.Println(p.Greet()) // person.go에서 정의한 Greet 메서드 호출
}

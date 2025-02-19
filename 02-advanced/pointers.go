// 02-advanced/pointers.go
package advanced

import "fmt"

func Increment(x *int) {
	*x++
}

func PointerExample() {
	a := 5
	fmt.Println("Before increment:", a)
	Increment(&a) // 포인터를 전달하여 값을 수정
	fmt.Println("After increment:", a)
}

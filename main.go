package main

import (
	basics "go-learning/01-basics"     // 'basics' 패키지를 임포트
	advanced "go-learning/02-advanced" // 'advanced' 패키지를 임포트
)

func main() {
	basics.DeclareVariables()
	basics.ConditionalExample()
	basics.LoopExample()

	advanced.StructExample()
	advanced.InterfaceExample()
	advanced.PointerExample()
	advanced.GoroutinesExample()
}

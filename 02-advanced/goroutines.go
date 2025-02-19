package advanced

import (
	"fmt"
	"sync"
)

func GoroutinesExample() {
	// 숫자 배열과 채널 선언
	numbers := make([]int, 100)
	for i := 0; i < 100; i++ {
		numbers[i] = i + 1 // 1부터 100까지 숫자 채우기
	}

	numGoroutines := 10
	var wg sync.WaitGroup
	ch := make(chan int, 100)           // 채널 생성
	uniqueNumbers := make(map[int]bool) // 중복 방지용 맵
	var mu sync.Mutex                   // 맵에 대한 동기화 처리용 뮤텍스

	// 10개의 고루틴을 실행
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(start int) {
			defer wg.Done()
			// 각 고루틴이 처리할 범위
			for j := start; j < 100; j += numGoroutines {
				ch <- numbers[j] // 채널에 숫자 전달
			}
		}(i) // 각 고루틴에 시작 인덱스를 전달
	}

	// 고루틴에서 보낸 숫자를 채널에서 읽고 중복 방지
	go func() {
		for num := range ch {
			mu.Lock() // 맵에 접근하기 전에 락을 걸어서 동시성 문제 해결
			// 중복 방지: 숫자가 이미 있으면 무시
			if !uniqueNumbers[num] {
				uniqueNumbers[num] = true
			}
			mu.Unlock() // 작업이 끝난 후 락을 해제
		}
	}()

	// 모든 고루틴이 완료될 때까지 기다림
	wg.Wait()

	// 채널을 닫은 후, 맵에 들어있는 고유한 숫자 출력
	close(ch)

	// 결과 출력
	fmt.Println("Unique numbers:")
	for num := range uniqueNumbers {
		fmt.Printf("%d ", num)
	}
	fmt.Println() // 새로운 줄
}

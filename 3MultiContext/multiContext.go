package _MultiContext

import (
	"context"
	"fmt"
	"time"
)

func RunMultiContextExample() {
	// 고루틴 1에 대한 컨텍스트 생성
	ctx1, cancel1 := context.WithCancel(context.Background())

	// 고루틴 2에 대한 컨텍스트 생성
	ctx2, cancel2 := context.WithCancel(context.Background())

	// 고루틴 3에 대한 컨텍스트 생성
	ctx3, cancel3 := context.WithCancel(context.Background())

	// 고루틴 1
	go func() {
		select {
		case <-ctx1.Done():
			fmt.Println("Goroutine 1: Context cancelled, exiting")
			return
		}
	}()

	// 고루틴 2
	go func() {
		select {
		case <-ctx2.Done():
			fmt.Println("Goroutine 2: Context cancelled, exiting")
			return
		}
	}()

	// 고루틴 3
	go func() {
		select {
		case <-ctx3.Done():
			fmt.Println("Goroutine 3: Context cancelled, exiting")
			return
		}
	}()

	// 각각의 컨텍스트에 대해 2초 간격으로 취소 신호 전송
	time.Sleep(2 * time.Second)
	fmt.Println("Sending cancel signal to Goroutine 1")
	cancel1()

	time.Sleep(2 * time.Second)
	fmt.Println("Sending cancel signal to Goroutine 2")
	cancel2()

	time.Sleep(2 * time.Second)
	fmt.Println("Sending cancel signal to Goroutine 3")
	cancel3()

	// 고루틴이 모두 종료될 시간을 기다리기 위해 대기
	time.Sleep(1 * time.Second)
}

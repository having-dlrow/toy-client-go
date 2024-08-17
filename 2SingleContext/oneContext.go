package singleContext

import (
	"context"
	"fmt"
	"time"
)

func RunSingleContextExample() {
	ctx, cancel := context.WithCancel(context.Background())

	// 고루틴 1
	go func() {
		select {
		case <-ctx.Done():
			fmt.Println("Goroutine 1: Context cancelled, exiting")
			return
		}
	}()

	// 고루틴 2
	go func() {
		select {
		case <-ctx.Done():
			fmt.Println("Goroutine 2: Context cancelled, exiting")
			return
		}
	}()

	// 고루틴 3
	go func() {
		select {
		case <-ctx.Done():
			fmt.Println("Goroutine 3: Context cancelled, exiting")
			return
		}
	}()

	// 2초 후에 취소 신호 전송
	time.Sleep(2 * time.Second)
	fmt.Println("Sending cancel signal")
	cancel()

	// 고루틴이 모두 종료될 시간을 기다리기 위해 대기
	time.Sleep(1 * time.Second)
}

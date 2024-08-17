package simpleChannel

import (
	"fmt"
	"time"
)

func RunSimpleChannel(stop chan struct{}) {
	ch := make(chan struct{}) // 채널 생성
	defer close(ch)           // 채널 종료

	go func() {
		for {
			select {
			case <-ch:
				fmt.Println("Received ch signal, exiting goroutine")
				return
			default:
				fmt.Println("Goroutine is running")
				time.Sleep(1 * time.Second) // 1초 대기
			}
		}
	}()

	// 채널에서 종료 신호를 받을 때까지 대기
	<-stop

}

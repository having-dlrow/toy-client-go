package channelWorker

import (
	"fmt"
	"sync"
	"time"
)

// Worker 구조체
type Worker struct {
	id          int
	taskChannel chan string
	status      chan bool
	wg          *sync.WaitGroup
}

// Worker의 작업을 수행하는 메서드
func (w *Worker) work() {
	for {
		fmt.Printf("%d loop!\n", w.id)
		select {
		case task, ok := <-w.taskChannel:
			if !ok {
				w.wg.Done()
				return
			}
			fmt.Printf("Worker %d received task: %s\n", w.id, task)
			time.Sleep(1 * time.Second) // 작업 시뮬레이션
			fmt.Printf("Worker %d completed task: %s\n", w.id, task)
			w.status <- true
		}
	}
}

func RunWorkers() {

	taskChannel := make(chan string)
	completeChannel := make(chan bool)
	var wg sync.WaitGroup

	// Worker 생성
	workers := []Worker{
		{id: 1, taskChannel: taskChannel, status: completeChannel, wg: &wg},
		{id: 2, taskChannel: taskChannel, status: completeChannel, wg: &wg},
		{id: 3, taskChannel: taskChannel, status: completeChannel, wg: &wg},
	}

	// Worker 고루틴 시작
	for i := range workers {
		wg.Add(i)
		go workers[i].work()
	}

	// 작업 할당
	taskList := []string{"task1", "task2", "task3", "task4", "task5"}

	go func() {
		for _, task := range taskList {
			fmt.Printf("Sent task: %s\n", task)
			taskChannel <- task
			<-completeChannel // 채널 Done, 작업 완료 대기
		}
		close(taskChannel)
	}()

	wg.Wait()
	fmt.Println("All taskChannel are completed.")
}

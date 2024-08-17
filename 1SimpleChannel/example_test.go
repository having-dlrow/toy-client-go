package simpleChannel

import (
	"testing"
	"time"
)

func TestRunSimpleChannel_Call(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "Test case1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stop := make(chan struct{})

			go RunSimpleChannel(stop)

			// 고루틴이 실행될 시간을 줍니다.
			time.Sleep(2 * time.Second)

			close(stop)

			// 잠시 대기하여 고루틴이 종료될 시간을 줍니다.
			time.Sleep(1 * time.Second)
		})
	}
}

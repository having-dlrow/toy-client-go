package singleContext

import "testing"

func TestRunSingleContextExample(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "Test Case1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RunSingleContextExample()
		})
	}
}

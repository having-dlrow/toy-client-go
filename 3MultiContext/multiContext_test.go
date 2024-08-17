package _MultiContext

import "testing"

func TestRunMultiContextExample(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "Test Case1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RunMultiContextExample()
		})
	}
}

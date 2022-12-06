package pab

import "testing"

func Test_zapTest(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "test_zapTest_1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			zapTest()
		})
	}
}

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

func zapTest() {
	logger := NewLogger("demo.log", "info")
	logger.Infof("info 日志 %s", "abc")
	logger.Errorf("error 日志 %d", 30)
	logger.Errorf("error 日志 %d", 30)
	logger.Errorf("error 日志 %d", 30)
	logger.Errorf("error 日志 %d", 30)
	logger.Errorf("error 日志 %d", 30)
	logger.Errorf("error 日志 %d", 30)
	logger.Errorf("error 日志 %d", 30)
	logger.Infof("info 日志 %s", "abc")
	logger.Infof("info 日志 %s", "abc")
	logger.Infof("info 日志 %s", "abc")
}

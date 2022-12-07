package logger

import (
	"github.com/leigme/pab"
	"reflect"
	"testing"
)

func TestMain(m *testing.M) {
	NewLogger("test.log", "info")
	m.Run()
}

func TestDPanic(t *testing.T) {
	type args struct {
		args []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DPanic(tt.args.args...)
		})
	}
}

func TestDPanicf(t *testing.T) {
	type args struct {
		template string
		args     []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DPanicf(tt.args.template, tt.args.args...)
		})
	}
}

func TestDebug(t *testing.T) {
	type args struct {
		args []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Debug(tt.args.args...)
		})
	}
}

func TestDebugf(t *testing.T) {
	type args struct {
		template string
		args     []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Debugf(tt.args.template, tt.args.args...)
		})
	}
}

func TestError(t *testing.T) {
	type args struct {
		args []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Error(tt.args.args...)
		})
	}
}

func TestErrorf(t *testing.T) {
	type args struct {
		template string
		args     []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Errorf(tt.args.template, tt.args.args...)
		})
	}
}

func TestFatal(t *testing.T) {
	type args struct {
		args []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Fatal(tt.args.args...)
		})
	}
}

func TestFatalf(t *testing.T) {
	type args struct {
		template string
		args     []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Fatalf(tt.args.template, tt.args.args...)
		})
	}
}

func TestInfo(t *testing.T) {
	type args struct {
		args []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Info(tt.args.args...)
		})
	}
}

func TestInfof(t *testing.T) {
	type args struct {
		template string
		args     []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test_Infof_1",
			args: struct {
				template string
				args     []interface{}
			}{template: "test%s", args: nil},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Infof(tt.args.template, tt.args.args...)
		})
	}
}

func TestNewLogger(t *testing.T) {
	type args struct {
		filename string
		level    string
	}
	tests := []struct {
		name string
		args args
		want pab.Logger
	}{
		{
			name: "test_newLogger_1",
			args: args{
				filename: "test.log",
				level:    "info",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewLogger(tt.args.filename, tt.args.level); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLogger() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWarn(t *testing.T) {
	type args struct {
		args []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Warn(tt.args.args...)
		})
	}
}

func TestWarnf(t *testing.T) {
	type args struct {
		template string
		args     []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Warnf(tt.args.template, tt.args.args...)
		})
	}
}

package logger

import (
	"go.uber.org/zap/zapcore"
	"reflect"
	"testing"
)

func TestMain(m *testing.M) {
	InitLoggerWithDefault("demo.log")
	Info("abcdefg")
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
		{
			name: "test_info_1",
			args: args{args: []interface{}{"abcdefg"}},
		},
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Infof(tt.args.template, tt.args.args...)
		})
	}
}

func TestInitLoggerWithConfig(t *testing.T) {
	type args struct {
		conf Config
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InitLoggerWithConfig(tt.args.conf)
		})
	}
}

func TestInitLoggerWithCustom(t *testing.T) {
	type args struct {
		encoder zapcore.Encoder
		writer  zapcore.WriteSyncer
		level   zapcore.LevelEnabler
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InitLoggerWithCustom(tt.args.encoder, tt.args.writer, tt.args.level)
		})
	}
}

func TestInitLoggerWithDefault(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InitLoggerWithDefault(tt.args.filename)
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

func Test_getCatEncoder(t *testing.T) {
	type args struct {
		logFmt string
	}
	tests := []struct {
		name string
		args args
		want zapcore.Encoder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getCatEncoder(tt.args.logFmt); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getCatEncoder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getLogWriter(t *testing.T) {
	type args struct {
		conf Config
	}
	tests := []struct {
		name string
		args args
		want zapcore.WriteSyncer
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getLogWriter(tt.args.conf); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getLogWriter() = %v, want %v", got, tt.want)
			}
		})
	}
}

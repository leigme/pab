package pab

import (
	_ "embed"
	"testing"
)

//go:embed .gitignore
var testData []byte

func TestBytes2File(t *testing.T) {
	type args struct {
		data     []byte
		filename string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test_bytes2file_1",
			args: args{
				data:     testData,
				filename: ".gitignore.test",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Bytes2File(tt.args.data, tt.args.filename); (err != nil) != tt.wantErr {
				t.Errorf("Bytes2File() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCopyFile(t *testing.T) {
	type args struct {
		srcPath string
		dstPath string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test_copyFile_1",
			args: args{
				srcPath: "go.mod",
				dstPath: "go.mod.test",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CopyFile(tt.args.srcPath, tt.args.dstPath); (err != nil) != tt.wantErr {
				t.Errorf("CopyFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCopyFileWithServer(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name: "test_copyFileWithServer_1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CopyFileWithServer(); (err != nil) != tt.wantErr {
				t.Errorf("CopyFileWithServerProperties() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCreateDateDir(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "test_createDateDir_1",
			args:    args{path: "./"},
			want:    "/home/leig/Developer/Github/pab/20221204/pab/20221203",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateDateDir(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateDateDir() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CreateDateDir() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBytes2FileWithApp(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test_bytes2FileWithApp_1",
			args: args{data: testData},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Bytes2FileWithApp(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Bytes2FileWithApp() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

package pab

import (
	"bufio"
	"github.com/leigme/pab/util"
	"io"
	"os"
	"os/user"
	"path/filepath"
	"strings"
	"time"
)

func Bytes2File(data []byte, filename string) error {
	f, err := os.Create(filename)
	defer f.Close()
	if err != nil {
		return err
	}
	w := bufio.NewWriter(f)
	_, err = w.Write(data)
	if err != nil {
		return err
	}
	err = w.Flush()
	if err != nil {
		return err
	}
	return nil
}

func Bytes2FileWithApp(data []byte) error {
	return Bytes2File(data, "app.properties")
}

func CopyFileWithServer() error {
	if strings.EqualFold(util.GetOS(), "linux") {
		u, err := user.Current()
		if err != nil {
			return err
		}
		spp := filepath.Join(u.HomeDir, ".settings", "server.properties")
		return CopyFile("/etc/server.properties", spp)
	}
	return nil
}

func CopyFile(srcPath, dstPath string) error {
	src, err := os.Open(srcPath)
	defer src.Close()
	if err != nil {
		return err
	}
	err = os.MkdirAll(filepath.Dir(dstPath), os.ModePerm)
	if err != nil {
		return err
	}
	dst, err := os.Create(dstPath)
	defer dst.Close()
	if err != nil {
		return err
	}
	w := bufio.NewWriter(dst)
	_, err = io.Copy(w, src)
	if err != nil {
		return err
	}
	err = w.Flush()
	if err != nil {
		return err
	}
	return nil
}

// CreateDateDir Golang的时间格式: 按美式时间格式 月日时分秒年 外加时区 排列起来依次是 01/02 03:04:05PM ‘06 -0700
// 中文时间格式: 2006-01-02 15:04:05 -0700 MST
func CreateDateDir(path string) (string, error) {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return "", err
	}
	folderName := time.Now().Format("20060102")
	folderPath := filepath.Join(path, folderName)
	if _, err = os.Stat(folderPath); os.IsNotExist(err) {
		err = os.Mkdir(folderPath, os.ModePerm)
		if err != nil {
			return "", err
		}
		err = os.Chmod(folderPath, os.ModePerm)
		if err != nil {
			return "", err
		}
	}
	result, err := filepath.Abs(folderPath)
	if err != nil {
		return "", err
	}
	return result, nil
}

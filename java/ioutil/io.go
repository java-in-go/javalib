package ioutil

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"javalib/java/util"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
)

func execCmd(name string, arg ...string) (string, string) {
	command := exec.Command(name, arg...)
	outInfo := bytes.Buffer{}
	errInfo := bytes.Buffer{}
	command.Stdout = &outInfo
	command.Stderr = &errInfo
	err := command.Start()
	if err != nil {
		fmt.Println(err.Error())
	}
	if err = command.Wait(); err != nil {
		fmt.Println(err.Error())
	} else {
		outputMsg := outInfo.String()
		errMsg := errInfo.String()
		return outputMsg, errMsg
	}
	return "", ""
}
func IsDir(name string) bool {
	if info, err := os.Stat(name); err == nil {
		return info.IsDir()
	}
	return false
}
func Exists(filename string) bool {
	existed := true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		existed = false
	}
	return existed
}
func MakeDir(dir string) error {
	if !Exists(dir) {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return err
		}
	}
	return nil
}
func CopyFileBufSize(src string, des string, bufSize int) (written int64, err error) {
	if bufSize <= 0 {
		bufSize = 1 * 1024 * 1024 //1M
	}
	buf := make([]byte, bufSize)
	srcFile, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer srcFile.Close()
	fi, _ := srcFile.Stat()
	perm := fi.Mode()
	desFile, err := os.OpenFile(des, os.O_CREATE|os.O_RDWR|os.O_TRUNC, perm)
	if err != nil {
		return 0, err
	}
	defer desFile.Close()
	count := 0
	for {
		n, err := srcFile.Read(buf)
		if err != nil && err != io.EOF {
			return 0, err
		}
		if n == 0 {
			break
		}
		if wn, err := desFile.Write(buf[:n]); err != nil {
			return 0, err
		} else {
			count += wn
		}
	}
	return int64(count), nil
}
func CopyFile(src string, dest string) (written int64, err error) {
	return CopyFileBufSize(src, dest, 1*1024*1024)
}

func CopyDir(srcPath string, destPath string) error {
	if srcInfo, err := os.Stat(srcPath); err != nil {
		return err
	} else {
		if !srcInfo.IsDir() {
			return errors.New("Invalid srcPath path")
		}
	}
	if desInfo, err := os.Stat(destPath); err != nil {
		return err
	} else {
		if !desInfo.IsDir() {
			return errors.New("Invalid destPath path")
		}
	}
	if strings.TrimSpace(srcPath) == strings.TrimSpace(destPath) {
		return errors.New(util.StrFormat("Files '{}' and '{}' are equal", srcPath, destPath))
	}
	err := filepath.Walk(srcPath, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if path == srcPath {
			return nil
		}
		//生成新路径
		destNewPath := strings.Replace(path, srcPath, destPath, -1)
		if !f.IsDir() {
			if _, err := CopyFile(path, destNewPath); err != nil {
				fmt.Println(err)
			}
		} else {
			if !Exists(destNewPath) {
				return MakeDir(destNewPath)
			}
		}
		return nil
	})
	return err
}

// 获取指定路径下的所有文件，只搜索当前路径，不会扫描子目录
func ListDir(dir string) (files []string, err error) {
	files = []string{}
	dirs, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	for _, file := range dirs {
		if file.IsDir() {
			continue
		}
		files = append(files, path.Join(dir, file.Name()))
	}
	return files, nil
}

// 递归遍历目录以及子目录中的所有文件
func LoopDir(dir string) (files []string, err error) {
	files = []string{}
	err = filepath.Walk(dir, func(fname string, fi os.FileInfo, err error) error {
		if fi.IsDir() {
			//忽略目录
			return nil
		}
		files = append(files, fname)
		return nil
	})
	return files, err
}

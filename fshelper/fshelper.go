package fshelper

import (
	"errors"
	// "fmt"
	funk "github.com/thoas/go-funk"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

func NormalizePath(path string) (string, error) {
	// if path == "" {
	// 	return "", errors.New("Invalid path : " + path)
	// }
	var err error = nil
	if !filepath.IsAbs(path) {
		var base string
		base, err = os.Getwd()
		if err != nil {
			return "", err
		}
		path = filepath.Join(base, path)
	}

	path, err = filepath.Abs(path)
	if err != nil {
		return "", err
	}
	return path, nil
}

type DirInfo struct {
	Path  string
	Files []os.FileInfo
	Dirs  []os.FileInfo
}

func FilterDir(path string, filer func(os.FileInfo) bool) ([]os.FileInfo, error) {
	fileInfos, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}
	filtered := funk.Filter(fileInfos, filer)

	if parsed, ok := filtered.([]os.FileInfo); ok {
		return parsed, nil
	}
	return nil, errors.New("Unexpected Error")
}

func ListDir(path string) (*DirInfo, error) {
	fileInfos, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}
	files := []os.FileInfo{}
	dirs := []os.FileInfo{}
	for _, fileInfo := range fileInfos {
		if fileInfo.IsDir() {
			dirs = append(dirs, fileInfo)
		} else {
			files = append(files, fileInfo)
		}
	}
	return &DirInfo{
		Path:  path,
		Files: files,
		Dirs:  dirs,
	}, err
}

// https://github.com/restic/restic/blob/master/build.go
func DirectoryExists(dirname string) bool {
	stat, err := os.Stat(dirname)
	if err != nil && os.IsNotExist(err) {
		return false
	}

	return stat.IsDir()
}

// CopyFile creates dst from src, preserving file attributes and timestamps.
func CopyFile(src, dst string) error {
	fi, err := os.Stat(src)
	if err != nil {
		return err
	}

	fsrc, err := os.Open(src)
	if err != nil {
		return err
	}

	if err = os.MkdirAll(filepath.Dir(dst), os.ModePerm); err != nil {
		// fmt.Printf("MkdirAll(%v)\n", filepath.Dir(dst))
		return err
	}

	fdst, err := os.Create(dst)
	if err != nil {
		return err
	}

	if _, err = io.Copy(fdst, fsrc); err != nil {
		return err
	}

	if err == nil {
		err = fsrc.Close()
	}

	if err == nil {
		err = fdst.Close()
	}

	if err == nil {
		err = os.Chmod(dst, fi.Mode())
	}

	if err == nil {
		err = os.Chtimes(dst, fi.ModTime(), fi.ModTime())
	}

	return nil
}

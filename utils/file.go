package utils

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"syscall"
)

/*
	Return the file content in a string.
*/
func FileReadAll(filepath string) (string, error) {
	contents, err := ioutil.ReadFile(filepath)
	return string(contents), err
}

/*
	Return the file content in a []byte.
*/
func FileReadAllBytes(filepath string) ([]byte, error) {
	contents, err := ioutil.ReadFile(filepath)
	return contents, err
}

/*
	Write a string in a file.
*/
func FileWriteAll(filepath, content string) error {
	return ioutil.WriteFile(filepath, []byte(content), 0600)
}

/*
	Write a []byte in a file.
*/
func FileWriteAllBytes(filepath string, content []byte) error {
	return ioutil.WriteFile(filepath, content, 0600)
}

/*
	Test if a file exists.
	(if the target is a dir, the function returns false)
*/
func FileExists(path string) bool {
	if info, err := os.Stat(path); err == nil {
		if !info.IsDir() {
			return true
		}
	}
	return false
}

/*
	Test if a dir exists.
	(if the target is a file, the function returns false)
*/
func DirExists(path string) bool {
	if info, err := os.Stat(path); err == nil {
		if info.IsDir() {
			return true
		}
	}
	return false
}

/*
	Size walks a directory tree and returns its total size in bytes.
	This function is from the docker project
*/
func DirectorySize(dir string) (size int64, err error) {
	data := make(map[uint64]struct{})
	err = filepath.Walk(dir, func(d string, fileInfo os.FileInfo, e error) error {
		// Ignore directory sizes
		if fileInfo == nil {
			return nil
		}

		s := fileInfo.Size()
		if fileInfo.IsDir() || s == 0 {
			return nil
		}

		// Check inode to handle hard links correctly
		inode := fileInfo.Sys().(*syscall.Stat_t).Ino
		// inode is not a uint64 on all platforms. Cast it to avoid issues.
		if _, exists := data[uint64(inode)]; exists {
			return nil
		}
		// inode is not a uint64 on all platforms. Cast it to avoid issues.
		data[uint64(inode)] = struct{}{}

		size += s

		return nil
	})
	return
}

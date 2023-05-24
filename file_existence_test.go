package sfs

import (
	"path/filepath"
	"testing"

	"github.com/vault-thirteen/tester"
)

func Test_getFileExistenceUsingCache(t *testing.T) {
	var aTest = tester.New(t)
	var fileExists bool
	var err error
	var sfs *SimpleFileServer
	var path string
	var ok bool

	// Test #1.
	sfs = &SimpleFileServer{
		fileExistenceMap: map[string]bool{},
	}
	path = filepath.Join(TestFolderA, TestFolderB, TestFile)
	fileExists, err = sfs.getFileExistenceUsingCache(path)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(fileExists, true)
	fileExists, ok = sfs.fileExistenceMap[path]
	aTest.MustBeEqual(fileExists, true)
	aTest.MustBeEqual(ok, true)
}

func Test_getFileExistenceWithoutCache(t *testing.T) {
	var aTest = tester.New(t)
	var fileExists bool
	var err error
	var sfs *SimpleFileServer
	var path string

	// Test #1.
	sfs = &SimpleFileServer{}
	path = filepath.Join(TestFolderA, TestFolderB, TestNonExistentFile)
	fileExists, err = sfs.getFileExistenceWithoutCache(path)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(fileExists, false)
}

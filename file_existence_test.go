package sfs

import (
	"path/filepath"
	"testing"

	ssc "github.com/vault-thirteen/Simple-File-Server/SSC"
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
		fileExistenceMap: ssc.NewSSC(10),
	}
	path = filepath.Join(TestFolderA, TestFolderB, TestFile)
	fileExists, err = sfs.getFileExistenceUsingCache(path)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(fileExists, true)
	fileExists, ok = sfs.fileExistenceMap.GetValue(path)
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

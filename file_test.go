package sfs

import (
	"path/filepath"
	"testing"

	"github.com/vault-thirteen/auxie/tester"
)

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

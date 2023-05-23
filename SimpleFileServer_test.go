package sfs

import (
	"path/filepath"
	"testing"

	"github.com/vault-thirteen/tester"
)

const TestFile = "a.txt"

func Test_NewSimpleFileServer(t *testing.T) {
	var aTest = tester.New(t)
	var err error
	var result *SimpleFileServer

	// Test #1.
	result, err = NewSimpleFileServer("non-existent-folder", 0, 0, 60)
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(result, (*SimpleFileServer)(nil))

	// Test #2.
	existingDataFolder := filepath.Join("test", "data")
	result, err = NewSimpleFileServer(existingDataFolder, 100, 1_000_000, 60)
	aTest.MustBeNoError(err)
	aTest.MustBeDifferent(result, (*SimpleFileServer)(nil))
}

func Test_GetFileContents(t *testing.T) {
	var aTest = tester.New(t)
	var err error
	var sfs *SimpleFileServer
	existingDataFolder := filepath.Join("test", "data")
	sfs, err = NewSimpleFileServer(existingDataFolder, 100, 1_000_000, 60)
	aTest.MustBeNoError(err)
	var data []byte
	var fileExists bool

	// Test #1.
	data, fileExists, err = sfs.GetFileContents(TestFile)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(data, []byte("This is a test file."))
	aTest.MustBeEqual(fileExists, true)

	// Test #2.
	data, fileExists, err = sfs.GetFileContents("non-existent-file")
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(data, ([]byte)(nil))
	aTest.MustBeEqual(fileExists, false)

	// Test #3.
	data, fileExists, err = sfs.GetFileContents("abc....def")
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(data, ([]byte)(nil))
	aTest.MustBeEqual(fileExists, false)
}

func Test_isFilePathValid(t *testing.T) {
	var aTest = tester.New(t)
	sfs := &SimpleFileServer{}
	aTest.MustBeEqual(sfs.isFilePathValid("q"), true)
	aTest.MustBeEqual(sfs.isFilePathValid(".."), false)
}

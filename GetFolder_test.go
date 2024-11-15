package sfs

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/vault-thirteen/auxie/tester"
)

func Test_GetFolder(t *testing.T) {
	var aTest = tester.New(t)
	var bytes []byte
	var err error
	var sfs *SimpleFileServer

	// Test #1. Empty list of default files.
	existingDataFolder := filepath.Join(TestFolderA, TestFolderB)
	sfs, err = NewSimpleFileServer(existingDataFolder, nil, false, 0, 0, 0)
	aTest.MustBeNoError(err)
	//
	bytes, err = sfs.GetFolder(string(os.PathSeparator))
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(bytes, []byte(nil))

	// Test #2. Invalid path.
	existingDataFolder = filepath.Join(TestFolderA, TestFolderB)
	sfs, err = NewSimpleFileServer(existingDataFolder, []string{"index.html"}, false, 0, 0, 0)
	aTest.MustBeNoError(err)
	//
	bytes, err = sfs.GetFolder("..")
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(bytes, []byte(nil))

	// Test #3. File does not exist.
	existingDataFolder = filepath.Join(TestFolderA, TestFolderB)
	sfs, err = NewSimpleFileServer(existingDataFolder, []string{"index.txt"}, false, 0, 0, 0)
	aTest.MustBeNoError(err)
	//
	bytes, err = sfs.GetFolder(string(os.PathSeparator))
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(bytes, []byte(nil))

	// Test #4. List of default file names contains several values.
	existingDataFolder = filepath.Join(TestFolderA, TestFolderB)
	sfs, err = NewSimpleFileServer(existingDataFolder, []string{"index.txt", "index.html"}, false, 0, 0, 0)
	aTest.MustBeNoError(err)
	//
	bytes, err = sfs.GetFolder(string(os.PathSeparator))
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(bytes, []byte(`<html>.</html>`))
}

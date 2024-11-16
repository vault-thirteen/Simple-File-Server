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

	// Test #1. Invalid path.
	existingDataFolder := filepath.Join(TestFolderA, TestFolderB)
	sfs, err = NewSimpleFileServer(existingDataFolder, []string{"index.html"}, false, 0, 0, 0)
	aTest.MustBeNoError(err)
	//
	bytes, err = sfs.GetFolder("..")
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(bytes, []byte(nil))

	// Test #2. Caching is disabled, file exists.
	existingDataFolder = filepath.Join(TestFolderA, TestFolderB)
	sfs, err = NewSimpleFileServer(existingDataFolder, []string{"index.php", "index.html"}, false, 0, 0, 0)
	aTest.MustBeNoError(err)
	//
	bytes, err = sfs.GetFolder(string(os.PathSeparator))
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(bytes, []byte("<html>.</html>"))

	// Test #4. Caching is enabled, file exists.
	existingDataFolder = filepath.Join(TestFolderA, TestFolderB)
	sfs, err = NewSimpleFileServer(existingDataFolder, []string{"index.php", "index.html"}, true, 10, 1_000_000, 60)
	aTest.MustBeNoError(err)
	//
	bytes, err = sfs.GetFolder(string(os.PathSeparator))
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(bytes, []byte("<html>.</html>"))
}

func Test_getFolderFromStorage(t *testing.T) {
	var aTest = tester.New(t)
	var bytes []byte
	var err error
	var sfs *SimpleFileServer

	// Test #1. Empty list of default files.
	existingDataFolder := filepath.Join(TestFolderA, TestFolderB)
	sfs, err = NewSimpleFileServer(existingDataFolder, nil, false, 0, 0, 0)
	aTest.MustBeNoError(err)
	//
	bytes, err = sfs.getFolderFromStorage(string(os.PathSeparator))
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(bytes, []byte(nil))

	// Test #2. File does not exist.
	existingDataFolder = filepath.Join(TestFolderA, TestFolderB)
	sfs, err = NewSimpleFileServer(existingDataFolder, []string{"index.txt"}, false, 0, 0, 0)
	aTest.MustBeNoError(err)
	//
	bytes, err = sfs.getFolderFromStorage(string(os.PathSeparator))
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(bytes, []byte(nil))

	// Test #3. List of default file names contains several values, one of them
	// exists.
	existingDataFolder = filepath.Join(TestFolderA, TestFolderB)
	sfs, err = NewSimpleFileServer(existingDataFolder, []string{"index.php", "index.html"}, false, 0, 0, 0)
	aTest.MustBeNoError(err)
	//
	bytes, err = sfs.getFolderFromStorage(string(os.PathSeparator))
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(bytes, []byte("<html>.</html>"))
}

package sfs

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/vault-thirteen/auxie/tester"
)

func Test_GetFolderDefaultFilename(t *testing.T) {
	var aTest = tester.New(t)
	var fileName string
	var err error
	var sfs *SimpleFileServer

	// Test #1. Empty list of default files.
	existingDataFolder := filepath.Join(TestFolderA, TestFolderB)
	sfs, err = NewSimpleFileServer(existingDataFolder, nil, false, 0, 0, 0)
	aTest.MustBeNoError(err)
	//
	fileName, err = sfs.GetFolderDefaultFilename(string(os.PathSeparator))
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(fileName, "")

	// Test #2. Invalid path.
	existingDataFolder = filepath.Join(TestFolderA, TestFolderB)
	sfs, err = NewSimpleFileServer(existingDataFolder, []string{"index.html"}, false, 0, 0, 0)
	aTest.MustBeNoError(err)
	//
	fileName, err = sfs.GetFolderDefaultFilename("..")
	aTest.MustBeAnError(err)

	// Test #3. File does not exist.
	existingDataFolder = filepath.Join(TestFolderA, TestFolderB)
	sfs, err = NewSimpleFileServer(existingDataFolder, []string{"index.txt"}, false, 0, 0, 0)
	aTest.MustBeNoError(err)
	//
	fileName, err = sfs.GetFolderDefaultFilename(string(os.PathSeparator))
	aTest.MustBeAnError(err)

	// Test #4. List of default file names contains several values.
	existingDataFolder = filepath.Join(TestFolderA, TestFolderB)
	sfs, err = NewSimpleFileServer(existingDataFolder, []string{"index.txt", "index.html"}, false, 0, 0, 0)
	aTest.MustBeNoError(err)
	//
	fileName, err = sfs.GetFolderDefaultFilename(string(os.PathSeparator))
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(fileName, "index.html")
}

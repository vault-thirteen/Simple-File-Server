package sfs

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/vault-thirteen/tester"
)

func Test_GetFolderDefaultFilename(t *testing.T) {
	var aTest = tester.New(t)
	var fileName string
	var fileExists bool
	var err error
	var sfs *SimpleFileServer

	// Test #1. Empty list of default files.
	existingDataFolder := filepath.Join(TestFolderA, TestFolderB)
	sfs, err = NewSimpleFileServer(existingDataFolder, nil, false, 0, 0, 0)
	aTest.MustBeNoError(err)
	//
	fileName, fileExists, err = sfs.GetFolderDefaultFilename(string(os.PathSeparator))
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(fileExists, false)
	aTest.MustBeEqual(fileName, "")

	// Test #2. Invalid path.
	existingDataFolder = filepath.Join(TestFolderA, TestFolderB)
	sfs, err = NewSimpleFileServer(existingDataFolder, []string{"index.html"}, false, 0, 0, 0)
	aTest.MustBeNoError(err)
	//
	fileName, fileExists, err = sfs.GetFolderDefaultFilename("..")
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(fileExists, false)
	aTest.MustBeEqual(fileName, "")

	// Test #3. Caching is enabled.
	existingDataFolder = filepath.Join(TestFolderA, TestFolderB)
	sfs, err = NewSimpleFileServer(existingDataFolder, []string{"index.html"}, true, 1000, 1_000_000, 60)
	aTest.MustBeNoError(err)
	//
	fileName, fileExists, err = sfs.GetFolderDefaultFilename(string(os.PathSeparator))
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(fileExists, true)
	aTest.MustBeEqual(fileName, "index.html")

	// Test #4. Caching is disabled.
	existingDataFolder = filepath.Join(TestFolderA, TestFolderB)
	sfs, err = NewSimpleFileServer(existingDataFolder, []string{"index.html"}, false, 0, 0, 0)
	aTest.MustBeNoError(err)
	//
	fileName, fileExists, err = sfs.GetFolderDefaultFilename(string(os.PathSeparator))
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(fileExists, true)
	aTest.MustBeEqual(fileName, "index.html")

	// Test #5. File does not exist.
	existingDataFolder = filepath.Join(TestFolderA, TestFolderB)
	sfs, err = NewSimpleFileServer(existingDataFolder, []string{"index.txt"}, false, 0, 0, 0)
	aTest.MustBeNoError(err)
	//
	fileName, fileExists, err = sfs.GetFolderDefaultFilename(string(os.PathSeparator))
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(fileExists, false)
	aTest.MustBeEqual(fileName, "")

	// Test #6. List of default file names contains several values.
	existingDataFolder = filepath.Join(TestFolderA, TestFolderB)
	sfs, err = NewSimpleFileServer(existingDataFolder, []string{"index.txt", "index.html"}, false, 0, 0, 0)
	aTest.MustBeNoError(err)
	//
	fileName, fileExists, err = sfs.GetFolderDefaultFilename(string(os.PathSeparator))
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(fileExists, true)
	aTest.MustBeEqual(fileName, "index.html")
}

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
	var fileExists bool
	var err error
	var sfs *SimpleFileServer

	// Test #1. Empty list of default files.
	existingDataFolder := filepath.Join(TestFolderA, TestFolderB)
	sfs, err = NewSimpleFileServer(existingDataFolder, nil, false, 0, 0, 0)
	aTest.MustBeNoError(err)
	//
	bytes, fileExists, err = sfs.GetFolder(string(os.PathSeparator))
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(fileExists, false)
	aTest.MustBeEqual(bytes, []byte(nil))

	// Test #2. Invalid path.
	existingDataFolder = filepath.Join(TestFolderA, TestFolderB)
	sfs, err = NewSimpleFileServer(existingDataFolder, []string{"index.html"}, false, 0, 0, 0)
	aTest.MustBeNoError(err)
	//
	bytes, fileExists, err = sfs.GetFolder("..")
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(fileExists, false)
	aTest.MustBeEqual(bytes, []byte(nil))

	// Test #3. Caching is enabled.
	existingDataFolder = filepath.Join(TestFolderA, TestFolderB)
	sfs, err = NewSimpleFileServer(existingDataFolder, []string{"index.html"}, true, 1000, 1_000_000, 60)
	aTest.MustBeNoError(err)
	//
	bytes, fileExists, err = sfs.GetFolder(string(os.PathSeparator))
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(fileExists, true)
	aTest.MustBeEqual(bytes, []byte(`<html>.</html>`))

	// Test #4. Caching is disabled.
	existingDataFolder = filepath.Join(TestFolderA, TestFolderB)
	sfs, err = NewSimpleFileServer(existingDataFolder, []string{"index.html"}, false, 0, 0, 0)
	aTest.MustBeNoError(err)
	//
	bytes, fileExists, err = sfs.GetFolder(string(os.PathSeparator))
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(fileExists, true)
	aTest.MustBeEqual(bytes, []byte(`<html>.</html>`))

	// Test #5. File does not exist.
	existingDataFolder = filepath.Join(TestFolderA, TestFolderB)
	sfs, err = NewSimpleFileServer(existingDataFolder, []string{"index.txt"}, false, 0, 0, 0)
	aTest.MustBeNoError(err)
	//
	bytes, fileExists, err = sfs.GetFolder(string(os.PathSeparator))
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(fileExists, false)
	aTest.MustBeEqual(bytes, []byte(nil))

	// Test #6. List of default file names contains several values.
	existingDataFolder = filepath.Join(TestFolderA, TestFolderB)
	sfs, err = NewSimpleFileServer(existingDataFolder, []string{"index.txt", "index.html"}, false, 0, 0, 0)
	aTest.MustBeNoError(err)
	//
	bytes, fileExists, err = sfs.GetFolder(string(os.PathSeparator))
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(fileExists, true)
	aTest.MustBeEqual(bytes, []byte(`<html>.</html>`))
}

func Test_getDefaultFileUsingCache(t *testing.T) {
	var aTest = tester.New(t)
	var bytes []byte
	var fileExists bool
	var err error
	var sfs *SimpleFileServer
	var path string

	// Test #1. File does not exist.
	existingDataFolder := filepath.Join(TestFolderA, TestFolderB)
	sfs, err = NewSimpleFileServer(existingDataFolder, []string{"index.html"}, true, 1000, 1_000_000, 60)
	aTest.MustBeNoError(err)
	//
	bytes, fileExists, err = sfs.getDefaultFileUsingCache("index.txt")
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(fileExists, false)
	aTest.MustBeEqual(bytes, []byte(nil))

	// Test #2.1. File exists. Cache is empty.
	existingDataFolder = filepath.Join(TestFolderA, TestFolderB)
	sfs, err = NewSimpleFileServer(existingDataFolder, []string{"index.html"}, true, 1000, 1_000_000, 60)
	aTest.MustBeNoError(err)
	//
	path = filepath.Join(existingDataFolder, "index.html")
	bytes, fileExists, err = sfs.getDefaultFileUsingCache(path)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(fileExists, true)
	aTest.MustBeEqual(bytes, []byte(`<html>.</html>`))

	// Test #2.2. File exists. Cache is set.
	bytes, fileExists, err = sfs.getDefaultFileUsingCache(path)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(fileExists, true)
	aTest.MustBeEqual(bytes, []byte(`<html>.</html>`))
}

func Test_getDefaultFileWithoutCache(t *testing.T) {
	var aTest = tester.New(t)
	var bytes []byte
	var fileExists bool
	var err error
	var sfs *SimpleFileServer
	var path string

	// Test #1.
	existingDataFolder := filepath.Join(TestFolderA, TestFolderB)
	sfs, err = NewSimpleFileServer(existingDataFolder, nil, false, 0, 0, 0)
	aTest.MustBeNoError(err)
	//
	path = filepath.Join(existingDataFolder, TestFile)
	bytes, fileExists, err = sfs.getDefaultFileWithoutCache(path)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(fileExists, true)
	aTest.MustBeEqual(bytes, []byte("This is a test file."))
}

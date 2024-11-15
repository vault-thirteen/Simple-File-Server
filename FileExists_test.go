package sfs

import (
	"path/filepath"
	"testing"

	"github.com/vault-thirteen/auxie/tester"
)

func Test_FileExists(t *testing.T) {
	var aTest = tester.New(t)
	var x bool
	var err error
	var sfs *SimpleFileServer

	// Test #1. No cache.
	existingDataFolder := filepath.Join(TestFolderA, TestFolderB)
	sfs, err = NewSimpleFileServer(existingDataFolder, nil, false, 5, 100, 60)
	aTest.MustBeNoError(err)
	//
	x, err = sfs.FileExists(TestFileA)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(x, true)

	// Test #2. Using cache.
	existingDataFolder = filepath.Join(TestFolderA, TestFolderB)
	sfs, err = NewSimpleFileServer(existingDataFolder, nil, true, 5, 100, 60)
	aTest.MustBeNoError(err)
	//
	x, err = sfs.FileExists(TestFileA)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(x, true)
}

func Test_fileExistsInCache(t *testing.T) {
	var aTest = tester.New(t)
	var x bool
	var err error
	var sfs *SimpleFileServer

	// Test #1.
	existingDataFolder := filepath.Join(TestFolderA, TestFolderB)
	sfs, err = NewSimpleFileServer(existingDataFolder, nil, true, 5, 100, 60)
	aTest.MustBeNoError(err)
	//
	x = sfs.fileExistsInCache(TestFileA)
	aTest.MustBeEqual(x, false)
	_, err = sfs.GetFile(TestFileA)
	aTest.MustBeNoError(err)
	x = sfs.fileExistsInCache(TestFileA)
	aTest.MustBeEqual(x, true)
}

func Test_fileExistsInStorage(t *testing.T) {
	var aTest = tester.New(t)
	var fileExists bool
	var err error
	var sfs *SimpleFileServer
	var path string

	// Test #1.
	sfs = &SimpleFileServer{}
	path = filepath.Join(TestFolderA, TestFolderB, TestNonExistentFile)
	fileExists, err = sfs.fileExistsInStorage(path)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(fileExists, false)
}

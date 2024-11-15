package sfs

import (
	"path/filepath"
	"testing"

	"github.com/vault-thirteen/auxie/tester"
)

func Test_GetFile(t *testing.T) {
	var aTest = tester.New(t)
	var bytes []byte
	var err error
	var sfs *SimpleFileServer

	// Test #1. Invalid path.
	existingDataFolder := filepath.Join(TestFolderA, TestFolderB)
	sfs, err = NewSimpleFileServer(existingDataFolder, nil, false, 0, 0, 0)
	aTest.MustBeNoError(err)
	//
	bytes, err = sfs.GetFile("x/../x")
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(bytes, []byte(nil))

	// Test #2. No caching.
	existingDataFolder = filepath.Join(TestFolderA, TestFolderB)
	sfs, err = NewSimpleFileServer(existingDataFolder, nil, false, 0, 0, 0)
	aTest.MustBeNoError(err)
	//
	bytes, err = sfs.GetFile(TestFileA)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(bytes, []byte("This is a test file."))

	// Test #2. File is in cache.
	existingDataFolder = filepath.Join(TestFolderA, TestFolderB)
	sfs, err = NewSimpleFileServer(existingDataFolder, nil, false, 0, 0, 0)
	aTest.MustBeNoError(err)
	//
	bytes, err = sfs.GetFile(TestFileA)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(bytes, []byte("This is a test file."))
	bytes, err = sfs.GetFile(TestFileA)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(bytes, []byte("This is a test file."))
}

func Test_getFileUsingCache(t *testing.T) {
	var aTest = tester.New(t)
	var bytes []byte
	var err error
	var sfs *SimpleFileServer

	// Test #1.
	existingDataFolder := filepath.Join(TestFolderA, TestFolderB)
	sfs, err = NewSimpleFileServer(existingDataFolder, nil, true, 100, 1_000_000, 60)
	aTest.MustBeNoError(err)
	//
	_, err = sfs.GetFile(TestFileA)
	aTest.MustBeNoError(err)
	bytes, err = sfs.getFileUsingCache(TestFileA)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(bytes, []byte("This is a test file."))
}

func Test_getFileWithoutCache(t *testing.T) {
	var aTest = tester.New(t)
	var bytes []byte
	var err error
	var sfs *SimpleFileServer

	// Test #1. File does not exist.
	existingDataFolder := filepath.Join(TestFolderA, TestFolderB)
	sfs, err = NewSimpleFileServer(existingDataFolder, nil, false, 0, 0, 0)
	aTest.MustBeNoError(err)
	//
	bytes, err = sfs.getFileWithoutCache(TestNonExistentFile)
	aTest.MustBeAnError(err)

	// Test #2. File exists.
	existingDataFolder = filepath.Join(TestFolderA, TestFolderB)
	sfs, err = NewSimpleFileServer(existingDataFolder, nil, false, 0, 0, 0)
	aTest.MustBeNoError(err)
	//
	bytes, err = sfs.getFileWithoutCache(TestFileA)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(bytes, []byte("This is a test file."))
}

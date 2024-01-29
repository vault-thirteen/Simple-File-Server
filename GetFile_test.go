package sfs

import (
	"path/filepath"
	"testing"

	"github.com/vault-thirteen/auxie/tester"
)

func Test_GetFile(t *testing.T) {
	var aTest = tester.New(t)
	var bytes []byte
	var fileExists bool
	var err error
	var sfs *SimpleFileServer

	// Test #1. Invalid path.
	existingDataFolder := filepath.Join(TestFolderA, TestFolderB)
	sfs, err = NewSimpleFileServer(existingDataFolder, nil, false, 0, 0, 0)
	aTest.MustBeNoError(err)
	//
	bytes, fileExists, err = sfs.GetFile("x/../x")
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(fileExists, false)
	aTest.MustBeEqual(bytes, []byte(nil))

	// Test #2. Valid path.
	existingDataFolder = filepath.Join(TestFolderA, TestFolderB)
	sfs, err = NewSimpleFileServer(existingDataFolder, nil, false, 0, 0, 0)
	aTest.MustBeNoError(err)
	//
	bytes, fileExists, err = sfs.GetFile(TestFile)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(fileExists, true)
	aTest.MustBeEqual(bytes, []byte("This is a test file."))
}

func Test_getFileUsingCache(t *testing.T) {
	var aTest = tester.New(t)
	var bytes []byte
	var fileExists bool
	var err error
	var sfs *SimpleFileServer

	// Test #1. File does not exist.
	existingDataFolder := filepath.Join(TestFolderA, TestFolderB)
	sfs, err = NewSimpleFileServer(existingDataFolder, nil, true, 100, 1_000_000, 60)
	aTest.MustBeNoError(err)
	//
	bytes, fileExists, err = sfs.getFileUsingCache(TestNonExistentFile)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(fileExists, false)
	aTest.MustBeEqual(bytes, []byte(nil))

	// Test #2.1. File exists and cache is empty.
	existingDataFolder = filepath.Join(TestFolderA, TestFolderB)
	sfs, err = NewSimpleFileServer(existingDataFolder, nil, true, 100, 1_000_000, 60)
	aTest.MustBeNoError(err)
	//
	bytes, fileExists, err = sfs.getFileUsingCache(TestFile)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(fileExists, true)
	aTest.MustBeEqual(bytes, []byte("This is a test file."))

	// Test #2.2. File exists and cache is set.
	bytes, fileExists, err = sfs.getFileUsingCache(TestFile)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(fileExists, true)
	aTest.MustBeEqual(bytes, []byte("This is a test file."))
}

func Test_getFileWithoutCache(t *testing.T) {
	var aTest = tester.New(t)
	var bytes []byte
	var fileExists bool
	var err error
	var sfs *SimpleFileServer

	// Test #1. File does not exist.
	existingDataFolder := filepath.Join(TestFolderA, TestFolderB)
	sfs, err = NewSimpleFileServer(existingDataFolder, nil, false, 0, 0, 0)
	aTest.MustBeNoError(err)
	//
	bytes, fileExists, err = sfs.getFileWithoutCache(TestNonExistentFile)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(fileExists, false)
	aTest.MustBeEqual(bytes, []byte(nil))

	// Test #2. File exists.
	existingDataFolder = filepath.Join(TestFolderA, TestFolderB)
	sfs, err = NewSimpleFileServer(existingDataFolder, nil, false, 0, 0, 0)
	aTest.MustBeNoError(err)
	//
	bytes, fileExists, err = sfs.getFileWithoutCache(TestFile)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(fileExists, true)
	aTest.MustBeEqual(bytes, []byte("This is a test file."))
}

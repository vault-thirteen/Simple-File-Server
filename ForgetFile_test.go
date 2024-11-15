package sfs

import (
	"path/filepath"
	"testing"

	"github.com/vault-thirteen/auxie/tester"
)

func Test_ForgetFile(t *testing.T) {
	var aTest = tester.New(t)
	var err error
	var sfs *SimpleFileServer
	var x bool

	// Test #1. Cache is disabled.
	existingDataFolder := filepath.Join(TestFolderA, TestFolderB)
	sfs, err = NewSimpleFileServer(existingDataFolder, nil, false, 0, 0, 0)
	aTest.MustBeNoError(err)
	//
	err = sfs.ForgetFile(TestFileA)
	aTest.MustBeAnError(err)

	// Test #2. Invalid path.
	existingDataFolder = filepath.Join(TestFolderA, TestFolderB)
	sfs, err = NewSimpleFileServer(existingDataFolder, nil, true, 0, 0, 60)
	aTest.MustBeNoError(err)
	//
	err = sfs.ForgetFile("x/../x")
	aTest.MustBeAnError(err)

	// Test #2. Valid path.
	existingDataFolder = filepath.Join(TestFolderA, TestFolderB)
	sfs, err = NewSimpleFileServer(existingDataFolder, nil, true, 5, 100, 60)
	aTest.MustBeNoError(err)
	//
	_, err = sfs.GetFile(TestFileA)
	aTest.MustBeNoError(err)
	x = sfs.fileExistsInCache(TestFileA)
	aTest.MustBeEqual(x, true)
	err = sfs.ForgetFile(TestFileA)
	aTest.MustBeNoError(err)
	x = sfs.fileExistsInCache(TestFileA)
	aTest.MustBeEqual(x, false)
}

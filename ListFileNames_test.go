package sfs

import (
	"path/filepath"
	"testing"

	"github.com/vault-thirteen/auxie/tester"
)

func Test_ListFileNames(t *testing.T) {
	var aTest = tester.New(t)

	// Test #1.
	existingDataFolder := filepath.Join(TestFolderA, TestFolderB)
	sfs, err := NewSimpleFileServer(existingDataFolder, nil, false, 0, 0, 0)
	aTest.MustBeNoError(err)
	//
	var fileNames []string
	fileNames, err = sfs.ListFileNames()
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(fileNames, []string{TestFileA, TestFileIndex})
}

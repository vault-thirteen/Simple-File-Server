package sfs

import (
	"path/filepath"
	"testing"

	"github.com/vault-thirteen/auxie/tester"
)

func Test_CountFiles(t *testing.T) {
	var aTest = tester.New(t)

	// Test #1.
	existingDataFolder := filepath.Join(TestFolderA, TestFolderB)
	sfs, err := NewSimpleFileServer(existingDataFolder, nil, false, 0, 0, 0)
	aTest.MustBeNoError(err)
	//
	var filesCount int
	filesCount, err = sfs.CountFiles()
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(filesCount, len([]string{TestFileA, TestFileIndex}))
}

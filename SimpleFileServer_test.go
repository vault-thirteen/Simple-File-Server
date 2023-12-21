package sfs

import (
	"path/filepath"
	"testing"

	"github.com/vault-thirteen/auxie/tester"
)

func Test_NewSimpleFileServer(t *testing.T) {
	var aTest = tester.New(t)
	var err error
	var result *SimpleFileServer

	// Test #1.
	result, err = NewSimpleFileServer(TestNonExistentFolder, nil, true, 0, 0, 60)
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(result, (*SimpleFileServer)(nil))

	// Test #2.
	existingDataFolder := filepath.Join(TestFolderA, TestFolderB)
	result, err = NewSimpleFileServer(existingDataFolder, []string{"x"}, true, 100, 1_000_000, 60)
	aTest.MustBeNoError(err)
	aTest.MustBeDifferent(result, (*SimpleFileServer)(nil))
}

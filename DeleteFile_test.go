package sfs

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/vault-thirteen/auxie/tester"
)

func Test_DeleteFile(t *testing.T) {
	var aTest = tester.New(t)

	existingDataFolder := filepath.Join(TestFolderA, TestFolderB)
	f, err := os.Create(filepath.Join(existingDataFolder, TestFileX))
	aTest.MustBeNoError(err)
	err = f.Close()
	aTest.MustBeNoError(err)
	time.Sleep(1 * time.Second)

	// Test #1.
	var sfs *SimpleFileServer
	sfs, err = NewSimpleFileServer(existingDataFolder, nil, false, 0, 0, 0)
	aTest.MustBeNoError(err)
	//
	err = sfs.DeleteFile(TestFileX)
	aTest.MustBeNoError(err)
}

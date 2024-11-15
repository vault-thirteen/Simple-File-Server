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
	var x bool

	existingDataFolder := filepath.Join(TestFolderA, TestFolderB)
	f, err := os.Create(filepath.Join(existingDataFolder, TestFileX))
	aTest.MustBeNoError(err)
	_, err = f.WriteString("test")
	aTest.MustBeNoError(err)
	err = f.Close()
	aTest.MustBeNoError(err)
	time.Sleep(1 * time.Second)

	// Test #1.
	var sfs *SimpleFileServer
	sfs, err = NewSimpleFileServer(existingDataFolder, nil, true, 0, 0, 60)
	aTest.MustBeNoError(err)
	//
	_, err = sfs.GetFile(TestFileX)
	aTest.MustBeNoError(err)
	x = sfs.fileExistsInCache(TestFileX)
	aTest.MustBeEqual(x, true)
	err = sfs.DeleteFile(TestFileX)
	aTest.MustBeNoError(err)
	x = sfs.fileExistsInCache(TestFileX)
	aTest.MustBeEqual(x, false)
	x, err = sfs.fileExistsInStorage(TestFileX)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(x, false)
}

package sfs

import (
	"path/filepath"
	"testing"

	"github.com/vault-thirteen/auxie/tester"
)

const (
	TestFolderA           = "test"
	TestFolderB           = "data"
	TestFileA             = "a.txt"
	TestFileIndex         = "index.html"
	TestFileX             = "x.tmp.txt"
	TestNonExistentFile   = "non-existent-file"
	TestNonExistentFolder = "non-existent-folder"
)

func Test_IsPathValid(t *testing.T) {
	var aTest = tester.New(t)
	aTest.MustBeEqual(IsPathValid("q"), true)
	aTest.MustBeEqual(IsPathValid(".."), false)
}

func Test_IsPathFolder(t *testing.T) {
	var aTest = tester.New(t)
	aTest.MustBeEqual(IsPathFolder(``), false)
	aTest.MustBeEqual(IsPathFolder(`a`), false)
	aTest.MustBeEqual(IsPathFolder(`/`), true)
	aTest.MustBeEqual(IsPathFolder(`\`), true)
	aTest.MustBeEqual(IsPathFolder(`/a`), false)
	aTest.MustBeEqual(IsPathFolder(`\a`), false)
	aTest.MustBeEqual(IsPathFolder(`a/`), true)
	aTest.MustBeEqual(IsPathFolder(`a\`), true)
}

func Test_GetAbsolutePath(t *testing.T) {
	aTest := tester.New(t)
	var sfs *SimpleFileServer
	var relpath string
	var err error

	// Test #1.
	existingDataFolder := filepath.Join(TestFolderA, TestFolderB)
	sfs, err = NewSimpleFileServer(existingDataFolder, []string{}, true, 1000, 1_000_000, 60)
	aTest.MustBeNoError(err)
	//
	relpath = `test`
	aTest.MustBeEqual(sfs.GetAbsolutePath(relpath), filepath.Join(existingDataFolder, relpath))
}

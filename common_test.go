package sfs

import (
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

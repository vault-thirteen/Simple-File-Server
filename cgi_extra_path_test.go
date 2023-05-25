package sfs

import (
	"path/filepath"
	"testing"

	"github.com/vault-thirteen/tester"
)

func Test_FindExtraPath(t *testing.T) {
	aTest := tester.New(t)
	var sfs *SimpleFileServer
	var relpath string
	var path string
	var extraPath string
	var err error

	// Test #1. Extra path exists.
	existingDataFolder := filepath.Join(TestFolderA, TestFolderB)
	sfs, err = NewSimpleFileServer(existingDataFolder, []string{}, true, 1000, 1_000_000, 60)
	aTest.MustBeNoError(err)
	//
	relpath = `/cgi-bin/display.pl/cgi/cgi_doc.txt`
	path, extraPath, err = sfs.FindExtraPath(relpath)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(path, `/cgi-bin/display.pl`)
	aTest.MustBeEqual(extraPath, `/cgi/cgi_doc.txt`)

	// Test #2. Extra path does not exist.
	existingDataFolder = filepath.Join(TestFolderA, TestFolderB)
	sfs, err = NewSimpleFileServer(existingDataFolder, []string{}, true, 1000, 1_000_000, 60)
	aTest.MustBeNoError(err)
	//
	relpath = `/cgi-bin/display.xyz/cgi/cgi_doc.txt`
	path, extraPath, err = sfs.FindExtraPath(relpath)
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(err.Error(), "CGI extra path is not found")
}

func Test_shortenPath(t *testing.T) {
	aTest := tester.New(t)
	aTest.MustBeEqual(shortenPath(``), ``)
	aTest.MustBeEqual(shortenPath(`a`), ``)
	aTest.MustBeEqual(shortenPath(`/a/`), `/a`)
	aTest.MustBeEqual(shortenPath(`/a/b`), `/a`)
	aTest.MustBeEqual(shortenPath(`/a/b/c`), `/a/b`)
	aTest.MustBeEqual(shortenPath(`\a\`), `\a`)
	aTest.MustBeEqual(shortenPath(`\a\b`), `\a`)
	aTest.MustBeEqual(shortenPath(`\a\b\c`), `\a\b`)
}

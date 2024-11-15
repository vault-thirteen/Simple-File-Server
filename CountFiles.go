package sfs

import (
	"github.com/vault-thirteen/auxie/file"
)

// CountFiles counts files in the storage folder. Cache is not used while may
// not contain all the files.
func (sfs *SimpleFileServer) CountFiles() (filesCount int, err error) {
	return file.CountFiles(sfs.rootFolderPath)
}

package sfs

import (
	"github.com/vault-thirteen/auxie/file"
)

// CountFiles counts files in the storage folder. Cache is not used while may
// not contain all the files.
func (sfs *SimpleFileServer) CountFiles() (filesCount int, err error) {
	filesCount, err = file.CountFiles(sfs.rootFolderPath)
	if err != nil {
		return filesCount, err
	}

	return filesCount, nil
}

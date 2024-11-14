package sfs

import (
	"github.com/vault-thirteen/auxie/file"
)

// ListFileNames lists name of all files in the storage folder. Cache is not
// used while may not contain all the files.
func (sfs *SimpleFileServer) ListFileNames() (fileNames []string, err error) {
	fileNames, err = file.ListFileNames(sfs.rootFolderPath)
	if err != nil {
		return nil, err
	}

	return fileNames, nil
}

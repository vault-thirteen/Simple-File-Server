package sfs

import "github.com/vault-thirteen/auxie/file"

// getFileExistenceUsingCache gets the existence flag of a default file using
// cache.
func (sfs *SimpleFileServer) getFileExistenceUsingCache(absFilePath string) (fileExists bool, err error) {
	var ok bool
	fileExists, ok = sfs.fileExistenceMap[absFilePath]
	if !ok {
		fileExists, err = file.FileExists(absFilePath)
		if err != nil {
			return false, err
		}

		sfs.fileExistenceMap[absFilePath] = fileExists
	}

	return fileExists, nil
}

// getFileExistenceWithoutCache gets the existence flag of a default file not
// using cache.
func (sfs *SimpleFileServer) getFileExistenceWithoutCache(absFilePath string) (fileExists bool, err error) {
	return file.FileExists(absFilePath)
}

package sfs

import "github.com/vault-thirteen/auxie/file"

// getFileExistenceWithoutCache gets the existence flag of a default file not
// using cache.
func (sfs *SimpleFileServer) getFileExistenceWithoutCache(absFilePath string) (fileExists bool, err error) {
	return file.FileExists(absFilePath)
}

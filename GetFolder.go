package sfs

import "errors"

func (sfs *SimpleFileServer) GetFolder(folderRelPath string) (bytes []byte, err error) {
	if !IsPathValid(folderRelPath) {
		return nil, errors.New(Err_PathIsNotValid)
	}

	if !sfs.isCachingEnabled {
		return sfs.getFolderFromStorage(folderRelPath)
	}

	return sfs.getFolderFromStorage(folderRelPath)
}

func (sfs *SimpleFileServer) getFolderFromCache(folderRelPath string) (bytes []byte, err error) {
	return nil, errors.New(Err_ActionIsImpossible)
}

func (sfs *SimpleFileServer) getFolderFromStorage(folderRelPath string) (bytes []byte, err error) {
	var fileName string
	fileName, err = sfs.GetFolderDefaultFilename(folderRelPath)
	if err != nil {
		return nil, err
	}

	if len(fileName) == 0 {
		return nil, nil
	}

	return sfs.GetFile(fileName)
}

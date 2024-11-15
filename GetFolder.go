package sfs

// GetFolder returns contents of the first found default file of a folder.
// Path is a relative path inside the file server's root folder.
func (sfs *SimpleFileServer) GetFolder(folderRelPath string) (bytes []byte, err error) {
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

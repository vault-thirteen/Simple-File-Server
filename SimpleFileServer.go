package sfs

import (
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"

	cache "github.com/vault-thirteen/Cache"
	"github.com/vault-thirteen/auxie/file"
	"github.com/vault-thirteen/errorz"
)

const (
	ErrFolderDoesNotExist = "folder does not exist"
	ErrFilePathIsNotValid = "file path is not valid"
)

const (
	PathLevelUp = ".."
)

type SimpleFileServer struct {
	rootFolderPath string
	cache          *cache.Cache[string, []byte]
}

// NewSimpleFileServer is a constructor of a SimpleFileServer object.
func NewSimpleFileServer(
	rootFolderPath string,
	cacheSizeLimit int,
	cacheVolumeLimit int,
	cacheRecordTtl uint,
) (sfs *SimpleFileServer, err error) {
	var ok bool
	ok, err = file.FolderExists(rootFolderPath)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, errors.New(ErrFolderDoesNotExist)
	}
	sfs = &SimpleFileServer{
		rootFolderPath: rootFolderPath,
	}

	sfs.cache = cache.NewCache[string, []byte](cacheSizeLimit, cacheVolumeLimit, cacheRecordTtl)

	return sfs, nil
}

// GetFileContents returns file's contents.
// File path is a relative path inside the file server's root folder.
func (sfs *SimpleFileServer) GetFileContents(filePath string) (bytes []byte, fileExists bool, err error) {
	if !sfs.isFilePathValid(filePath) {
		return nil, false, errors.New(ErrFilePathIsNotValid)
	}

	filePath = filepath.Join(sfs.rootFolderPath, filePath)

	bytes, err = sfs.cache.GetRecord(filePath)
	if err == nil {
		// File is cached.
		return bytes, true, nil
	}

	// File is not cached.
	fileExists, err = file.FileExists(filePath)
	if err != nil {
		return nil, false, err
	}
	if !fileExists {
		return nil, false, nil
	}

	bytes, err = sfs.readFileFromOs(filePath)
	if err != nil {
		return nil, true, err
	}

	err = sfs.cache.AddRecord(filePath, bytes)
	if err != nil {
		return nil, true, err
	}

	return bytes, true, nil
}

// readFileFromOs reads file's contents from an operating system.
func (sfs *SimpleFileServer) readFileFromOs(filePath string) (bytes []byte, err error) {
	var f *os.File
	f, err = os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer func() {
		derr := f.Close()
		if derr != nil {
			err = errorz.Combine(err, derr)
		}
	}()

	bytes, err = io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

// isFilePathValid checks validity of the file path.
func (sfs *SimpleFileServer) isFilePathValid(filePath string) bool {
	if strings.Contains(filePath, PathLevelUp) {
		return false
	}
	return true
}

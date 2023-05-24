package sfs

import (
	"errors"

	cache "github.com/vault-thirteen/Cache"
	"github.com/vault-thirteen/auxie/file"
)

const (
	PathLevelUp = ".."
)

const (
	ErrFolderDoesNotExist = "folder does not exist"
	ErrPathIsNotValid     = "path is not valid"
)

type SimpleFileServer struct {
	rootFolderPath     string
	folderDefaultFiles []string
	isCachingEnabled   bool
	cache              *cache.Cache[string, []byte]

	// fileExistenceMap caches flags showing existence of a file.
	// Key: absolute path to file; Value: existence of the file.
	fileExistenceMap map[string]bool
}

// NewSimpleFileServer is a constructor of a SimpleFileServer object.
func NewSimpleFileServer(
	rootFolderPath string,
	folderDefaultFiles []string,
	isCachingEnabled bool,
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
		rootFolderPath:     rootFolderPath,
		folderDefaultFiles: folderDefaultFiles,
		isCachingEnabled:   isCachingEnabled,
	}

	if sfs.isCachingEnabled {
		sfs.cache = cache.NewCache[string, []byte](cacheSizeLimit, cacheVolumeLimit, cacheRecordTtl)
		sfs.fileExistenceMap = map[string]bool{}
	}

	return sfs, nil
}

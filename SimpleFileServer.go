package sfs

import (
	"errors"

	"github.com/vault-thirteen/Cache/VL"
	"github.com/vault-thirteen/auxie/file"
)

type SimpleFileServer struct {
	rootFolderPath     string
	folderDefaultFiles []string
	isCachingEnabled   bool

	// Cached contents of files.
	cache *vl.Cache[string, []byte]
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
		return nil, errors.New(Err_FolderDoesNotExist)
	}

	sfs = &SimpleFileServer{
		rootFolderPath:     rootFolderPath,
		folderDefaultFiles: folderDefaultFiles,
		isCachingEnabled:   isCachingEnabled,
	}

	if sfs.isCachingEnabled {
		sfs.cache = vl.NewCache[string, []byte](cacheSizeLimit, cacheVolumeLimit, cacheRecordTtl)
	}

	return sfs, nil
}

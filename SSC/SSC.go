package ssc

import "sync"

// SSC is a Simple Stupid Cache.
// This type of cache is used for very simple tasks.
type SSC struct {
	cache map[string]bool
	lock  sync.RWMutex

	// Maximum size of cache.
	// When it is <= 0, size limit is ignored.
	maxSize       int
	isSizeLimited bool
}

func NewSSC(maxSize int) (ssc *SSC) {
	return &SSC{
		cache:         map[string]bool{},
		lock:          sync.RWMutex{},
		maxSize:       maxSize,
		isSizeLimited: (maxSize > 0),
	}
}

// GetValue gets the cache record.
func (ssc *SSC) GetValue(key string) (value bool, recordExists bool) {
	ssc.lock.RLock()
	defer ssc.lock.RUnlock()

	// Golang does not allow to return map's value directly,
	// so this code is not compilable:
	// return ssc.cache[key]
	value, recordExists = ssc.cache[key]
	return value, recordExists
}

// SetValue sets the cache record value.
func (ssc *SSC) SetValue(key string, value bool) (recordAlreadyExists bool) {
	ssc.lock.Lock()
	defer ssc.lock.Unlock()

	_, recordAlreadyExists = ssc.cache[key]
	ssc.cache[key] = value

	// Check size for new records.
	if !recordAlreadyExists {
		ssc.checkSize()
	}

	return recordAlreadyExists
}

// checkSize removes extra cache records in case of size overflow.
func (ssc *SSC) checkSize() {
	if !ssc.isSizeLimited {
		return
	}

	if len(ssc.cache) <= ssc.maxSize {
		return
	}

	nRecordsToRemove := len(ssc.cache) - ssc.maxSize
	for i := 1; i <= nRecordsToRemove; i++ {
		ssc.removeRandomRecord()
	}
}

// removeRandomRecord removes a single random record from cache.
func (ssc *SSC) removeRandomRecord() {
	l := len(ssc.cache)
	if l == 0 {
		return
	}

	iX := (l - 1) / 2
	i := 0
	for k, _ := range ssc.cache {
		if i == iX {
			delete(ssc.cache, k)
			return
		}

		i++
	}
}

// GetSize returns the cache size.
func (ssc *SSC) GetSize() (cacheSize int) {
	ssc.lock.RLock()
	defer ssc.lock.RUnlock()

	return len(ssc.cache)
}

// Reset clears the cache.
func (ssc *SSC) Reset() {
	ssc.lock.Lock()
	defer ssc.lock.Unlock()

	ssc.cache = map[string]bool{}

	return
}

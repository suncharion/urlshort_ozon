package storagememory

import (
	"fmt"
	"ozon/internal/model"
	"sync"
)

type StorageMemory struct {
	data      map[string]*model.ShortenedUrl
	mutexData sync.RWMutex
}

func NewStorageMemory() (*StorageMemory, error) {
	return &StorageMemory{
		data: make(map[string]*model.ShortenedUrl),
	}, nil
}

func (sm *StorageMemory) Put(url *model.ShortenedUrl) error {
	sm.mutexData.Lock()
	defer sm.mutexData.Unlock()

	sm.data[url.Short] = url

	return nil
}

func (sm *StorageMemory) Get(shortUrl string) (*model.ShortenedUrl, error) {
	sm.mutexData.RLock()
	defer sm.mutexData.RUnlock()

	if val, ok := sm.data[shortUrl]; ok {
		return val, nil
	}

	return nil, fmt.Errorf("Not found")
}

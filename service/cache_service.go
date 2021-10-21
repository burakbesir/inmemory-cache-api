package service

import (
	"encoding/json"
	"errors"
	internalerror "github.com/rahmanbesir/inmemory-cache-api/error"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
)

const (
	cacheFileName               = "cache.json"
	cacheFilePath               = "tmp"
	cacheFileSaveIntervalSecond = 5
)

var fileName = filepath.Join(cacheFilePath, cacheFileName)

type CacheService interface {
	Put(key string, value string)
	Get(key string) (string, error)
}

type cacheService struct {
	cache *sync.Map
}

func NewCacheService() CacheService {
	cacheService := &cacheService{
		cache: loadCache(),
	}
	runCacheSaveJob(cacheService)
	return cacheService
}

func (c *cacheService) Put(key string, value string) {
	c.cache.Store(key, value)
}

func (c *cacheService) Get(key string) (string, error) {
	value, ok := c.cache.Load(key)
	if !ok {
		return "", internalerror.CreateNotFoundError(errors.New(key + " key doesn't exist"))
	}

	return value.(string), nil
}

func loadCache() *sync.Map {
	file, err := os.ReadFile(fileName)
	syncMap := &sync.Map{}
	if err != nil {
		return syncMap
	}

	m := make(map[string]string)
	_ = json.Unmarshal(file, &m)

	for k, v := range m {
		syncMap.Store(k, v)
	}
	return syncMap
}

func runCacheSaveJob(cacheService *cacheService) {
	err := os.MkdirAll(cacheFilePath, os.ModePerm)
	if err != nil {
		log.Printf(	"cache file path couldn't create : %v", err)
	} else {
		go func() {
			for range time.Tick(time.Second * cacheFileSaveIntervalSecond) {
				saveCache(cacheService.cache)
			}
		}()
	}
}

func saveCache(cache *sync.Map) {
	data, _ := json.Marshal(convertSyncMapToMap(cache))

	f, err := os.Create(fileName)

	if err != nil {
		log.Printf("cache file cannot create error : %v", err)
		return
	}

	defer f.Close()

	_, err = f.Write(data)

	if err != nil {
		log.Printf("cache file couldn't write error : %v", err)
	}
}

func convertSyncMapToMap(syncMap *sync.Map) map[string]string {
	m := make(map[string]string)
	syncMap.Range(func(key, value interface{}) bool {
		m[key.(string)] = value.(string)
		return true
	})
	return m
}

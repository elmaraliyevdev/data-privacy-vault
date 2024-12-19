package storage

import "sync"

type MemoryStore struct {
    sync.Mutex
    Data map[string]map[string]string
}

var Store = &MemoryStore{
    Data: make(map[string]map[string]string),
}
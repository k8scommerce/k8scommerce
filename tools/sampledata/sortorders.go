package main

import "sync"

type SortOrders struct {
	Category map[string]int64
	mutex    sync.Mutex
}

func (i *SortOrders) NextCategorySortOrder(path string) int64 {
	i.mutex.Lock()
	defer i.mutex.Unlock()
	if i.Category == nil {
		i.Category = make(map[string]int64)
	}

	i.Category[path]++
	return i.Category[path]
}

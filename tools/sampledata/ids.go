package main

import "sync"

type IDs struct {
	CategoryID int64
	ProductID  int64
	mutex      sync.Mutex
}

func (i *IDs) NextCategoryID() int64 {
	i.mutex.Lock()
	defer i.mutex.Unlock()
	i.CategoryID++
	return i.CategoryID
}

func (i *IDs) NextProductID() int64 {
	i.mutex.Lock()
	defer i.mutex.Unlock()
	i.ProductID++
	return i.ProductID
}

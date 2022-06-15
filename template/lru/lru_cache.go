package lru

import (
	"container/list"
	"errors"
)

type LRU struct {
	size      int
	innerList *list.List
	innerMap  map[int]*list.Element
}

type entry struct {
	key int
	val int
}

func New(size int) (*LRU, error) {
	if size <= 0 {
		return nil, errors.New("must provide a positive size")
	}

	return &LRU{
		size:      size,
		innerList: list.New(),
		innerMap:  make(map[int]*list.Element),
	}, nil
}

func (lru *LRU) Get(key int) (int, bool) {
	if e, ok := lru.innerMap[key]; ok {
		lru.innerList.MoveToFront(e)
		return e.Value.(*entry).val, true
	}
	return -1, false
}

func (lru *LRU) Put(key int, value int) bool {
	if e, ok := lru.innerMap[key]; ok {
		lru.innerList.MoveToFront(e)
		e.Value.(*entry).val = value
		return false
	}

	e := lru.innerList.PushFront(
		&entry{key: key, val: value},
	)
	lru.innerMap[key] = e

	if lru.innerList.Len() > lru.size {
		delete(lru.innerMap, lru.innerList.Back().Value.(*entry).key)
		lru.innerList.Remove(lru.innerList.Back())
		return true
	}
	return false
}

package mcache

import "sync"

type Element struct {
	prev  *Element
	next  *Element
	Key   interface{}
	Value interface{}
}

func (e *Element) Next() *Element {
	return e.next
}

func (e *Element) Prev() *Element {
	return e.prev
}

type LRUCache struct {
	lock     *sync.Mutex
	cache    map[interface{}]*Element
	head     *Element
	tail     *Element
	capacity int
}

func New(capacity int) *LRUCache {
	return &LRUCache{new(sync.Mutex), make(map[interface{}]*Element), nil, nil, capacity}
}

func (lc *LRUCache) Set(key interface{}, value interface{}) {
	lc.lock.Lock()
	defer func() {
		lc.lock.Unlock()
	}()
	//检查是否缓存中存在这个键，
	// 1.存在，则更新为新的值，并且刷新记录最近使用的key的链表状态
	if e, ok := lc.cache[key]; ok {
		e.Value = value
		lc.refresh(e)
		return
	}
	//不存在先检查，缓存中是否还有空间存放
	if lc.capacity == 0 {
		return
	} else if len(lc.cache) >= lc.capacity {
		//空间不足，删除在记录最近使用的链表尾部的元素
		delete(lc.cache, lc.tail.Key)
		lc.remove(lc.tail)
	}
	e := &Element{nil, lc.head, key, value}
	lc.cache[key] = e
	//长度为一的时候，直接添加到链表尾部，不是从链表头部添加
	if len(lc.cache) != 1 {
		lc.head.prev = e
	} else {
		lc.tail = e
	}
	lc.head = e
}

func (lc *LRUCache) Get(key interface{}) (interface{}, bool) {
	lc.lock.Lock()
	defer func() {
		lc.lock.Unlock()
	}()
	if e, ok := lc.cache[key]; ok {
		//取出元素后刷新列表
		lc.refresh(e)
		return e.Value, ok
	}else{
		return nil, ok
	}
}

func (lc *LRUCache) Delete(key interface{}) {
	lc.lock.Lock()
	defer func() {
		lc.lock.Unlock()
	}()
	if e, ok := lc.cache[key]; ok {
		//删除对应的键，以及删除最近使用的键的链表中删除对应的元素
		delete(lc.cache, key)
		lc.remove(e)
	}
}

func (lc *LRUCache) Range(f func(key, value interface{}) bool) {
	lc.lock.Lock()
	defer func() {
		lc.lock.Unlock()
	}()
	for i := lc.head; i != nil; i = i.Next() {
		if !f(i.Key, i.Value) {
			break
		}
	}
}

func (lc *LRUCache) UpdateElementValue(key interface{}, f func(value *interface{})) {
	lc.lock.Lock()
	defer func() {
		lc.lock.Unlock()
	}()
	if e, ok := lc.cache[key]; ok {
		f(&e.Value)
		lc.refresh(e)
	}
}

func (lc *LRUCache) Front() *Element {
	return lc.head
}

func (lc *LRUCache) Back() *Element {
	return lc.tail
}

func (lc *LRUCache) Len() int {
	return len(lc.cache)
}

func (lc *LRUCache) Capacity() int {
	return lc.capacity
}

func (lc *LRUCache) refresh(e *Element) {
	if e.prev != nil {
		//将最近使用的元素从对应位置转到链表头
		e.prev.next = e.next
		if e.next == nil {
			lc.tail = e.prev
		} else {
			e.next.prev = e.prev
		}
		e.prev = nil
		e.next = lc.head
		lc.head.prev = e
		lc.head = e
	}
}

func (lc *LRUCache) remove(e *Element) {
	//判断是否链表头
	if e.prev == nil {
		lc.head = e.next
	} else {
		e.prev.next = e.next
	}
	//判断是否是链表尾部
	if e.next == nil {
		lc.tail = e.prev
	} else {
		e.next.prev = e.prev
	}
}

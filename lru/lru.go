//go:build !solution

package lrucache

import (
	"container/list"
)

type lruCache struct {
	//data            map[int]int
	accessTimeOrder *list.List
	elems           map[int]*list.Element
	cap             int
}

type pair struct {
	key, value int
}

// Get returns value associated with the key.
//
// The second value is a bool that is true if the key exists in the cache,
// and false if not.
func (l *lruCache) Get(key int) (int, bool) {
	val, ok := l.elems[key]
	if ok {
		l.accessTimeOrder.MoveToBack(val)
		return val.Value.(*pair).value, ok
	}
	return 0, false
}

// Set updates value associated with the key.
//
// If there is no key in the cache new (key, value) pair is created.
func (l *lruCache) Set(key, value int) {
	currentElement, ok := l.elems[key]
	if ok {
		l.accessTimeOrder.MoveToBack(currentElement)
		ce := currentElement.Value.(*pair)
		ce.value = value
	} else {
		curr := l.accessTimeOrder.PushBack(&pair{key, value})
		l.elems[key] = curr
		if len(l.elems) > l.cap {
			earliest := l.accessTimeOrder.Front()
			earlyKey := earliest.Value.(*pair).key
			delete(l.elems, earlyKey)
			l.accessTimeOrder.Remove(earliest)
		}
	}

}

// Range calls function f on all elements of the cache
// in increasing access time order.
//
// Stops earlier if f returns false.
func (l *lruCache) Range(f func(key, value int) bool) {
	flag := true
	for e := l.accessTimeOrder.Front(); e != nil && flag; e = e.Next() {
		p := e.Value.(*pair)
		flag = f(p.key, p.value)
	}
}

// Clear removes all keys and values from the cache.
func (l *lruCache) Clear() {
	l.accessTimeOrder = list.New()
	for k := range l.elems {
		delete(l.elems, k)
	}
}

func New(cap int) Cache {
	return &lruCache{list.New(), make(map[int]*list.Element, cap), cap}
}

/*

package lrucache

import (
	"container/list"
)

type lruCache struct {
	//data            map[int]int
	accessTimeOrder *list.List
	elems           map[int]*list.Element
	cap             int
}

type pair struct {
	key, value int
}

// Get returns value associated with the key.
//
// The second value is a bool that is true if the key exists in the cache,
// and false if not.
func (l *lruCache) Get(key int) (int, bool) {
	val, ok := l.elems[key]
	if ok {
		l.accessTimeOrder.MoveToBack(l.elems[key])
		return val.Value.(int), ok
	}
	return 0, false
}

// Set updates value associated with the key.
//
// If there is no key in the cache new (key, value) pair is created.
func (l *lruCache) Set(key, value int) {
	currentElement, ok := l.elems[key]
	if ok {
		l.accessTimeOrder.MoveToBack(currentElement)
		currentElement.Value = value
	} else {
		curr := l.accessTimeOrder.PushBack(value)
		l.elems[key] = curr
		if len(l.elems) > l.cap {
			earliest := l.accessTimeOrder.Front()
			earlyKey := earliest.Value.(pair).key
			delete(l.elems, earlyKey)
			l.accessTimeOrder.Remove(earliest)
		}
	}

}

// Range calls function f on all elements of the cache
// in increasing access time order.
//
// Stops earlier if f returns false.
func (l *lruCache) Range(f func(key, value int) bool) {
	flag := true
	for e := l.accessTimeOrder.Front(); e != nil && flag; e = e.Next() {
		p := e.Value.(pair)
		flag = f(p.key, p.value)
	}
}

// Clear removes all keys and values from the cache.
func (l *lruCache) Clear() {
	l.accessTimeOrder = list.New()
	for k := range l.elems {
		delete(l.elems, k)
	}
}

func New(cap int) Cache {
	return &lruCache{list.New(), make(map[int]*list.Element, cap), cap}
}

*/

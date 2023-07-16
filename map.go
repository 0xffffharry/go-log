package log

import (
	"sync"
	"sync/atomic"
)

type syncMap[K comparable, V any] struct {
	sync.Map
}

func (m *syncMap[K, V]) Load(key K) (value V, ok bool) {
	var v any
	v, ok = m.Map.Load(key)
	if ok && v != nil {
		value = v.(V)
	}
	return
}

func (m *syncMap[K, V]) Store(key K, value V) {
	m.Map.Store(key, value)
}

func (m *syncMap[K, V]) Delete(key K) {
	m.Map.Delete(key)
}

func (m *syncMap[K, V]) Range(f func(key K, value V) bool) {
	m.Map.Range(func(keyAny any, valueAny any) bool {
		var key K
		var value V
		if keyAny != nil {
			key = keyAny.(K)
		}
		if valueAny != nil {
			value = valueAny.(V)
		}
		return f(key, value)
	})
}

func (m *syncMap[K, V]) GoRange(f func(key K, value V)) {
	var wg sync.WaitGroup
	m.Map.Range(func(keyAny any, valueAny any) bool {
		var key K
		var value V
		if keyAny != nil {
			key = keyAny.(K)
		}
		if valueAny != nil {
			value = valueAny.(V)
		}
		wg.Add(1)
		go func(key K, value V) {
			defer wg.Done()
			f(key, value)
		}(key, value)
		return true
	})
	wg.Wait()
}

func (m *syncMap[K, V]) LoadOrStore(key K, value V) (actual V, loaded bool) {
	var v any
	v, loaded = m.Map.LoadOrStore(key, value)
	if loaded && v != nil {
		actual = v.(V)
	}
	return
}

func (m *syncMap[K, V]) LoadAndDelete(key K) (value V, loaded bool) {
	var v any
	v, loaded = m.Map.LoadAndDelete(key)
	if loaded && v != nil {
		value = v.(V)
	}
	return
}

func (m *syncMap[K, V]) Swap(key K, old V, new V) (swapped bool) {
	return m.Map.CompareAndSwap(key, old, new)
}

func (m *syncMap[K, V]) CompareAndSwap(key K, old V, new V) (swapped bool) {
	return m.Map.CompareAndSwap(key, old, new)
}

func (m *syncMap[K, V]) CompareAndDelete(key K, old V) (deleted bool) {
	return m.Map.CompareAndDelete(key, old)
}

func (m *syncMap[K, V]) Len() int {
	var length atomic.Int64
	m.GoRange(func(_ K, _ V) {
		length.Add(1)
	})
	lengthInt := int(length.Load())
	return lengthInt
}

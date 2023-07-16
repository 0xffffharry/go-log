package log

import "sync"

type syncPool[T any] struct {
	sync.Pool
}

func (p *syncPool[T]) New(f func() T) {
	p.Pool.New = func() any {
		return f()
	}
}

func (p *syncPool[T]) Get() (v T) {
	value := p.Pool.Get()
	if value != nil {
		v = value.(T)
	}
	return
}

func (p *syncPool[T]) Put(v T) {
	p.Pool.Put(v)
}

package log

import (
	"context"
	"sync"
)

type LimitTask struct {
	queue  chan struct{}
	wg     sync.WaitGroup
	once   sync.Once
	ctx    context.Context
	cancel context.CancelFunc
}

func NewLimitTask(size int) *LimitTask {
	if size <= 0 {
		size = 1
	}
	t := &LimitTask{
		queue: make(chan struct{}, size),
	}
	for i := 0; i < size; i++ {
		t.queue <- struct{}{}
	}
	return t
}

func (t *LimitTask) do(ctx context.Context, f func(ctx context.Context)) {
	<-t.queue
	defer func() {
		t.queue <- struct{}{}
	}()
	f(ctx)
}

func (t *LimitTask) Do(f func(ctx context.Context)) {
	t.once.Do(func() {
		t.ctx, t.cancel = context.WithCancel(context.Background())
	})
	t.wg.Add(1)
	go func() {
		defer t.wg.Done()
		t.do(t.ctx, f)
	}()
}

func (t *LimitTask) Cancel() {
	t.cancel()
}

func (t *LimitTask) Wait() {
	t.wg.Wait()
}

package log

import (
	"context"
	"math/rand"
	"time"
)

type contextKey struct{}

type ContextMessage interface {
	GetID() string
	GetAddTime() time.Time
}

type contextMessage struct {
	id      string
	addTime time.Time
}

var randomNum = "0123456789"

func NewDefaultContextMessage() ContextMessage {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var id string
	for i := 0; i < 8; i++ {
		id += string(randomNum[r.Intn(len(randomNum))])
	}
	return &contextMessage{
		id:      id,
		addTime: time.Now(),
	}
}

func (m *contextMessage) GetID() string {
	return m.id
}

func (m *contextMessage) GetAddTime() time.Time {
	return m.addTime
}

func WithContextMessage(ctx context.Context, message ContextMessage) context.Context {
	return context.WithValue(ctx, (*contextKey)(nil), message)
}

func WithDefaultContextMessage(ctx context.Context) context.Context {
	return WithContextMessage(ctx, NewDefaultContextMessage())
}

func GetContextMessage(ctx context.Context) ContextMessage {
	value := ctx.Value((*contextKey)(nil))
	if value == nil {
		return nil
	}
	message, ok := value.(ContextMessage)
	if !ok {
		return nil
	}
	return message
}

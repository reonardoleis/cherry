package state

import (
	"github.com/reonardoleis/cherry/pkg/router"
)

type state[T any] struct {
	data *T
}

type SetFunc[T any] func(T)

func SetState[T any](initialValue T) (*T, SetFunc[T]) {
	state := &state[T]{data: &initialValue}

	return state.data, state.Set
}

func (s state[T]) Get() T {
	return *s.data
}

func (s state[T]) Set(newValue T) {
	*s.data = newValue
	router.Instance().UpdateDOM()
}

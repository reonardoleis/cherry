package state

import (
	"github.com/reonardoleis/cherry/pkg/dom"
)

type state[T any] struct {
	name string
	data *T
}

type SetFunc[T any] func(T)

func SetState[T any](name string, initialValue T) (*T, SetFunc[T]) {
	state := &state[T]{name: name, data: &initialValue}

	return state.data, state.Set
}

func (s state[T]) Get() T {
	return *s.data
}

func (s state[T]) Set(newValue T) {
	*s.data = newValue
	dom.UpdateDOM()
}

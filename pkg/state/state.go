package state

import (
	"github.com/reonardoleis/cherry/pkg/router"
)

type State[T any] struct {
	data    *T
	boundTo []string
}

type SetFunc[T any] func(T)

func SetState[T any](initialValue T) *State[T] {
	state := &State[T]{data: &initialValue}
	return state
}

func (s State[T]) Get() T {
	return *s.data
}

func (s *State[T]) Set(newValue T) {
	*s.data = newValue
	for _, bind := range s.boundTo {
		router.Instance().UpdateBinds(bind, *s.data)
	}
}

func (s *State[T]) Bind(key string) {
	s.boundTo = append(s.boundTo, key)
}

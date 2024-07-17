package component

import "reflect"

type Component[T any] interface {
	Render() string
	Children() map[string][]Component[T]
}

type ComponentFunction = string

type Base[T any] struct {
	children map[string][]Component[T]
}

func (b *Base[T]) Register(component Component[T]) {
	if b.children == nil {
		b.children = make(map[string][]Component[T])
	}

	name := reflect.TypeOf(component).Name()

	v, ok := b.children[name]
	if !ok {
		b.children[name] = []Component[T]{component}
	} else {
		v = append(v, component)
	}
}

func (b Base[T]) Children() map[string][]Component[T] {
	return b.children
}

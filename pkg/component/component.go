package component

type Component[T any] interface {
	Render() string
}

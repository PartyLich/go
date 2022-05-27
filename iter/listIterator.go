package iter

import "container/list"

// ListIterator is a lazy iterator over a container/list
type ListIterator[T any] struct {
	list *list.List
	el   *list.Element
}

// New creates a new lazy iterator over the provided container/list
func FromList[T any](list *list.List) *ListIterator[T] {
	return &ListIterator[T]{list, list.Front()}
}

// Next advances the iterator and returns the next value.
//
// Returns nil when iteration is finished, or if the next value does not conform
// to the specified type.
func (it *ListIterator[T]) Next() *T {
	if it.el == nil {
		return nil
	}

	el, ok := it.el.Value.(T)
	if !ok {
		return nil
	}

	it.el = it.el.Next()

	return &el
}

//go:generate go run ./cmd/gen/ -name ListIterator -output listIterator_ext_gen.go

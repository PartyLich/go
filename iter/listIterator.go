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

// Find searches for an element of an iterator that satisfies a predicate.
//
// Takes a function that returns true or false. It applies this function to
// each element of the iterator, and if any of them return true, then Find
// returns a pointer to the element. If they all return false, it returns
// nil.
//
// Find is short-circuiting; in other words, it will stop processing as soon as
// the predicate returns true.
func (it *ListIterator[T]) Find(pred func(T) bool) *T {
	for next := it.Next(); next != nil; next = it.Next() {
		if pred(*next) {
			return next
		}
	}

	return nil
}

//go:generate go run ./cmd/gen/ -name ListIterator -output listIterator_ext_gen.go

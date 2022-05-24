package iter

type Iterator[T any] struct {
	idx   int
	slice []T
}

// New creates a new lazy iterator over the provided slice
func New[T any](slice []T) *Iterator[T] {
	return &Iterator[T]{0, slice}
}

// Next advances the iterator and returns the next value.
//
// Returns nil when iteration is finished.
func (iter *Iterator[T]) Next() *T {
	if iter.idx >= len(iter.slice) {
		return nil
	}

	next := &iter.slice[iter.idx]
	iter.idx += 1

	return next
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
func (iter *Iterator[T]) Find(pred func(T) bool) *T {
	for iter.idx < len(iter.slice) {
		next := iter.slice[iter.idx]
		iter.idx++

		if pred(next) {
			return &next
		}
	}

	return nil
}

// Rev reverses the iteration order of this iterator
func (iter *Iterator[T]) Rev() *RevIterator[T] {
	var idx int

	if iter.idx == len(iter.slice)-1 {
		idx = 0
	}
	if iter.idx == 0 {
		idx = len(iter.slice) - 1
	}

	return &RevIterator[T]{
		Iterator[T]{idx, iter.slice},
	}
}

//go:generate go run ./cmd/gen/ -name Iterator -output iterator_ext_gen.go

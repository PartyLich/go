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

package iter

// An Iterable with the direction reversed.
type RevIterator[T any] struct {
	it Iterator[T]
}

// Next advances the iterator and returns the next value.
//
// Returns nil when iteration is finished.
func (iter *RevIterator[T]) Next() *T {
	if iter.it.idx < 0 {
		return nil
	}

	next := &iter.it.slice[iter.it.idx]
	iter.it.idx -= 1

	return next
}

//go:generate go run ./cmd/gen/ -name RevIterator -output revIterator_ext_gen.go

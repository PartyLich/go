package iter

type RevIterator[T any] struct {
	it Iterator[T]
}

func (iter *RevIterator[T]) Next() *T {
	if iter.it.idx < 0 {
		return nil
	}

	next := &iter.it.slice[iter.it.idx]
	iter.it.idx -= 1

	return next
}

// Find searches for an element of an iterator that satisfies a predicate.
//
// takes a function that returns `true` or `false`. It applies this function to
// each element of the iterator, and if any of them return `true`, then Find
// returns a pointer to the element. If they all return `false`, it returns
// `nil`.
//
// Find is short-circuiting; in other words, it will stop processing as soon as
// the closure returns `true`.
func (iter *RevIterator[T]) Find(pred func(T) bool) *T {
	for iter.it.idx >= 0 {
		next := iter.it.slice[iter.it.idx]
		iter.it.idx--

		if pred(next) {
			return &next
		}
	}

	return nil
}

//go:generate go run ./cmd/gen/ -name RevIterator -output revIterator_ext_gen.go

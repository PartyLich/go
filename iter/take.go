package iter

// Taken is an iterator that only iterates over the first n elements.
type Taken[T any] struct {
	iter Iterable[T]
	n    int
}

// Take creates an iterator that yields the first n elements, or fewer if the
// underlying iterator ends sooner.
func Take[T any](iter Iterable[T], n int) *Taken[T] {
	return &Taken[T]{iter, n}
}

func (s *Taken[T]) Next() *T {
	if s.n == 0 {
		return nil
	}

	s.n -= 1
	return s.iter.Next()
}

func (iter *Taken[T]) Find(pred func(T) bool) *T {
	for next := iter.Next(); next != nil; next = iter.Next() {
		if pred(*next) {
			return next
		}
	}

	return nil
}

//go:generate go run ./cmd/gen/ -name Taken -output take_ext_gen.go

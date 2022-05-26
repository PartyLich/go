package iter

// Skipped is an iterator that skips over n elements.
type Skipped[T any] struct {
	iter Iterable[T]
	n    int
}

// Skip creates an iterator that skips the first n elements.
func Skip[T any](iter Iterable[T], n int) *Skipped[T] {
	if n < 0 {
		panic("Skip requires n >= 0")
	}

	return &Skipped[T]{iter, n}
}

// Next advances the iterator and returns the next value.
//
// Returns nil when iteration is finished.
func (s *Skipped[T]) Next() *T {
	for s.n != 0 {
		s.n -= 1
		s.iter.Next()
	}

	return s.iter.Next()
}

//go:generate go run ./cmd/gen/ -name Skipped -output skip_ext_gen.go

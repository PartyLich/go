package iter

type Chained[T any] struct {
	a, b Iterable[T]
}

// Chain takes two iterators and creates a new iterator over both in sequence.
//
// Chain will return a new iterator which will first iterate over values from
// the first iterator and then over values from the second iterator.
func Chain[T any](a, b Iterable[T]) *Chained[T] {
	return &Chained[T]{a, b}
}

func (c *Chained[T]) Next() *T {
	next := c.a.Next()

	if next == nil {
		return c.b.Next()
	}

	return next
}

//go:generate go run ./cmd/gen/ -name Chained -output chain_ext_gen.go

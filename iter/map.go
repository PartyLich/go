package iter

type Mapped[T any, O any] struct {
	iter Iterable[T]
	fn   func(T) O
}

// Map returns an iterator that applies a function to every element.
func Map[T any, O any](iter Iterable[T], fn func(T) O) *Mapped[T, O] {
	return &Mapped[T, O]{iter, fn}
}

func (m *Mapped[T, O]) Next() *O {
	next := m.iter.Next()

	if next == nil {
		return nil
	}
	result := m.fn(*next)

	return &result
}

func (iter *Mapped[T, O]) Find(pred func(O) bool) *O {
	for next := iter.Next(); next != nil; next = iter.Next() {
		if pred(*next) {
			return next
		}
	}

	return nil
}

//go:generate go run ./cmd/gen/ -name Mapped -otype O -output map_ext_gen.go

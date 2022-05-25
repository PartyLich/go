package iter

type Filtered[T any] struct {
	iter Iterable[T]
	pred func(T) bool
}

// Filter returns an iterator which uses a predicate function to determine if an
// element should be yielded.
//
// The returned iterator will yield only the elements for which the predicate
// returns true.
func Filter[T any](iter Iterable[T], pred func(T) bool) *Filtered[T] {
	return &Filtered[T]{iter, pred}
}

func (f *Filtered[T]) Next() *T {
	next := f.iter.Find(f.pred)

	if next == nil || !f.pred(*next) {
		return nil
	}

	return next
}

//go:generate go run ./cmd/gen/ -name Filtered -output filter_ext_gen.go

package iter

// Flat is an Iterable that flattens one level of nesting in an Iterable of Iteraables
type Flat[I any] struct {
	outer Iterable[Iterable[I]]
	inner *Iterable[I]
}

// Flatten creates an iterator that flattens nested structure.
func Flatten[I any](it Iterable[Iterable[I]]) *Flat[I] {
	return &Flat[I]{it, it.Next()}
}

// Next advances the iterator and returns the next value.
//
// Returns nil when iteration is finished.
func (f *Flat[T]) Next() *T {
	for {
		if f.inner == nil {
			return nil
		}

		next := (*f.inner).Next()
		if next == nil {
			f.inner = f.outer.Next()
			continue
		}

		return next
	}
}

//go:generate go run ./cmd/gen/ -name Flat -output flatten_ext_gen.go

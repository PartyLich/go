package iter

// SkipWhile is an Iterable that rejects elements while predicate returns true.
type SkipWhileT[T any] struct {
	iter Iterable[T]
	flag bool
	pred func(T) bool
}

// SkipWhile creates an iterator that skips elements based on a predicate.
func SkipWhile[T any](iter Iterable[T], pred func(T) bool) *SkipWhileT[T] {
	return &SkipWhileT[T]{iter, false, pred}
}

// Next advances the iterator and returns the next value.
//
// Returns nil when iteration is finished.
func (s *SkipWhileT[T]) Next() *T {
	check := func(flag *bool, pred func(T) bool) func(T) bool {
		return func(t T) bool {
			if *flag || !pred(t) {
				*flag = true
				return true
			} else {
				return false
			}
		}
	}

	return s.iter.Find(check(&s.flag, s.pred))
}

//go:generate go run ./cmd/gen/ -name SkipWhileT -output skipWhile_ext_gen.go

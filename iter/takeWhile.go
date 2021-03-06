package iter

// TakeWhile is an Iterable that only yields elements while a predicate returns true.
type TakeWhileT[T any] struct {
	iter Iterable[T]
	flag bool
	pred func(T) bool
}

// TakeWhile Creates an iterator that yields elements based on a predicate.
//
// TakeWhile takes a predicate function as an argument. It will call this
// function on each element of the iterator, and yield elements while it returns
// true.
func TakeWhile[T any](iter Iterable[T], pred func(T) bool) *TakeWhileT[T] {
	return &TakeWhileT[T]{iter, true, pred}
}

// Next advances the iterator and returns the next value.
//
// Returns nil when iteration is finished.
func (s *TakeWhileT[T]) Next() *T {
	check := func(flag *bool, pred func(T) bool) func(T) bool {
		return func(t T) bool {
			if *flag && pred(t) {
				return true
			} else {
				*flag = false
				return false
			}
		}
	}

	return s.iter.Find(check(&s.flag, s.pred))
}

//go:generate go run ./cmd/gen/ -name TakeWhileT -output takeWhile_ext_gen.go

package iter

// SkipWhile iterable adapter
type SkipWhileT[T any] struct {
	iter Iterable[T]
	flag bool
	pred func(T) bool
}

// SkipWhile creates an iterator that skips elements based on a predicate.
func SkipWhile[T any](iter Iterable[T], pred func(T) bool) *SkipWhileT[T] {
	return &SkipWhileT[T]{iter, false, pred}
}

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

func (iter *SkipWhileT[T]) Find(pred func(T) bool) *T {
	for next := iter.Next(); next != nil; next = iter.Next() {
		if pred(*next) {
			return next
		}
	}

	return nil
}

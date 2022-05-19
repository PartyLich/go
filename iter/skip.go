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

func (s *Skipped[T]) Next() *T {
	for s.n != 0 {
		s.n -= 1
		s.iter.Next()
	}

	return s.iter.Next()
}

func (iter *Skipped[T]) Find(pred func(T) bool) *T {
	for next := iter.Next(); next != nil; next = iter.Next() {
		if pred(*next) {
			return next
		}
	}

	return nil
}

func (iter *Skipped[T]) Count() int {
	return Count[T](iter)
}

func (iter *Skipped[T]) Partition(pred func(T) bool) ([]T, []T) {
	return Partition[T](iter, pred)
}

func (iter *Skipped[T]) Filter(pred func(T) bool) *Filtered[T] {
	return Filter[T](iter, pred)
}

func (iter *Skipped[T]) SkipWhile(pred func(T) bool) *SkipWhileT[T] {
	return SkipWhile[T](iter, pred)
}

func (iter *Skipped[T]) TakeWhile(pred func(T) bool) *TakeWhileT[T] {
	return TakeWhile[T](iter, pred)
}

func (iter *Skipped[T]) Chain(b Iterable[T]) *Chained[T] {
	return Chain[T](iter, b)
}

func (iter *Skipped[T]) StepBy(step int) *Stepped[T] {
	return StepBy[T](iter, step)
}

func (iter *Skipped[T]) Take(n int) *Taken[T] {
	return Take[T](iter, n)
}

func (iter *Skipped[T]) Collect() []T {
	return Collect[T](iter)
}

func (iter *Skipped[T]) ForEach(fn func(T)) {
	ForEach[T](iter, fn)
}

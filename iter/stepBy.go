package iter

type Stepped[T any] struct {
	iter  Iterable[T]
	step  int
	first bool
}

// StepBy creates an iterator starting at the same point, but stepping by the
// given amount at each iteration.
//
// The method will panic if the given step is <= 0.
//
// Note 1: The first element of the iterator will always be returned, regardless
// of the step given.
func StepBy[T any](a Iterable[T], step int) *Stepped[T] {
	if step <= 0 {
		panic("StepBy requires a step value > 0")
	}

	return &Stepped[T]{a, step, true}
}

func (s *Stepped[T]) Next() *T {
	if s.first {
		s.first = false
		return s.iter.Next()
	}

	next := s.iter.Next()
	for i := 1; i < s.step; i++ {
		next = s.iter.Next()
	}

	return next
}

func (s *Stepped[T]) Find(pred func(T) bool) *T {
	var next *T

	for next = s.Next(); next != nil; next = s.Next() {
		if pred(*next) {
			break
		}
	}

	return next
}

func (iter *Stepped[T]) Count() int {
	return Count[T](iter)
}

func (iter *Stepped[T]) Partition(pred func(T) bool) ([]T, []T) {
	return Partition[T](iter, pred)
}

func (iter *Stepped[T]) Filter(pred func(T) bool) *Filtered[T] {
	return Filter[T](iter, pred)
}

func (iter *Stepped[T]) SkipWhile(pred func(T) bool) *SkipWhileT[T] {
	return SkipWhile[T](iter, pred)
}

func (iter *Stepped[T]) TakeWhile(pred func(T) bool) *TakeWhileT[T] {
	return TakeWhile[T](iter, pred)
}

func (iter *Stepped[T]) Chain(b Iterable[T]) *Chained[T] {
	return Chain[T](iter, b)
}

func (iter *Stepped[T]) Skip(n int) *Skipped[T] {
	return Skip[T](iter, n)
}

func (iter *Stepped[T]) Take(n int) *Taken[T] {
	return Take[T](iter, n)
}

func (iter *Stepped[T]) Collect() []T {
	return Collect[T](iter)
}

func (iter *Stepped[T]) ForEach(fn func(T)) {
	ForEach[T](iter, fn)
}

func (iter *Stepped[T]) Nth(n int) *T {
	return Nth[T](iter, n)
}

func (iter *Stepped[T]) All(pred func(T) bool) bool {
	return All[T](iter, pred)
}

func (iter *Stepped[T]) Any(pred func(T) bool) bool {
	return Any[T](iter, pred)
}

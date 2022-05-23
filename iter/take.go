package iter

// Taken is an iterator that only iterates over the first n elements.
type Taken[T any] struct {
	iter Iterable[T]
	n    int
}

// Take creates an iterator that yields the first n elements, or fewer if the
// underlying iterator ends sooner.
func Take[T any](iter Iterable[T], n int) *Taken[T] {
	return &Taken[T]{iter, n}
}

func (s *Taken[T]) Next() *T {
	if s.n == 0 {
		return nil
	}

	s.n -= 1
	return s.iter.Next()
}

func (iter *Taken[T]) Find(pred func(T) bool) *T {
	for next := iter.Next(); next != nil; next = iter.Next() {
		if pred(*next) {
			return next
		}
	}

	return nil
}

func (iter *Taken[T]) Count() int {
	return Count[T](iter)
}

func (iter *Taken[T]) Partition(pred func(T) bool) ([]T, []T) {
	return Partition[T](iter, pred)
}

func (iter *Taken[T]) Filter(pred func(T) bool) *Filtered[T] {
	return Filter[T](iter, pred)
}

func (iter *Taken[T]) SkipWhile(pred func(T) bool) *SkipWhileT[T] {
	return SkipWhile[T](iter, pred)
}

func (iter *Taken[T]) TakeWhile(pred func(T) bool) *TakeWhileT[T] {
	return TakeWhile[T](iter, pred)
}

func (iter *Taken[T]) Chain(b Iterable[T]) *Chained[T] {
	return Chain[T](iter, b)
}

func (iter *Taken[T]) StepBy(step int) *Stepped[T] {
	return StepBy[T](iter, step)
}

func (iter *Taken[T]) Skip(n int) *Skipped[T] {
	return Skip[T](iter, n)
}

func (iter *Taken[T]) Collect() []T {
	return Collect[T](iter)
}

func (iter *Taken[T]) ForEach(fn func(T)) {
	ForEach[T](iter, fn)
}

func (iter *Taken[T]) Nth(n int) *T {
	return Nth[T](iter, n)
}

func (iter *Taken[T]) All(pred func(T) bool) bool {
	return All[T](iter, pred)
}

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

func (c *Chained[T]) Find(pred func(T) bool) *T {
	for next := c.a.Next(); next != nil; next = c.a.Next() {
		if pred(*next) {
			return next
		}
	}
	for next := c.b.Next(); next != nil; next = c.b.Next() {
		if pred(*next) {
			return next
		}
	}

	return nil
}

func (iter *Chained[T]) Count() int {
	return Count[T](iter)
}

func (iter *Chained[T]) Partition(pred func(T) bool) ([]T, []T) {
	return Partition[T](iter, pred)
}

func (iter *Chained[T]) Filter(pred func(T) bool) *Filtered[T] {
	return Filter[T](iter, pred)
}

func (iter *Chained[T]) SkipWhile(pred func(T) bool) *SkipWhileT[T] {
	return SkipWhile[T](iter, pred)
}

func (iter *Chained[T]) TakeWhile(pred func(T) bool) *TakeWhileT[T] {
	return TakeWhile[T](iter, pred)
}

func (iter *Chained[T]) Chain(b Iterable[T]) *Chained[T] {
	return Chain[T](iter, b)
}

func (iter *Chained[T]) StepBy(step int) *Stepped[T] {
	return StepBy[T](iter, step)
}

func (iter *Chained[T]) Skip(n int) *Skipped[T] {
	return Skip[T](iter, n)
}

func (iter *Chained[T]) Take(n int) *Taken[T] {
	return Take[T](iter, n)
}

func (iter *Chained[T]) Collect() []T {
	return Collect[T](iter)
}

func (iter *Chained[T]) ForEach(fn func(T)) {
	ForEach[T](iter, fn)
}

func (iter *Chained[T]) Nth(n int) *T {
	return Nth[T](iter, n)
}

func (iter *Chained[T]) All(pred func(T) bool) bool {
	return All[T](iter, pred)
}

func (iter *Chained[T]) Any(pred func(T) bool) bool {
	return Any[T](iter, pred)
}

package iter

type RevIterator[T any] struct {
	it Iterator[T]
}

func (iter *RevIterator[T]) Next() *T {
	if iter.it.idx < 0 {
		return nil
	}

	next := &iter.it.slice[iter.it.idx]
	iter.it.idx -= 1

	return next
}

// Find searches for an element of an iterator that satisfies a predicate.
//
// takes a function that returns `true` or `false`. It applies this function to
// each element of the iterator, and if any of them return `true`, then Find
// returns a pointer to the element. If they all return `false`, it returns
// `nil`.
//
// Find is short-circuiting; in other words, it will stop processing as soon as
// the closure returns `true`.
func (iter *RevIterator[T]) Find(pred func(T) bool) *T {
	for iter.it.idx >= 0 {
		next := iter.it.slice[iter.it.idx]
		iter.it.idx--

		if pred(next) {
			return &next
		}
	}

	return nil
}

func (iter *RevIterator[T]) Count() int {
	return Count[T](iter)
}

func (iter *RevIterator[T]) Partition(pred func(T) bool) ([]T, []T) {
	return Partition[T](iter, pred)
}

func (iter *RevIterator[T]) Filter(pred func(T) bool) *Filtered[T] {
	return Filter[T](iter, pred)
}

func (iter *RevIterator[T]) SkipWhile(pred func(T) bool) *SkipWhileT[T] {
	return SkipWhile[T](iter, pred)
}

func (iter *RevIterator[T]) TakeWhile(pred func(T) bool) *TakeWhileT[T] {
	return TakeWhile[T](iter, pred)
}

func (iter *RevIterator[T]) Chain(b Iterable[T]) *Chained[T] {
	return Chain[T](iter, b)
}

func (iter *RevIterator[T]) StepBy(step int) *Stepped[T] {
	return StepBy[T](iter, step)
}

func (iter *RevIterator[T]) Skip(n int) *Skipped[T] {
	return Skip[T](iter, n)
}

func (iter *RevIterator[T]) Take(n int) *Taken[T] {
	return Take[T](iter, n)
}

func (iter *RevIterator[T]) Collect() []T {
	return Collect[T](iter)
}

func (iter *RevIterator[T]) ForEach(fn func(T)) {
	ForEach[T](iter, fn)
}

func (iter *RevIterator[T]) Nth(n int) *T {
	return Nth[T](iter, n)
}

func (iter *RevIterator[T]) All(pred func(T) bool) bool {
	return All[T](iter, pred)
}

func (iter *RevIterator[T]) Any(pred func(T) bool) bool {
	return Any[T](iter, pred)
}

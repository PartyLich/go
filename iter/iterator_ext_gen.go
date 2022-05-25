// Code generated from template. May be overwritten if modified manually. DO NOT EDIT.

package iter

// Find searches for an element of an iterator that satisfies a predicate.
//
// Takes a function that returns true or false. It applies this function to
// each element of the iterator, and if any of them return true, then Find
// returns a pointer to the element. If they all return false, it returns
// nil.
//
// Find is short-circuiting; in other words, it will stop processing as soon as
// the predicate returns true.
func (iter *Iterator[T]) Find(pred func(T) bool) *T {
	return Find[T](iter, pred)
}

func (iter *Iterator[T]) Count() int {
	return Count[T](iter)
}

func (iter *Iterator[T]) Partition(pred func(T) bool) ([]T, []T) {
	return Partition[T](iter, pred)
}

func (iter *Iterator[T]) Filter(pred func(T) bool) *Filtered[T] {
	return Filter[T](iter, pred)
}

func (iter *Iterator[T]) SkipWhile(pred func(T) bool) *SkipWhileT[T] {
	return SkipWhile[T](iter, pred)
}

func (iter *Iterator[T]) TakeWhile(pred func(T) bool) *TakeWhileT[T] {
	return TakeWhile[T](iter, pred)
}

func (iter *Iterator[T]) Chain(b Iterable[T]) *Chained[T] {
	return Chain[T](iter, b)
}

func (iter *Iterator[T]) StepBy(step int) *Stepped[T] {
	return StepBy[T](iter, step)
}

func (iter *Iterator[T]) Skip(n int) *Skipped[T] {
	return Skip[T](iter, n)
}

func (iter *Iterator[T]) Take(n int) *Taken[T] {
	return Take[T](iter, n)
}

func (iter *Iterator[T]) Collect() []T {
	return Collect[T](iter)
}

func (iter *Iterator[T]) ForEach(fn func(T)) {
	ForEach[T](iter, fn)
}

func (iter *Iterator[T]) Nth(n int) *T {
	return Nth[T](iter, n)
}

func (iter *Iterator[T]) All(pred func(T) bool) bool {
	return All[T](iter, pred)
}

func (iter *Iterator[T]) Any(pred func(T) bool) bool {
	return Any[T](iter, pred)
}

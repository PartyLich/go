// Code generated from template. May be overwritten if modified manually. DO NOT EDIT.

package iter

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
// Code generated from template. May be overwritten if modified manually. DO NOT EDIT.

package iter

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

func (iter *Skipped[T]) Skip(n int) *Skipped[T] {
	return Skip[T](iter, n)
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

func (iter *Skipped[T]) Nth(n int) *T {
	return Nth[T](iter, n)
}

func (iter *Skipped[T]) All(pred func(T) bool) bool {
	return All[T](iter, pred)
}

func (iter *Skipped[T]) Any(pred func(T) bool) bool {
	return Any[T](iter, pred)
}

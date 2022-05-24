// Code generated from template. May be overwritten if modified manually. DO NOT EDIT.

package iter

func (iter *SkipWhileT[T]) Count() int {
	return Count[T](iter)
}

func (iter *SkipWhileT[T]) Partition(pred func(T) bool) ([]T, []T) {
	return Partition[T](iter, pred)
}

func (iter *SkipWhileT[T]) Filter(pred func(T) bool) *Filtered[T] {
	return Filter[T](iter, pred)
}

func (iter *SkipWhileT[T]) SkipWhile(pred func(T) bool) *SkipWhileT[T] {
	return SkipWhile[T](iter, pred)
}

func (iter *SkipWhileT[T]) TakeWhile(pred func(T) bool) *TakeWhileT[T] {
	return TakeWhile[T](iter, pred)
}

func (iter *SkipWhileT[T]) Chain(b Iterable[T]) *Chained[T] {
	return Chain[T](iter, b)
}

func (iter *SkipWhileT[T]) StepBy(step int) *Stepped[T] {
	return StepBy[T](iter, step)
}

func (iter *SkipWhileT[T]) Skip(n int) *Skipped[T] {
	return Skip[T](iter, n)
}

func (iter *SkipWhileT[T]) Take(n int) *Taken[T] {
	return Take[T](iter, n)
}

func (iter *SkipWhileT[T]) Collect() []T {
	return Collect[T](iter)
}

func (iter *SkipWhileT[T]) ForEach(fn func(T)) {
	ForEach[T](iter, fn)
}

func (iter *SkipWhileT[T]) Nth(n int) *T {
	return Nth[T](iter, n)
}

func (iter *SkipWhileT[T]) All(pred func(T) bool) bool {
	return All[T](iter, pred)
}

func (iter *SkipWhileT[T]) Any(pred func(T) bool) bool {
	return Any[T](iter, pred)
}

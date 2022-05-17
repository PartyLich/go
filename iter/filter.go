package iter

func (iter *Filtered[T]) Count() int {
	return Count[T](iter)
}

func (iter *Filtered[T]) Partition(pred func(T) bool) ([]T, []T) {
	return Partition[T](iter, pred)
}

func (iter *Filtered[T]) Filter(pred func(T) bool) *Filtered[T] {
	return Filter[T](iter, pred)
}

func (iter *Filtered[T]) SkipWhile(pred func(T) bool) *SkipWhileT[T] {
	return SkipWhile[T](iter, pred)
}

func (iter *Filtered[T]) TakeWhile(pred func(T) bool) *TakeWhileT[T] {
	return TakeWhile[T](iter, pred)
}

func (iter *Filtered[T]) Chain(b Iterable[T]) *Chained[T] {
	return Chain[T](iter, b)
}

func (iter *Filtered[T]) StepBy(step int) *Stepped[T] {
	return StepBy[T](iter, step)
}

func (iter *Filtered[T]) Skip(n int) *Skipped[T] {
	return Skip[T](iter, n)
}

func (iter *Filtered[T]) Take(n int) *Taken[T] {
	return Take[T](iter, n)
}
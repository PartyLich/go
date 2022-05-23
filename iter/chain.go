package iter

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

package iter

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

package iter

func (iter *Mapped[T, O]) Count() int {
	return Count[O](iter)
}

func (iter *Mapped[T, O]) Partition(pred func(O) bool) ([]O, []O) {
	return Partition[O](iter, pred)
}

func (iter *Mapped[T, O]) Filter(pred func(O) bool) *Filtered[O] {
	return Filter[O](iter, pred)
}

func (iter *Mapped[T, O]) SkipWhile(pred func(O) bool) *SkipWhileT[O] {
	return SkipWhile[O](iter, pred)
}

func (iter *Mapped[T, O]) TakeWhile(pred func(O) bool) *TakeWhileT[O] {
	return TakeWhile[O](iter, pred)
}

func (iter *Mapped[T, O]) Chain(b Iterable[O]) *Chained[O] {
	return Chain[O](iter, b)
}

func (iter *Mapped[T, O]) StepBy(step int) *Stepped[O] {
	return StepBy[O](iter, step)
}

func (iter *Mapped[T, O]) Skip(n int) *Skipped[O] {
	return Skip[O](iter, n)
}

func (iter *Mapped[T, O]) Take(n int) *Taken[O] {
	return Take[O](iter, n)
}

func (iter *Mapped[T, O]) Collect() []O {
	return Collect[O](iter)
}

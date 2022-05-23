package iter

type Mapped[T any, O any] struct {
	iter Iterable[T]
	fn   func(T) O
}

// Map returns an iterator that applies a function to every element.
func Map[T any, O any](iter Iterable[T], fn func(T) O) *Mapped[T, O] {
	return &Mapped[T, O]{iter, fn}
}

func (m *Mapped[T, O]) Next() *O {
	next := m.iter.Next()

	if next == nil {
		return nil
	}
	result := m.fn(*next)

	return &result
}

func (iter *Mapped[T, O]) Find(pred func(O) bool) *O {
	for next := iter.Next(); next != nil; next = iter.Next() {
		if pred(*next) {
			return next
		}
	}

	return nil
}

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

func (iter *Mapped[T, O]) ForEach(fn func(O)) {
	ForEach[O](iter, fn)
}

func (iter *Mapped[T, O]) Nth(n int) *O {
	return Nth[O](iter, n)
}

func (iter *Mapped[T, O]) All(pred func(O) bool) bool {
	return All[O](iter, pred)
}

func (iter *Mapped[T, O]) Any(pred func(O) bool) bool {
	return Any[O](iter, pred)
}

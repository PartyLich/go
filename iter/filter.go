package iter

type Filtered[T any] struct {
	iter Iterable[T]
	pred func(T) bool
}

// Filter returns an iterator which uses a predicate function to determine if an
// element should be yielded.
//
// The returned iterator will yield only the elements for which the predicate
// returns true.
func Filter[T any](iter Iterable[T], pred func(T) bool) *Filtered[T] {
	return &Filtered[T]{iter, pred}
}

func (f *Filtered[T]) Next() *T {
	next := f.iter.Find(f.pred)

	if next == nil || !f.pred(*next) {
		return nil
	}

	return next
}

func (iter *Filtered[T]) Find(pred func(T) bool) *T {
	for next := iter.Next(); next != nil; next = iter.Next() {
		if pred(*next) {
			return next
		}
	}

	return nil
}

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

func (iter *Filtered[T]) Collect() []T {
	return Collect[T](iter)
}

func (iter *Filtered[T]) ForEach(fn func(T)) {
	ForEach[T](iter, fn)
}

func (iter *Filtered[T]) Nth(n int) *T {
	return Nth[T](iter, n)
}

func (iter *Filtered[T]) All(pred func(T) bool) bool {
	return All[T](iter, pred)
}

func (iter *Filtered[T]) Any(pred func(T) bool) bool {
	return Any[T](iter, pred)
}

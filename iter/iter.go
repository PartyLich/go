package iter

type Iterable[T any] interface {
	Next() *T
	Find(pred func(T) bool) *T
}

type Iterator[T any] struct {
	idx   int
	slice []T
}

func New[T any](slice []T) *Iterator[T] {
	return &Iterator[T]{0, slice}
}

// Next advances the iterator and returns the next value.
//
// Returns nil when iteration is finished.
func (iter *Iterator[T]) Next() *T {
	if iter.idx >= len(iter.slice) {
		return nil
	}

	next := &iter.slice[iter.idx]
	iter.idx += 1

	return next
}

// Find searches for an element of an iterator that satisfies a predicate.
//
// takes a function that returns `true` or `false`. It applies this function to
// each element of the iterator, and if any of them return `true`, then Find
// returns a pointer to the element. If they all return `false`, it returns
// `nil`.
//
// Find is short-circuiting; in other words, it will stop processing as soon as
// the closure returns `true`.
func (iter *Iterator[T]) Find(pred func(T) bool) *T {
	for iter.idx < len(iter.slice) {
		next := iter.slice[iter.idx]
		iter.idx++

		if pred(next) {
			return &next
		}
	}

	return nil
}

type RevIterator[T any] struct {
	Iterator[T]
}

// Rev reverses the iteration order of this iterator
func (iter *Iterator[T]) Rev() *RevIterator[T] {
	var idx int

	if iter.idx == len(iter.slice)-1 {
		idx = 0
	}
	if iter.idx == 0 {
		idx = len(iter.slice) - 1
	}

	return &RevIterator[T]{
		Iterator[T]{idx, iter.slice},
	}
}

func (iter *RevIterator[T]) Next() *T {
	if iter.idx < 0 {
		return nil
	}

	next := &iter.slice[iter.idx]
	iter.idx -= 1

	return next
}

// Find searches for an element of an iterator that satisfies a predicate.
//
// takes a function that returns `true` or `false`. It applies this function to
// each element of the iterator, and if any of them return `true`, then Find
// returns a pointer to the element. If they all return `false`, it returns
// `nil`.
//
// Find is short-circuiting; in other words, it will stop processing as soon as
// the closure returns `true`.
func (iter *RevIterator[T]) Find(pred func(T) bool) *T {
	for iter.idx >= 0 {
		next := iter.slice[iter.idx]
		iter.idx--

		if pred(next) {
			return &next
		}
	}

	return nil
}

type Mapped[T any, O any] struct {
	iter Iterable[T]
	fn   func(T) O
}

// Map returns an iterator that applies a function to every element.
func Map[T any, O any](iter Iterable[T], fn func(T) O) Mapped[T, O] {
	return Mapped[T, O]{iter, fn}
}

func (m *Mapped[T, O]) Next() *O {
	next := m.iter.Next()

	if next == nil {
		return nil
	}
	result := m.fn(*next)

	return &result
}

type Filtered[T any] struct {
	iter Iterable[T]
	pred func(T) bool
}

// Filter returns an iterator which uses a function to determine if an
// element should be yielded.
//
// The returned iterator will yield only the elements for which the closure
// returns true.
func Filter[T any](iter Iterable[T], pred func(T) bool) Filtered[T] {
	return Filtered[T]{iter, pred}
}

func (f *Filtered[T]) Next() *T {
	next := f.iter.Find(f.pred)

	if next == nil || !f.pred(*next) {
		return nil
	}

	return next
}

// Reduce repeatedly applies a reducing operation, reducing the iterator to a
// single element
func Reduce[T any, O any](iter Iterable[T], init O, fn func(O, T) O) O {
	accum := init

	for val := iter.Next(); val != nil; val = iter.Next() {
		accum = fn(accum, *val)
	}

	return accum
}

// Fold repeatedly applies a reducing operation, reducing the iterator to a
// single element
//
// Alias for Reduce
func Fold[T any, O any](iter Iterable[T], init O, fn func(O, T) O) O {
	return Reduce(iter, init, fn)
}

// SkipWhile iterable adapter
type SkipWhileT[T any] struct {
	iter Iterable[T]
	flag bool
	pred func(T) bool
}

// SkipWhile creates an iterator that skips elements based on a predicate.
func SkipWhile[T any](iter Iterable[T], pred func(T) bool) SkipWhileT[T] {
	return SkipWhileT[T]{iter, false, pred}
}

func (s *SkipWhileT[T]) Next() *T {
	check := func(flag *bool, pred func(T) bool) func(T) bool {
		return func(t T) bool {
			if *flag || !pred(t) {
				*flag = true
				return true
			} else {
				return false
			}
		}
	}

	return s.iter.Find(check(&s.flag, s.pred))
}

func (iter *SkipWhileT[T]) Find(pred func(T) bool) *T {
	for next := iter.Next(); next != nil; next = iter.Next() {
		if pred(*next) {
			return next
		}
	}

	return nil
}

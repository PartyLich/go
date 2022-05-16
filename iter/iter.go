package iter

type Iterable[T any] interface {
	Next() *T
	Find(pred func(T) bool) *T
}

type Iterator[T any] struct {
	idx   int
	slice []T
}

// New creates a new lazy iterator over the provided slice
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

type Filtered[T any] struct {
	iter Iterable[T]
	pred func(T) bool
}

// Filter returns an iterator which uses a function to determine if an
// element should be yielded.
//
// The returned iterator will yield only the elements for which the closure
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

// Partition consumes an iterator, creating two slices from it.
//
// The first slice contains all of the elements for which the predicate returned
// true, and the second slice contains all of the elements for which it returned
// false.
func Partition[T any](iter Iterable[T], pred func(T) bool) ([]T, []T) {
	var a, b []T

	for next := iter.Next(); next != nil; next = iter.Next() {
		if pred(*next) {
			a = append(a, *next)
		} else {
			b = append(b, *next)
		}
	}

	return a, b
}

type Chainer[T any] interface {
	Chain(Iterable[T]) Iterable[T]
}

type Chained[T any] struct {
	a, b Iterable[T]
}

// Chain takes two iterators and creates a new iterator over both in sequence.
//
// Chain will return a new iterator which will first iterate over values from
// the first iterator and then over values from the second iterator.
func Chain[T any](a, b Iterable[T]) *Chained[T] {
	return &Chained[T]{a, b}
}

func (c *Chained[T]) Next() *T {
	next := c.a.Next()

	if next == nil {
		return c.b.Next()
	}

	return next
}

func (c *Chained[T]) Find(pred func(T) bool) *T {
	for next := c.a.Next(); next != nil; next = c.a.Next() {
		if pred(*next) {
			return next
		}
	}
	for next := c.b.Next(); next != nil; next = c.b.Next() {
		if pred(*next) {
			return next
		}
	}

	return nil
}

type Stepped[T any] struct {
	iter  Iterable[T]
	step  int
	first bool
}

// StepBy creates an iterator starting at the same point, but stepping by the
// given amount at each iteration.
//
// The method will panic if the given step is <= 0.
//
// Note 1: The first element of the iterator will always be returned, regardless
// of the step given.
func StepBy[T any](a Iterable[T], step int) *Stepped[T] {
	if step <= 0 {
		panic("StepBy requires a step value > 0")
	}

	return &Stepped[T]{a, step, true}
}

func (s *Stepped[T]) Next() *T {
	if s.first {
		s.first = false
		return s.iter.Next()
	}

	next := s.iter.Next()
	for i := 1; i < s.step; i++ {
		next = s.iter.Next()
	}

	return next
}

func (s *Stepped[T]) Find(pred func(T) bool) *T {
	var next *T

	for next = s.Next(); next != nil; next = s.Next() {
		if pred(*next) {
			break
		}
	}

	return next
}

// TakeWhile iterable adapter
type TakeWhileT[T any] struct {
	iter Iterable[T]
	flag bool
	pred func(T) bool
}

// TakeWhile Creates an iterator that yields elements based on a predicate.
//
// TakeWhile takes a predicate function as an argument. It will call this
// function on each element of the iterator, and yield elements while it returns
// true.
func TakeWhile[T any](iter Iterable[T], pred func(T) bool) *TakeWhileT[T] {
	return &TakeWhileT[T]{iter, true, pred}
}

func (s *TakeWhileT[T]) Next() *T {
	check := func(flag *bool, pred func(T) bool) func(T) bool {
		return func(t T) bool {
			if *flag && pred(t) {
				return true
			} else {
				*flag = false
				return false
			}
		}
	}

	return s.iter.Find(check(&s.flag, s.pred))
}

func (iter *TakeWhileT[T]) Find(pred func(T) bool) *T {
	for next := iter.Next(); next != nil; next = iter.Next() {
		if pred(*next) {
			return next
		}
	}

	return nil
}

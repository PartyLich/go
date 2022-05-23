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

// Collect transforms an iterator into a slice.
func Collect[T any](iter Iterable[T]) []T {
	var out []T

	for next := iter.Next(); next != nil; next = iter.Next() {
		out = append(out, *next)
	}

	return out
}

// Count consumes the iterator, counting the number of iterations and returning
// it.
func Count[T any](iter Iterable[T]) int {
	count := 0

	for next := iter.Next(); next != nil; next = iter.Next() {
		count += 1
	}

	return count
}

// Calls a function on each element of an iterator.
//
// This is equivalent to using a for loop on the iterator, although break and
// continue are not possible.
func ForEach[T any](iter Iterable[T], fn func(T)) {
	for val := iter.Next(); val != nil; val = iter.Next() {
		fn(*val)
	}
}

// Nth returns the `n`th element of the iterator.
//
// Like most indexing operations, the count starts from zero, so `Nth(0)`
// returns the first value, `nth(1)` the second, and so on.
//
// Note that all preceding elements, as well as the returned element, will be
// consumed from the iterator. That means that the preceding elements will be
// discarded, and also that calling `Nth(0)` multiple times on the same iterator
// will return different elements.
//
// Nth will return `nil` if `n` is greater than or equal to the length of the
// iterator.
func Nth[T any](iter Iterable[T], n int) *T {
	if n < 0 {
		panic("Nth expected n to be >= 0")
	}

	var next *T

	for ; n >= 0; n-- {
		next = iter.Next()
		if next == nil {
			break
		}
	}

	return next
}

// SkipWhile iterable adapter
type SkipWhileT[T any] struct {
	iter Iterable[T]
	flag bool
	pred func(T) bool
}

// SkipWhile creates an iterator that skips elements based on a predicate.
func SkipWhile[T any](iter Iterable[T], pred func(T) bool) *SkipWhileT[T] {
	return &SkipWhileT[T]{iter, false, pred}
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

// All tests if every element of the iterator matches a predicate.
//
// All takes a function that returns true or false. It applies this function to
// each element of the iterator, and if they all return true, then so does All.
// If any of them return false, it returns false.
//
// All is short-circuiting; in other words, it will stop processing as soon as
// it finds a false, given that no matter what else happens, the result will
// also be false.
//
// An empty iterator returns true.
func All[T any](iter Iterable[T], pred func(T) bool) bool {
	for next := iter.Next(); next != nil; next = iter.Next() {
		if !pred(*next) {
			return false
		}
	}

	return true
}

// Any tests if any element of the iterator matches a predicate.
//
// Any takes a function that returns true or false. It applies this function to
// each element of the iterator, and if any of them return true, then so does
// Any. If they all return false, it returns false.
//
// Any is short-circuiting; in other words, it will stop processing as soon as
// it finds a true, given that no matter what else happens, the result will also
// be true.
//
// An empty iterator returns false.
func Any[T any](iter Iterable[T], pred func(T) bool) bool {
	for next := iter.Next(); next != nil; next = iter.Next() {
		if pred(*next) {
			return true
		}
	}

	return false
}

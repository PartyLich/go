// Code generated from template. May be overwritten if modified manually. DO NOT EDIT.

package iter

// Find searches for an element of an iterator that satisfies a predicate.
//
// Takes a function that returns true or false. It applies this function to
// each element of the iterator, and if any of them return true, then Find
// returns a pointer to the element. If they all return false, it returns
// nil.
//
// Find is short-circuiting; in other words, it will stop processing as soon as
// the predicate returns true.
func (iter *Skipped[T]) Find(pred func(T) bool) *T {
	return Find[T](iter, pred)
}

// Count consumes the iterator, counting the number of iterations and returning
// it.
func (iter *Skipped[T]) Count() int {
	return Count[T](iter)
}

// Partition consumes an iterator, creating two slices from it.
//
// The first slice contains all of the elements for which the predicate returned
// true, and the second slice contains all of the elements for which it returned
// false.
func (iter *Skipped[T]) Partition(pred func(T) bool) ([]T, []T) {
	return Partition[T](iter, pred)
}

// Filter returns an iterator which uses a predicate function to determine if an
// element should be yielded.
//
// The returned iterator will yield only the elements for which the predicate
// returns true.
func (iter *Skipped[T]) Filter(pred func(T) bool) *Filtered[T] {
	return Filter[T](iter, pred)
}

// SkipWhile creates an iterator that skips elements based on a predicate.
func (iter *Skipped[T]) SkipWhile(pred func(T) bool) *SkipWhileT[T] {
	return SkipWhile[T](iter, pred)
}

// TakeWhile Creates an iterator that yields elements based on a predicate.
func (iter *Skipped[T]) TakeWhile(pred func(T) bool) *TakeWhileT[T] {
	return TakeWhile[T](iter, pred)
}

// Chain takes two iterators and creates a new iterator over both in sequence.
//
// Chain will return a new iterator which will first iterate over values from
// the first iterator and then over values from the second iterator.
func (iter *Skipped[T]) Chain(b Iterable[T]) *Chained[T] {
	return Chain[T](iter, b)
}

// StepBy creates an iterator starting at the same point, but stepping by the
// given amount at each iteration.
//
// The method will panic if the given step is <= 0.
//
// Note 1: The first element of the iterator will always be returned, regardless
// of the step given.
func (iter *Skipped[T]) StepBy(step int) *Stepped[T] {
	return StepBy[T](iter, step)
}

// Skip creates an iterator that skips the first n elements.
func (iter *Skipped[T]) Skip(n int) *Skipped[T] {
	return Skip[T](iter, n)
}

// Take creates an iterator that yields the first n elements, or fewer if the
// underlying iterator ends sooner.
func (iter *Skipped[T]) Take(n int) *Taken[T] {
	return Take[T](iter, n)
}

// Collect transforms an iterator into a slice.
func (iter *Skipped[T]) Collect() []T {
	return Collect[T](iter)
}

// ForEach calls a function on each element of an iterator.
//
// This is equivalent to using a for loop on the iterator, although break and
// continue are not possible.
func (iter *Skipped[T]) ForEach(fn func(T)) {
	ForEach[T](iter, fn)
}

// Nth returns the nth element of the iterator.
//
// Like most indexing operations, the count starts from zero, so Nth(0)
// returns the first value, nth(1) the second, and so on.
//
// Note that all preceding elements, as well as the returned element, will be
// consumed from the iterator. That means that the preceding elements will be
// discarded, and also that calling Nth(0) multiple times on the same iterator
// will return different elements.
//
// Nth will return nil if n is greater than or equal to the length of the
// iterator.
func (iter *Skipped[T]) Nth(n int) *T {
	return Nth[T](iter, n)
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
func (iter *Skipped[T]) All(pred func(T) bool) bool {
	return All[T](iter, pred)
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
func (iter *Skipped[T]) Any(pred func(T) bool) bool {
	return Any[T](iter, pred)
}

// Last consumes the iterator, returning the last element.
//
// This method will evaluate the iterator until it returns nil. While doing so,
// it keeps track of the current element. After nil is returned, Last will then
// return the last element it saw.
func (iter *Skipped[T]) Last() *T {
	return Last[T](iter)
}

package iter

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

//go:generate go run ./cmd/gen/ -name Stepped -output stepBy_ext_gen.go

package iter

import "testing"

func assertEq[T comparable](t *testing.T, a, b T) {
	if a != b {
		t.Errorf("expected %v == %v", a, b)
	}
}

func TestNext(t *testing.T) {
	cases := []struct {
		list []int
	}{
		{[]int{1, 2, 3}},
		{[]int{}},
	}

	for _, c := range cases {
		i := New(c.list)
		idx := 0
		for have := i.Next(); have != nil; have = i.Next() {
			if *have != c.list[idx] {
				t.Errorf("Next \n\thave %v\n\twant %v", *have, c.list[idx])
			}

			idx += 1
		}
	}
}

func TestRevNext(t *testing.T) {
	cases := []struct {
		list []int
	}{
		{[]int{1, 2, 3}},
		{[]int{}},
	}

	for _, c := range cases {
		i := New(c.list).Rev()

		idx := len(c.list) - 1
		for have := i.Next(); have != nil; have = i.Next() {
			assertEq(t, *have, c.list[idx])
			idx -= 1
		}
	}
}

func TestFind(t *testing.T) {
	list := []int{1, 2, 3, 4}
	pred := func(i int) bool { return i == 2 }

	i := New(list)
	assertEq(t, *i.Find(pred), 2)
	assertEq(t, i.idx, 2)
	assertEq(t, i.Find(pred), nil)
}

func TestRevFind(t *testing.T) {
	list := []int{1, 2, 3, 4}
	pred := func(i int) bool { return i == 2 }

	i := New(list).Rev()
	assertEq(t, *i.Find(pred), 2)
	assertEq(t, i.idx, 0)
	assertEq(t, *i.Next(), 1)
	assertEq(t, i.Find(pred), nil)
}

func TestFilter(t *testing.T) {
	list := []int{1, 2, 3, 4}
	filter := func(i int) bool { return i%2 == 0 }

	iter := New(list)
	i := Filter[int](iter, filter)
	assertEq(t, *i.Next(), 2)
	assertEq(t, *i.Next(), 4)
	assertEq(t, i.Next(), nil)
}

func TestFilterFind(t *testing.T) {
	list := []int{-4, -2, 1, 2}
	isNeg := func(a int) bool { return a < 0 }
	pred := func(i int) bool { return i%2 == 0 }

	iter := New(list)
	i := Filter[int](iter, isNeg)

	assertEq(t, *i.Find(pred), -4)
	assertEq(t, *i.Find(pred), -2)
	assertEq(t, i.Find(pred), nil)
}

func TestMap(t *testing.T) {
	list := []int{1, 2, 3, 4}
	want := []int{2, 4, 6, 8}
	fn := func(i int) int { return i * 2 }

	iter := New(list)
	i := Map[int, int](iter, fn)

	for have, idx := i.Next(), 0; have != nil; have, idx = i.Next(), idx+1 {
		assertEq(t, *have, want[idx])
	}
}

func TestMapFind(t *testing.T) {
	list := []int{1, 2, 3, 4}
	pred := func(i int) bool { return i%2 == 0 }
	fn := func(i int) int { return i + 1 }

	iter := New(list)
	i := Map[int, int](iter, fn)

	assertEq(t, *i.Find(pred), 2)
	assertEq(t, *i.Find(pred), 4)
	assertEq(t, i.Find(pred), nil)
}

func TestReduce(t *testing.T) {
	list := []int{1, 2, 3, 4}
	sum := func(a, b int) int { return a + b }

	iter := New(list)
	have := Reduce[int, int](iter, 0, sum)
	assertEq(t, have, 10)
	assertEq(t, iter.Next(), nil)

	// Fold alias
	iter = New(list)
	have = Fold[int, int](iter, 0, sum)
	assertEq(t, have, 10)
	assertEq(t, iter.Next(), nil)
}

func TestRevReduce(t *testing.T) {
	list := []int{1, 2, 3, 4}
	sub := func(a, b int) int { return a - b }

	r := New(list).Rev()
	have := Reduce[int, int](r, 0, sub)
	assertEq(t, have, -10)
	assertEq(t, r.Next(), nil)

	// Fold alias
	r = New(list).Rev()
	have = Reduce[int, int](r, 0, sub)
	assertEq(t, have, -10)
	assertEq(t, r.Next(), nil)
}

func TestSkipWhile(t *testing.T) {
	list := []int{-1, 2, 3, 4}
	isNeg := func(a int) bool { return a < 0 }

	iter := New(list)
	i := SkipWhile[int](iter, isNeg)
	assertEq(t, *i.Next(), 2)
	assertEq(t, *i.Next(), 3)
	assertEq(t, *i.Next(), 4)
	assertEq(t, i.Next(), nil)
}

func TestSkipWhileFind(t *testing.T) {
	list := []int{-4, -2, 1, 2}
	isNeg := func(a int) bool { return a < 0 }
	pred := func(i int) bool { return i%2 == 0 }

	iter := New(list)
	i := SkipWhile[int](iter, isNeg)

	assertEq(t, *i.Find(pred), 2)
	assertEq(t, i.Find(pred), nil)
}

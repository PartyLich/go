package iter

import (
	"container/list"
	"testing"
)

func assertEq[T comparable](t *testing.T, a, b T) {
	if a != b {
		t.Errorf("expected %v == %v", a, b)
	}
}

func assertPanic(t *testing.T, f func()) {
	defer func() {
		if recover() == nil {
			t.Errorf("The code did not panic")
		}
	}()
	f()
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
	assertEq(t, i.it.idx, 0)
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

func TestCollect(t *testing.T) {
	list := []int{1, 2, 3, 4}

	iter := New(list)
	have := Collect[int](iter)

	for idx, want := range list {
		assertEq(t, have[idx], want)
	}
	// iterator consumed
	assertEq(t, iter.Next(), nil)

	// original list unchanged
	have[0] = -4
	assertEq(t, 1, list[0])
	assertEq(t, -4, have[0])
}

func TestForEach(t *testing.T) {
	list := []int{1, 2, 3, 4}
	have := make([]int, 0, len(list))
	fn := func(i int) {
		have = append(have, i)
	}

	iter := New(list)
	ForEach[int](iter, fn)

	for idx, want := range list {
		assertEq(t, have[idx], want)
	}
	// iterator consumed
	assertEq(t, iter.Next(), nil)
}

func TestNth(t *testing.T) {
	list := []int{1, 2, 3, 4}
	have := *Nth[int](New(list), 1)

	assertEq(t, have, 2)
	assertPanic(t, func() { Nth[int](New(list), -1) })
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

func TestPartition(t *testing.T) {
	list := []int{1, 2, 3, 4}
	isEven := func(a int) bool { return a%2 == 0 }

	iter := New(list)
	a, b := Partition[int](iter, isEven)
	wantA := []int{2, 4}
	wantB := []int{1, 3}

	for i, v := range a {
		assertEq(t, v, wantA[i])
	}
	for i, v := range b {
		assertEq(t, v, wantB[i])
	}
}

func TestAdapterIsIterable(t *testing.T) {
	list := []int{1, 2, 3, 4}
	ident := func(i int) int { return i }
	all := func(i int) bool { return true }

	var it Iterable[int]
	it = New(list)
	it = New(list).Rev()
	it = Filter[int](New(list), all)
	it = Map[int, int](New(list), ident)
	it = Chain[int](New(list), New(list))
	it = StepBy[int](New(list), 2)
	it = SkipWhile[int](New(list), all)
	it = TakeWhile[int](New(list), all)
	it = Skip[int](New(list), 2)
	it = Take[int](New(list), 2)
	_ = it
}

func TestChain(t *testing.T) {
	a1 := []int{1, 2, 3}
	a2 := []int{4, 5, 6}

	iter := Chain[int](New(a1), New(a2))
	want := []int{1, 2, 3, 4, 5, 6}

	for _, v := range want {
		assertEq(t, v, *iter.Next())
	}
}

func TestChain_Find(t *testing.T) {
	a1 := []int{1, 2, 3}
	a2 := []int{4, 5, 6}
	pred := func(i int) bool { return i > 4 }

	i := Chain[int](New(a1), New(a2))

	assertEq(t, *i.Find(pred), 5)
	assertEq(t, *i.Find(pred), 6)
	assertEq(t, i.Find(pred), nil)
}

func TestStepBy(t *testing.T) {
	list := []int{1, 2, 3, 4, 5, 6}

	i := StepBy[int](New(list), 2)
	want := []int{1, 3, 5}

	for _, v := range want {
		assertEq(t, v, *i.Next())
	}
	assertEq(t, i.Next(), nil)

	// should panic on invalid step
	assertPanic(t, func() { StepBy[int](New(list), 0) })
	assertPanic(t, func() { StepBy[int](New(list), -1) })
}

func TestStepBy_Find(t *testing.T) {
	list := []int{1, 2, 3, 4, 5, 6}
	pred := func(i int) bool { return i == 2 || i == 5 }

	i := StepBy[int](New(list), 2)

	assertEq(t, *i.Find(pred), 5)
	assertEq(t, i.Find(pred), nil)
}

func TestTakeWhile(t *testing.T) {
	list := []int{-1, -2, 3, 4}
	isNeg := func(a int) bool { return a < 0 }

	iter := New(list)
	i := TakeWhile[int](iter, isNeg)
	assertEq(t, *i.Next(), -1)
	assertEq(t, *i.Next(), -2)
	assertEq(t, i.Next(), nil)
}

func TestTakeWhile_Find(t *testing.T) {
	list := []int{-4, -2, 1, 2, 4}
	isNeg := func(a int) bool { return a < 0 }
	pred := func(i int) bool { return i%2 == 0 }

	iter := New(list)
	i := TakeWhile[int](iter, isNeg)

	assertEq(t, *i.Find(pred), -4)
	assertEq(t, *i.Find(pred), -2)
	assertEq(t, i.Find(pred), nil)
}

func TestCount(t *testing.T) {
	list := []int{1, 2, 3, 4, 5}
	i := New(list)

	assertEq(t, Count[int](i), 5)
}

func TestSkip(t *testing.T) {
	list := []int{-1, -2, 3, 4}

	iter := New(list)
	i := Skip[int](iter, 2)
	assertEq(t, *i.Next(), 3)
	assertEq(t, *i.Next(), 4)
	assertEq(t, i.Next(), nil)

	assertPanic(t, func() { Skip[int](iter, -1) })
}

func TestSkip_Find(t *testing.T) {
	list := []int{-4, -2, 1, 2, 4}
	pred := func(i int) bool { return i%2 == 0 }

	iter := New(list)
	i := Skip[int](iter, 3)

	assertEq(t, *i.Find(pred), 2)
	assertEq(t, *i.Find(pred), 4)
	assertEq(t, i.Find(pred), nil)
}

func TestTake(t *testing.T) {
	list := []int{-1, -2, 3, 4}

	iter := New(list)
	i := Take[int](iter, 2)
	assertEq(t, *i.Next(), -1)
	assertEq(t, *i.Next(), -2)
	assertEq(t, i.Next(), nil)
}

func TestTake_Find(t *testing.T) {
	list := []int{-4, -2, 1, 2, 4}
	pred := func(i int) bool { return i%2 == 0 }

	iter := New(list)
	i := Take[int](iter, 3)

	assertEq(t, *i.Find(pred), -4)
	assertEq(t, *i.Find(pred), -2)
	assertEq(t, i.Find(pred), nil)
}

func TestAll(t *testing.T) {
	list := []int{-4, -2, 2, 4}
	isPos := func(i int) bool { return i >= 0 }
	isEven := func(a int) bool { return a%2 == 0 }

	assertEq(t, All[int](New(list), isPos), false)
	assertEq(t, All[int](New(list), isEven), true)

	assertEq(t, All[int](New([]int{}), isEven), true)
	assertEq(t, All[int](New([]int{}), isPos), true)
}

func TestAny(t *testing.T) {
	list := []int{-4, -2, 2, 4}
	isPos := func(i int) bool { return i >= 0 }
	isOdd := func(a int) bool { return a%2 != 0 }

	assertEq(t, Any[int](New(list), isPos), true)
	assertEq(t, Any[int](New(list), isOdd), false)

	assertEq(t, Any[int](New([]int{}), isOdd), false)
	assertEq(t, Any[int](New([]int{}), isPos), false)
}

func TestMin(t *testing.T) {
	list := []int{-4, -2, 2, 4}
	i := New(list)

	assertEq(t, *Min[int](i), -4)

	assertEq(t, Min[int](New([]int{})), nil)

	s := []string{"abc", "bcd"}
	assertEq(t, *Min[string](New(s)), "abc")
}

func TestMax(t *testing.T) {
	list := []int{-4, -2, 2, 4}
	i := New(list)

	assertEq(t, *Max[int](i), 4)

	assertEq(t, Max[int](New([]int{})), nil)

	s := []string{"abc", "bcd"}
	assertEq(t, *Max[string](New(s)), "bcd")
}

func TestLast(t *testing.T) {
	list := []int{-4, -2, 2, 4}
	i := New(list)

	assertEq(t, *Last[int](i), 4)
	assertEq(t, i.Next(), nil)

	empty := New([]int{})
	assertEq(t, Last[int](empty), nil)

	s := []string{"abc", "bcd"}
	assertEq(t, *Last[string](New(s)), "bcd")
}

func makeList(n int) *list.List {
	l := list.New()
	for i := 1; i <= n; i++ {
		l.PushBack(i)
	}

	return l
}

func TestListIterator_Next(t *testing.T) {
	cases := []struct {
		list *list.List
		want []int
	}{
		{makeList(3), []int{1, 2, 3}},
		{makeList(0), []int{}},
		{makeList(3), []int{1, 2, 3}},
	}
	cases[2].list.PushBack("different type")

	for _, c := range cases {
		i := FromList[int](c.list)
		for idx, have := 0, i.Next(); have != nil; idx, have = idx+1, i.Next() {
			if *have != c.want[idx] {
				t.Errorf("Next \n\thave %v\n\twant %v", *have, c.want[idx])
			}
		}
	}
}

func TestListIterator_Find(t *testing.T) {
	pred := func(i int) bool { return i == 2 }
	l := makeList(3)
	i := FromList[int](l)

	assertEq(t, *i.Find(pred), 2)
	assertEq(t, i.Find(pred), nil)
}

func TestFlat_Next(t *testing.T) {
	data := New([]Iterable[int]{
		New([]int{1, 2}),
		New([]int{3, 4}),
	})
	f := Flatten[int](data)

	for i := 1; i <= 4; i++ {
		assertEq(t, *f.Next(), i)
	}
	assertEq(t, f.Next(), nil)
}

func TestFlat_Find(t *testing.T) {
	pred := func(i int) bool { return i == 2 }
	data := New([]Iterable[int]{
		New([]int{1, 2}),
		New([]int{3, 4}),
	})
	f := Flatten[int](data)

	assertEq(t, *f.Find(pred), 2)
	assertEq(t, f.Find(pred), nil)
}

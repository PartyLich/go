package iter_test

import (
	"fmt"

	"github.com/partylich/go/iter"
)

func ExampleFlat_Find() {
	data := iter.New([]iter.Iterable[int]{
		iter.New([]int{1, 2}),
		iter.New([]int{3, 4}),
	})
	isTwo := func(i int) bool { return i == 2 }
	i := iter.Flatten[int](data)

	fmt.Println(*i.Find(isTwo))
	fmt.Println(i.Find(isTwo))
	// Output:
	// 2
	// <nil>
}

func ExampleFlat_Find_mapped() {
	toIter := func(a []int) iter.Iterable[int] { return iter.New(a) }
	data := [][]int{{1, 2}, {3, 4}}
	m := iter.Map[[]int](iter.New(data), toIter)
	i := iter.Flatten[int](m)

	isTwo := func(i int) bool { return i == 2 }

	fmt.Println(*i.Find(isTwo))
	fmt.Println(i.Find(isTwo))
	// Output:
	// 2
	// <nil>
}

func ExampleFlat_Count() {
	data := iter.New([]iter.Iterable[int]{
		iter.New([]int{1, 2}),
		iter.New([]int{3, 4}),
	})
	i := iter.Flatten[int](data)

	fmt.Println(i.Count())
	// Output:
	// 4
}

func ExampleFlat_Count_mapped() {
	toIter := func(a []int) iter.Iterable[int] { return iter.New(a) }
	data := [][]int{{1, 2}, {3, 4}}
	m := iter.Map[[]int](iter.New(data), toIter)
	i := iter.Flatten[int](m)

	fmt.Println(i.Count())
	// Output:
	// 4
}

func ExampleFlat_Filter() {
	data := iter.New([]iter.Iterable[int]{
		iter.New([]int{1, 2}),
		iter.New([]int{3, 4}),
	})
	i := iter.Flatten[int](data)

	f := i.Filter(func(n int) bool { return n > 2 })

	// alternatively,
	//  gt2 := func(n int) bool { return n > 2 }
	//  f := i.Filter(gt5)

	for val := f.Next(); val != nil; val = f.Next() {
		fmt.Println(*val)
	}
	// Output:
	// 3
	// 4
}

func ExampleFlat_Filter_mapped() {
	toIter := func(a []int) iter.Iterable[int] { return iter.New(a) }
	data := [][]int{{1, 2}, {3, 4}}
	m := iter.Map[[]int](iter.New(data), toIter)
	i := iter.Flatten[int](m)

	f := i.Filter(func(n int) bool { return n > 2 })

	// alternatively,
	//  gt2 := func(n int) bool { return n > 2 }
	//  f := i.Filter(gt5)

	for val := f.Next(); val != nil; val = f.Next() {
		fmt.Println(*val)
	}
	// Output:
	// 3
	// 4
}

func ExampleFlat_SkipWhile() {
	isNeg := func(a int) bool { return a < 0 }
	data := iter.New([]iter.Iterable[int]{
		iter.New([]int{-1, 2}),
		iter.New([]int{3, 4}),
	})

	i := iter.Flatten[int](data).SkipWhile(isNeg)

	fmt.Println(*i.Next())
	fmt.Println(*i.Next())
	fmt.Println(*i.Next())
	fmt.Println(i.Next())
	// Output:
	// 2
	// 3
	// 4
	// <nil>
}

func ExampleFlat_SkipWhile_mapped() {
	toIter := func(a []int) iter.Iterable[int] { return iter.New(a) }
	data := [][]int{{-1, 2}, {3, 4}}
	m := iter.Map[[]int](iter.New(data), toIter)

	isNeg := func(a int) bool { return a < 0 }
	i := iter.Flatten[int](m).SkipWhile(isNeg)

	fmt.Println(*i.Next())
	fmt.Println(*i.Next())
	fmt.Println(*i.Next())
	fmt.Println(i.Next())
	// Output:
	// 2
	// 3
	// 4
	// <nil>
}

func ExampleFlat_Partition() {
	isEven := func(a int) bool { return a%2 == 0 }
	data := iter.New([]iter.Iterable[int]{
		iter.New([]int{1, 2}),
		iter.New([]int{3, 4}),
	})
	i := iter.Flatten[int](data)

	a, b := i.Partition(isEven)

	fmt.Println(a)
	fmt.Println(b)
	// Output:
	// [2 4]
	// [1 3]
}

func ExampleFlat_Partition_mapped() {
	toIter := func(a []int) iter.Iterable[int] { return iter.New(a) }
	data := [][]int{{1, 2}, {3, 4}}
	m := iter.Map[[]int](iter.New(data), toIter)
	i := iter.Flatten[int](m)

	isEven := func(a int) bool { return a%2 == 0 }
	a, b := i.Partition(isEven)

	fmt.Println(a)
	fmt.Println(b)
	// Output:
	// [2 4]
	// [1 3]
}

func ExampleFlat_Chain() {
	l1 := iter.New([]iter.Iterable[int]{
		iter.New([]int{1, 2}),
		iter.New([]int{3}),
	})
	l2 := iter.New([]iter.Iterable[int]{
		iter.New([]int{4, 5}),
		iter.New([]int{6}),
	})
	a1 := iter.Flatten[int](l1)
	a2 := iter.Flatten[int](l2)

	i := a1.Chain(a2)

	for val := i.Next(); val != nil; val = i.Next() {
		fmt.Println(*val)
	}
	// Output:
	// 1
	// 2
	// 3
	// 4
	// 5
	// 6
}

func ExampleFlat_Chain_mapped() {
	toIter := func(a []int) iter.Iterable[int] { return iter.New(a) }
	l1 := [][]int{{1, 2}, {3}}
	m1 := iter.Map[[]int](iter.New(l1), toIter)
	a1 := iter.Flatten[int](m1)

	l2 := [][]int{{4, 5}, {6}}
	m2 := iter.Map[[]int](iter.New(l2), toIter)
	a2 := iter.Flatten[int](m2)

	i := a1.Chain(a2)

	for val := i.Next(); val != nil; val = i.Next() {
		fmt.Println(*val)
	}
	// Output:
	// 1
	// 2
	// 3
	// 4
	// 5
	// 6
}

func ExampleFlat_StepBy() {
	data := iter.New([]iter.Iterable[int]{
		iter.New([]int{1, 2}),
		iter.New([]int{3, 4, 5, 6}),
	})
	i := iter.Flatten[int](data).StepBy(2)

	for val := i.Next(); val != nil; val = i.Next() {
		fmt.Println(*val)
	}
	// Output:
	// 1
	// 3
	// 5
}

func ExampleFlat_StepBy_mapped() {
	toIter := func(a []int) iter.Iterable[int] { return iter.New(a) }
	data := [][]int{{1, 2}, {3, 4, 5, 6}}
	m := iter.Map[[]int](iter.New(data), toIter)
	i := iter.Flatten[int](m).StepBy(2)

	for val := i.Next(); val != nil; val = i.Next() {
		fmt.Println(*val)
	}
	// Output:
	// 1
	// 3
	// 5
}

func ExampleFlat_TakeWhile() {
	isNeg := func(a int) bool { return a < 0 }

	data := iter.New([]iter.Iterable[int]{
		iter.New([]int{-1, -2}),
		iter.New([]int{3, 4}),
	})
	i := iter.Flatten[int](data).TakeWhile(isNeg)

	fmt.Println(*i.Next())
	fmt.Println(*i.Next())
	fmt.Println(i.Next())
	// Output:
	// -1
	// -2
	// <nil>
}

func ExampleFlat_TakeWhile_mapped() {
	toIter := func(a []int) iter.Iterable[int] { return iter.New(a) }
	data := [][]int{{-1, -2}, {3, 4}}
	m := iter.Map[[]int](iter.New(data), toIter)

	isNeg := func(a int) bool { return a < 0 }
	i := iter.Flatten[int](m).TakeWhile(isNeg)

	fmt.Println(*i.Next())
	fmt.Println(*i.Next())
	fmt.Println(i.Next())
	// Output:
	// -1
	// -2
	// <nil>
}

func ExampleFlat_Skip() {
	data := iter.New([]iter.Iterable[int]{
		iter.New([]int{-1, -2}),
		iter.New([]int{-3, 4}),
	})
	i := iter.Flatten[int](data).Skip(1)

	for val := i.Next(); val != nil; val = i.Next() {
		fmt.Println(*val)
	}
	// Output:
	// -2
	// -3
	// 4
}

func ExampleFlat_Skip_mapped() {
	toIter := func(a []int) iter.Iterable[int] { return iter.New(a) }
	data := [][]int{{1, -2}, {-3, 4}}
	m := iter.Map[[]int](iter.New(data), toIter)
	i := iter.Flatten[int](m).Skip(1)

	for val := i.Next(); val != nil; val = i.Next() {
		fmt.Println(*val)
	}
	// Output:
	// -2
	// -3
	// 4
}

func ExampleFlat_Take() {
	data := iter.New([]iter.Iterable[int]{
		iter.New([]int{1, 2}),
		iter.New([]int{3, 4}),
	})
	i := iter.Flatten[int](data).Take(1)

	for val := i.Next(); val != nil; val = i.Next() {
		fmt.Println(*val)
	}
	// Output:
	// 1
}

func ExampleFlat_Take_mapped() {
	toIter := func(a []int) iter.Iterable[int] { return iter.New(a) }
	data := [][]int{{1, 2}, {3, 4}}
	m := iter.Map[[]int](iter.New(data), toIter)
	i := iter.Flatten[int](m).Take(1)

	for val := i.Next(); val != nil; val = i.Next() {
		fmt.Println(*val)
	}
	// Output:
	// 1
}

func ExampleFlat_Collect() {
	data := iter.New([]iter.Iterable[int]{
		iter.New([]int{1, 2}),
		iter.New([]int{3, 4}),
	})
	i := iter.Flatten[int](data)

	copied := i.Collect()
	fmt.Println(copied)

	// Collected slice should contain copies
	copied[0] = 42
	fmt.Println(copied)
	// Output:
	// [1 2 3 4]
	// [42 2 3 4]
}

func ExampleFlat_Collect_mapped() {
	toIter := func(a []int) iter.Iterable[int] { return iter.New(a) }
	data := [][]int{{1, 2}, {3, 4}}
	m := iter.Map[[]int](iter.New(data), toIter)
	i := iter.Flatten[int](m)

	copied := i.Collect()
	fmt.Println(copied)

	// Collected slice should contain copies
	copied[0] = 42
	fmt.Println(copied)
	// Output:
	// [1 2 3 4]
	// [42 2 3 4]
}

func ExampleFlat_ForEach() {
	data := iter.New([]iter.Iterable[int]{
		iter.New([]int{1, 2}),
		iter.New([]int{3, 4}),
	})

	iter.Flatten[int](data).
		ForEach(func(i int) { fmt.Println(i) })
	// Output:
	// 1
	// 2
	// 3
	// 4
}

func ExampleFlat_ForEach_mapped() {
	toIter := func(a []int) iter.Iterable[int] { return iter.New(a) }
	data := [][]int{{1, 2}, {3, 4}}
	m := iter.Map[[]int](iter.New(data), toIter)

	iter.Flatten[int](m).
		ForEach(func(i int) { fmt.Println(i) })
	// Output:
	// 1
	// 2
	// 3
	// 4
}

func ExampleFlat_Nth() {
	data := iter.New([]iter.Iterable[int]{
		iter.New([]int{1, -2}),
		iter.New([]int{3, 4}),
	})
	i := iter.Flatten[int](data).
		Nth(1)

	fmt.Println(*i)
	// Output:
	// -2
}

func ExampleFlat_Nth_mapped() {
	toIter := func(a []int) iter.Iterable[int] { return iter.New(a) }
	data := [][]int{{1, -2}, {3, 4}}
	m := iter.Map[[]int](iter.New(data), toIter)
	i := iter.Flatten[int](m)

	fmt.Println(*i.Nth(1))
	// Output:
	// -2
}

func ExampleFlat_All() {
	gt0 := func(a int) bool { return a > 0 }
	data := iter.New([]iter.Iterable[int]{
		iter.New([]int{1, 2}),
		iter.New([]int{3}),
	})
	t := iter.Flatten[int](data).
		All(gt0)
	fmt.Println(t)
	// Output:
	// true
}

func ExampleFlat_All_mapped() {
	toIter := func(a []int) iter.Iterable[int] { return iter.New(a) }
	data := [][]int{{1, 2}, {3, 4}}

	gt0 := func(a int) bool { return a > 0 }
	m1 := iter.Map[[]int](iter.New(data), toIter)
	t := iter.Flatten[int](m1).All(gt0)
	fmt.Println(t)

	gt2 := func(a int) bool { return a > 2 }
	m2 := iter.Map[[]int](iter.New(data), toIter)
	i := iter.Flatten[int](m2)
	f := i.All(gt2)
	fmt.Println(f)

	// All stops at the first false, so there are still more elements
	fmt.Println(*i.Next())
	fmt.Println(*i.Next())
	// Output:
	// true
	// false
	// 2
	// 3
}

func ExampleFlat_Any() {
	gt0 := func(a int) bool { return a > 0 }
	data := iter.New([]iter.Iterable[int]{
		iter.New([]int{1, 2}),
		iter.New([]int{3}),
	})
	t := iter.Flatten[int](data).
		Any(gt0)
	fmt.Println(t)
	// Output:
	// true
}

func ExampleFlat_Any_mapped() {
	toIter := func(a []int) iter.Iterable[int] { return iter.New(a) }
	data := [][]int{{1, 2}, {3, 4}}

	gt0 := func(a int) bool { return a > 0 }
	m1 := iter.Map[[]int](iter.New(data), toIter)
	t := iter.Flatten[int](m1).Any(gt0)
	fmt.Println(t)

	ne2 := func(a int) bool { return a != 2 }
	m2 := iter.Map[[]int](iter.New(data), toIter)
	i := iter.Flatten[int](m2)
	f := i.Any(ne2)
	fmt.Println(f)

	// Any stops at the first true, so there are still more elements
	fmt.Println(*i.Next())
	fmt.Println(*i.Next())
	// Output:
	// true
	// true
	// 2
	// 3
}

func ExampleFlat_Last_mapped() {
	toIter := func(a []int) iter.Iterable[int] { return iter.New(a) }
	data := [][]int{{1, 2}, {3, 4}}
	m := iter.Map[[]int](iter.New(data), toIter)
	i := iter.Flatten[int](m)

	fmt.Println(*i.Last())
	// Output:
	// 4
}

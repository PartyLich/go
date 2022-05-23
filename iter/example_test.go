package iter_test

import (
	"fmt"

	"github.com/partylich/go/iter"
)

func ExampleNew() {
	slice := []int{1, 2, 3}
	i := iter.New(slice)

	fmt.Println(*i.Next())
	fmt.Println(*i.Next())
	fmt.Println(*i.Next())
	fmt.Println(i.Next())
	// Output:
	// 1
	// 2
	// 3
	// <nil>
}

func ExampleIterable_Next() {
	slice := []int{1, 2, 3}
	i := iter.New(slice)

	fmt.Println(*i.Next())
	fmt.Println(*i.Next())
	fmt.Println(*i.Next())
	fmt.Println(i.Next())
	// Output:
	// 1
	// 2
	// 3
	// <nil>
}

func ExampleIterable_Find() {
	list := []int{1, 2, 3, 4}
	isTwo := func(i int) bool { return i == 2 }
	i := iter.New(list)

	fmt.Println(*i.Find(isTwo))
	fmt.Println(i.Find(isTwo))
	// Output:
	// 2
	// <nil>
}

func ExampleIterator_Rev() {
	slice := []int{1, 2, 3}
	i := iter.New(slice).Rev()

	fmt.Println(*i.Next())
	fmt.Println(*i.Next())
	fmt.Println(*i.Next())
	fmt.Println(i.Next())
	// Output:
	// 3
	// 2
	// 1
	// <nil>
}

func ExampleMap() {
	list := []int{1, 2, 3, 4}
	i := iter.New(list)

	doubled := iter.Map[int, int](i, func(n int) int { return n * 2 })

	for val := doubled.Next(); val != nil; val = doubled.Next() {
		fmt.Println(*val)
	}
	// Output:
	// 2
	// 4
	// 6
	// 8
}

func Example_compose() {
	list := []int{1, 2, 3, 4}
	i := iter.New(list)

	doubled := iter.Map[int, int](i, func(n int) int { return n * 2 })
	f := iter.Filter[int](doubled, func(n int) bool { return n > 5 })

	for val := f.Next(); val != nil; val = f.Next() {
		fmt.Println(*val)
	}
	// Output:
	// 6
	// 8
}

func Example_compose_fluent() {
	list := []int{1, 2, 3, 4}
	double := func(n int) int { return n * 2 }
	gt5 := func(n int) bool { return n > 5 }

	i := iter.New(list)
	iter.Map[int, int](i, double).
		Filter(gt5).
		ForEach(func(val int) { fmt.Println(val) })
	// Output:
	// 6
	// 8
}

func ExampleCollect() {
	list := []int{1, 2, 3, 4}
	i := iter.New(list)

	doubled := iter.Map[int, int](i, func(n int) int { return n * 2 }).
		Filter(func(n int) bool { return n > 5 })

	fmt.Println(iter.Collect[int](doubled))
	// Output:
	// [6 8]
}

func ExampleCollect_fluent() {
	list := []int{1, 2, 3, 4}
	i := iter.New(list)

	doubled := iter.Map[int, int](i, func(n int) int { return n * 2 }).
		Filter(func(n int) bool { return n > 5 }).
		Collect()

	fmt.Println(doubled)
	// Output:
	// [6 8]
}

func ExampleForEach() {
	list := []int{1, 2, 3, 4}
	fn := func(i int) {
		fmt.Println(i)
	}

	i := iter.New(list)
	iter.ForEach[int](i, fn)
	// Output:
	// 1
	// 2
	// 3
	// 4
}

func ExampleForEach_inline() {
	list := []int{1, 2, 3, 4}

	i := iter.New(list)
	iter.ForEach[int](i, func(i int) {
		fmt.Println(i)
	})
	// Output:
	// 1
	// 2
	// 3
	// 4
}

func ExampleForEach_fluent() {
	list := []int{1, 2, 3, 4}

	iter.New(list).
		ForEach(func(i int) { fmt.Println(i) })
	// Output:
	// 1
	// 2
	// 3
	// 4
}

func ExampleNth() {
	list := []int{1, 2, 3, 4}
	i := iter.New(list)

	fmt.Println(*iter.Nth[int](i, 0))
	fmt.Println(*iter.Nth[int](i, 0))
	// i only has 2 elements left, so n = 2 is out of range
	fmt.Println(iter.Nth[int](i, 2))
	// Output:
	// 1
	// 2
	// <nil>
}

func ExampleNth_fluent() {
	list := []int{1, 2, 3, 4}
	n := iter.New(list).Nth(2)

	fmt.Println(*n)
	// Output:
	// 3
}

func ExampleChain() {
	a1 := []int{1, 2}
	a2 := []int{3, 4}

	iter := iter.Chain[int](iter.New(a1), iter.New(a2))

	fmt.Println(*iter.Next())
	fmt.Println(*iter.Next())
	fmt.Println(*iter.Next())
	fmt.Println(*iter.Next())
	fmt.Println(iter.Next())
	// Output:
	// 1
	// 2
	// 3
	// 4
	// <nil>
}

func ExampleTakeWhile() {
	list := []int{-1, -2, 3, 4}
	isNeg := func(a int) bool { return a < 0 }

	iter := iter.TakeWhile[int](iter.New(list), isNeg)

	for val := iter.Next(); val != nil; val = iter.Next() {
		fmt.Println(*val)
	}
	fmt.Println(iter.Next())
	// Output:
	// -1
	// -2
	// <nil>
}

func ExampleSkipWhile() {
	list := []int{-1, -2, 3, 4}
	isNeg := func(a int) bool { return a < 0 }

	iter := iter.SkipWhile[int](iter.New(list), isNeg)

	for val := iter.Next(); val != nil; val = iter.Next() {
		fmt.Println(*val)
	}
	fmt.Println(iter.Next())
	// Output:
	// 3
	// 4
	// <nil>
}

func ExampleCount() {
	i := iter.New([]int{1, 2, 3, 4, 5})

	fmt.Println(iter.Count[int](i))
	// Output:
	// 5
}

func Example_filterCount() {
	list := []int{1, 2, 3, 4}
	i := iter.New(list)

	f := iter.Filter[int](i, func(n int) bool { return n%2 == 0 })

	fmt.Println(iter.Count[int](f))
	// Output:
	// 2
}

func Example_filterCount_fluent() {
	list := []int{1, 2, 3, 4}
	c := iter.New(list).
		Filter(func(n int) bool { return n%2 == 0 }).
		Count()

	fmt.Println(c)
	// Output:
	// 2
}

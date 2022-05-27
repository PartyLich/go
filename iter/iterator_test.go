package iter_test

import (
	"fmt"

	"github.com/partylich/go/iter"
)

func ExampleIterator_Find() {
	list := []int{1, 2, 3, 4}
	isTwo := func(i int) bool { return i == 2 }
	i := iter.New(list)

	fmt.Println(*i.Find(isTwo))
	fmt.Println(i.Find(isTwo))
	// Output:
	// 2
	// <nil>
}

func ExampleIterator_Count() {
	i := iter.New([]int{1, 2, 3, 4, 5})

	fmt.Println(i.Count())
	// Output:
	// 5
}

func ExampleIterator_Filter() {
	i := iter.New([]int{1, 2, 3, 4})

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

func ExampleIterator_SkipWhile() {
	list := []int{-1, 2, 3, 4}
	isNeg := func(a int) bool { return a < 0 }

	i := iter.New(list).SkipWhile(isNeg)

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

func ExampleIterator_Partition() {
	list := []int{1, 2, 3, 4}
	isEven := func(a int) bool { return a%2 == 0 }

	a, b := iter.New(list).Partition(isEven)

	for _, v := range a {
		fmt.Println(v)
	}
	for _, v := range b {
		fmt.Println(v)
	}
	// Output:
	// 2
	// 4
	// 1
	// 3
}

func ExampleIterator_Chain() {
	a1 := []int{1, 2, 3}
	a2 := []int{4, 5, 6}

	i := iter.New(a1).Chain(iter.New(a2))

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

func ExampleIterator_StepBy() {
	list := []int{1, 2, 3, 4, 5, 6}
	i := iter.New(list).StepBy(2)

	for val := i.Next(); val != nil; val = i.Next() {
		fmt.Println(*val)
	}
	// Output:
	// 1
	// 3
	// 5
}

func ExampleIterator_TakeWhile() {
	list := []int{-1, -2, 3, 4}
	isNeg := func(a int) bool { return a < 0 }

	i := iter.New(list).TakeWhile(isNeg)
	fmt.Println(*i.Next())
	fmt.Println(*i.Next())
	fmt.Println(i.Next())
	// Output:
	// -1
	// -2
	// <nil>
}

func ExampleIterator_Skip() {
	list := []int{-1, -2, -3, 4}

	i := iter.New(list).Skip(1)

	for val := i.Next(); val != nil; val = i.Next() {
		fmt.Println(*val)
	}
	// Output:
	// -2
	// -3
	// 4
}

func ExampleIterator_Take() {
	list := []int{-1, -2, -3, 4}
	i := iter.New(list).Take(1)

	for val := i.Next(); val != nil; val = i.Next() {
		fmt.Println(*val)
	}
	// Output:
	// -1
}

func ExampleIterator_Collect() {
	list := []int{1, 2, 3, 4}
	i := iter.New(list)

	copied := i.Collect()
	fmt.Println(copied)

	// Collected slice should contain copies
	copied[0] = 42
	fmt.Println(list)
	fmt.Println(copied)
	// Output:
	// [1 2 3 4]
	// [1 2 3 4]
	// [42 2 3 4]
}

func ExampleIterator_ForEach() {
	list := []int{1, 2, 3, 4}

	iter.New(list).
		ForEach(func(i int) { fmt.Println(i) })
	// Output:
	// 1
	// 2
	// 3
	// 4
}

func ExampleIterator_Nth() {
	list := []int{-1, -2, -3, 4}
	i := iter.New(list).Nth(1)

	fmt.Println(*i)
	// Output:
	// -2
}

func ExampleIterator_All() {
	gt0 := func(a int) bool { return a > 0 }
	gt2 := func(a int) bool { return a > 2 }
	list := []int{1, 2, 3}

	t := iter.New(list).All(gt0)
	fmt.Println(t)

	i := iter.New(list)
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

func ExampleIterator_Any() {
	gt0 := func(a int) bool { return a > 0 }
	ne2 := func(a int) bool { return a != 2 }
	list := []int{1, 2, 3}

	t := iter.New(list).Any(gt0)
	fmt.Println(t)

	i := iter.New(list)
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

func ExampleIterator_Last() {
	i := iter.New([]int{1, 2, 3, 4, 5})

	fmt.Println(*i.Last())
	// Output:
	// 5
}

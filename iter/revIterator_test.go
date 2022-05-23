package iter_test

import (
	"fmt"

	"github.com/partylich/go/iter"
)

func ExampleRevIterator_Count() {
	i := iter.New([]int{1, 2, 3, 4, 5}).Rev()

	fmt.Println(i.Count())
	// Output:
	// 5
}

func ExampleRevIterator_Filter() {
	list := []int{1, 2, 3, 4}
	f := iter.New(list).
		Rev().
		Filter(func(n int) bool { return n > 2 })

	// alternatively,
	//  gt2 := func(n int) bool { return n > 2 }
	//  f := iter.New(list).Rev().Filter(gt5)

	for val := f.Next(); val != nil; val = f.Next() {
		fmt.Println(*val)
	}
	// Output:
	// 4
	// 3
}

func ExampleRevIterator_SkipWhile() {
	list := []int{-1, 2, 3, 4}
	isNeg := func(a int) bool { return a < 0 }

	i := iter.New(list).Rev().SkipWhile(isNeg)

	fmt.Println(*i.Next())
	fmt.Println(*i.Next())
	fmt.Println(*i.Next())
	fmt.Println(*i.Next())
	fmt.Println(i.Next())
	// Output:
	// 4
	// 3
	// 2
	// -1
	// <nil>
}

func ExampleRevIterator_Partition() {
	list := []int{1, 2, 3, 4}
	isEven := func(a int) bool { return a%2 == 0 }

	a, b := iter.New(list).Rev().Partition(isEven)

	for _, v := range a {
		fmt.Println(v)
	}
	for _, v := range b {
		fmt.Println(v)
	}
	// Output:
	// 4
	// 2
	// 3
	// 1
}

func ExampleRevIterator_Chain() {
	a1 := []int{1, 2, 3}
	a2 := []int{4, 5, 6}

	i := iter.New(a1).Rev().Chain(iter.New(a2))

	for val := i.Next(); val != nil; val = i.Next() {
		fmt.Println(*val)
	}
	// Output:
	// 3
	// 2
	// 1
	// 4
	// 5
	// 6
}

func ExampleRevIterator_StepBy() {
	list := []int{1, 2, 3, 4, 5, 6}
	i := iter.New(list).Rev().StepBy(2)

	for val := i.Next(); val != nil; val = i.Next() {
		fmt.Println(*val)
	}
	// Output:
	// 6
	// 4
	// 2
}

func ExampleRevIterator_TakeWhile() {
	list := []int{-1, -2, 3, 4}
	isPos := func(a int) bool { return a > 0 }

	i := iter.New(list).Rev().TakeWhile(isPos)
	fmt.Println(*i.Next())
	fmt.Println(*i.Next())
	fmt.Println(i.Next())
	// Output:
	// 4
	// 3
	// <nil>
}

func ExampleRevIterator_Skip() {
	list := []int{-1, -2, -3, 4}

	i := iter.New(list).Rev().Skip(1)

	for val := i.Next(); val != nil; val = i.Next() {
		fmt.Println(*val)
	}
	// Output:
	// -3
	// -2
	// -1
}

func ExampleRevIterator_Take() {
	list := []int{-1, -2, -3, 4}
	i := iter.New(list).Rev().Take(1)

	for val := i.Next(); val != nil; val = i.Next() {
		fmt.Println(*val)
	}
	// Output:
	// 4
}

func ExampleRevIterator_ForEach() {
	list := []int{1, 2, 3, 4}

	iter.New(list).Rev().
		ForEach(func(i int) { fmt.Println(i) })
	// Output:
	// 4
	// 3
	// 2
	// 1
}

func ExampleRevIterator_Nth() {
	list := []int{-1, -2, -3, 4}
	i := iter.New(list).Rev().Nth(1)

	fmt.Println(*i)
	// Output:
	// -3
}

func ExampleRevIterator_All() {
	gt0 := func(a int) bool { return a > 0 }
	gt2 := func(a int) bool { return a > 2 }
	list := []int{1, 2, 3}

	t := iter.New(list).Rev().All(gt0)
	fmt.Println(t)

	i := iter.New(list).Rev()
	f := i.All(gt2)
	fmt.Println(f)
	// All stops at the first false, so there are still more elements
	fmt.Println(*i.Next())
	// Output:
	// true
	// false
	// 1
}

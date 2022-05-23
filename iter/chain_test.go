package iter_test

import (
	"fmt"

	"github.com/partylich/go/iter"
)

func ExampleChained_Count() {
	a1 := []int{1, 2, 3}
	a2 := []int{4, 5, 6}

	i := iter.New(a1).Chain(iter.New(a2))

	fmt.Println(i.Count())
	// Output:
	// 6
}

func ExampleChained_Filter() {
	a1 := []int{1, 2, 3}
	a2 := []int{4, 5, 6}

	i := iter.New(a1).Chain(iter.New(a2)).
		Filter(func(n int) bool { return n%2 == 0 })

	// alternatively,
	//  isEven := func(n int) bool { return n%2 == 0 }
	//  i := iter.New(a1).Chain(iter.New(a2)).Filter(isEven)

	for val := i.Next(); val != nil; val = i.Next() {
		fmt.Println(*val)
	}
	// Output:
	// 2
	// 4
	// 6
}

func ExampleChained_Partition() {
	a1 := []int{1, 2, 3}
	a2 := []int{6}
	gt5 := func(a int) bool { return a > 5 }

	a, b := iter.New(a1).Chain(iter.New(a2)).Partition(gt5)

	for _, v := range a {
		fmt.Println(v)
	}
	for _, v := range b {
		fmt.Println(v)
	}
	// Output:
	// 6
	// 1
	// 2
	// 3
}

func ExampleChained_Chain() {
	a1 := []int{1, 2}
	a2 := []int{3, 4}
	a3 := []int{5, 6}

	i := iter.New(a1).Chain(iter.New(a2)).Chain(iter.New(a3))

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

func ExampleChained_StepBy() {
	a1 := []int{2, 3}
	a2 := []int{4, 5, 6}

	i := iter.New(a1).Chain(iter.New(a2)).
		StepBy(2)

	for val := i.Next(); val != nil; val = i.Next() {
		fmt.Println(*val)
	}
	// Output:
	// 2
	// 4
	// 6
}

func ExampleChained_TakeWhile() {
	a1 := []int{1, 2, 3}
	a2 := []int{4, -5, -6}
	isPos := func(a int) bool { return a > 0 }

	i := iter.New(a1).Chain(iter.New(a2)).TakeWhile(isPos)

	for val := i.Next(); val != nil; val = i.Next() {
		fmt.Println(*val)
	}
	// Output:
	// 1
	// 2
	// 3
	// 4
}

func ExampleChained_SkipWhile() {
	a1 := []int{-1, -2, 3}
	a2 := []int{4, 5}
	isNeg := func(a int) bool { return a < 0 }

	i := iter.New(a1).Chain(iter.New(a2)).SkipWhile(isNeg)

	for val := i.Next(); val != nil; val = i.Next() {
		fmt.Println(*val)
	}
	// Output:
	// 3
	// 4
	// 5
}

func ExampleChained_Skip() {
	a1 := []int{1, 2, 3}
	a2 := []int{4, 5}

	i := iter.New(a1).Chain(iter.New(a2)).Skip(2)

	for val := i.Next(); val != nil; val = i.Next() {
		fmt.Println(*val)
	}
	// Output:
	// 3
	// 4
	// 5
}

func ExampleChained_Take() {
	a1 := []int{1, 2, 3}
	a2 := []int{4, 5, 6}

	i := iter.New(a1).Chain(iter.New(a2)).Take(2)

	for val := i.Next(); val != nil; val = i.Next() {
		fmt.Println(*val)
	}
	// Output:
	// 1
	// 2
}

func ExampleChained_ForEach() {
	a1 := []int{1, 2, 3}
	a2 := []int{4, 5, 6}

	iter.New(a1).Chain(iter.New(a2)).
		ForEach(func(i int) { fmt.Println(i) })
	// Output:
	// 1
	// 2
	// 3
	// 4
	// 5
	// 6
}

func ExampleChained_Nth() {
	a1 := []int{1, 2, 3}
	a2 := []int{4, 5, 6}

	i := iter.New(a1).Chain(iter.New(a2)).Nth(3)

	fmt.Println(*i)
	// Output:
	// 4
}

func ExampleChained_All() {
	gt0 := func(a int) bool { return a > 0 }
	gt2 := func(a int) bool { return a > 2 }
	a1 := []int{1, 2, 3}
	a2 := []int{4, 5, 6}

	t := iter.New(a1).Chain(iter.New(a2)).All(gt0)
	fmt.Println(t)

	i := iter.New(a1).Chain(iter.New(a2))
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

func ExampleChained_Any() {
	gt0 := func(a int) bool { return a > 0 }
	ne2 := func(a int) bool { return a != 2 }
	a1 := []int{1, 2}
	a2 := []int{3, 4}

	t := iter.New(a1).Chain(iter.New(a2)).Any(gt0)
	fmt.Println(t)

	i := iter.New(a1).Chain(iter.New(a2))
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

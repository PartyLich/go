package iter_test

import (
	"fmt"

	"github.com/partylich/go/iter"
)

func ExampleStepped_Count() {
	i := iter.New([]int{1, 2, 3, 4, 5}).StepBy(2)

	fmt.Println(i.Count())
	// Output:
	// 3
}

func ExampleStepped_Filter() {
	list := []int{2, 3, 4, 5, 6}
	f := iter.New(list).StepBy(2).
		Filter(func(n int) bool { return n%2 == 0 })

	// alternatively,
	//  isEven := func(n int) bool { return n%2 == 0 }
	//  f := iter.New(list).StepBy(2).Filter(isEven)

	for val := f.Next(); val != nil; val = f.Next() {
		fmt.Println(*val)
	}
	// Output:
	// 2
	// 4
	// 6
}

func ExampleStepped_Partition() {
	list := []int{1, 2, 3, 4, 5, 6}
	gte5 := func(a int) bool { return a >= 5 }

	a, b := iter.New(list).StepBy(2).Partition(gte5)

	fmt.Println(a)
	fmt.Println(b)
	// Output:
	// [5]
	// [1 3]
}

func ExampleStepped_Chain() {
	a1 := []int{1, 2, 3}
	a2 := []int{4, 5, 6}

	i := iter.New(a1).StepBy(2).Chain(iter.New(a2))

	for val := i.Next(); val != nil; val = i.Next() {
		fmt.Println(*val)
	}
	// Output:
	// 1
	// 3
	// 4
	// 5
	// 6
}

func ExampleStepped_TakeWhile() {
	list := []int{1, 2, 6, -7, -3, -4}
	isPos := func(a int) bool { return a > 0 }

	i := iter.New(list).StepBy(2).TakeWhile(isPos)
	fmt.Println(*i.Next())
	fmt.Println(*i.Next())
	fmt.Println(i.Next())
	// Output:
	// 1
	// 6
	// <nil>
}

func ExampleStepped_SkipWhile() {
	list := []int{-1, -2, -3, 4, 5}
	isNeg := func(a int) bool { return a < 0 }

	i := iter.New(list).StepBy(2).SkipWhile(isNeg)

	fmt.Println(*i.Next())
	fmt.Println(i.Next())
	// Output:
	// 5
	// <nil>
}

func ExampleStepped_Skip() {
	list := []int{-1, -2, -3, 4}

	i := iter.New(list).StepBy(2).Skip(1)

	fmt.Println(*i.Next())
	fmt.Println(i.Next())
	// Output:
	// -3
	// <nil>
}

func ExampleStepped_Take() {
	list := []int{-1, -2, -3, 4}

	i := iter.New(list).StepBy(2).Take(1)

	fmt.Println(*i.Next())
	fmt.Println(i.Next())
	// Output:
	// -1
	// <nil>
}

func ExampleStepped_ForEach() {
	list := []int{1, 2, 3, 4}

	iter.New(list).StepBy(2).
		ForEach(func(i int) { fmt.Println(i) })
	// Output:
	// 1
	// 3
}

func ExampleStepped_Nth() {
	list := []int{-1, -2, -3, 4}

	i := iter.New(list).StepBy(2).Nth(1)

	fmt.Println(*i)
	// Output:
	// -3
}

func ExampleStepped_All() {
	gt0 := func(a int) bool { return a > 0 }
	gt2 := func(a int) bool { return a > 2 }
	list := []int{1, 2, 3, 4}

	t := iter.New(list).StepBy(2).All(gt0)
	fmt.Println(t)

	i := iter.New(list).StepBy(2)
	f := i.All(gt2)
	fmt.Println(f)
	// All stops at the first false, so there are still more elements
	fmt.Println(*i.Next())
	// Output:
	// true
	// false
	// 3
}

func ExampleStepped_Any() {
	gt0 := func(a int) bool { return a > 0 }
	ne2 := func(a int) bool { return a != 2 }
	list := []int{1, 2, 3, 4}

	t := iter.New(list).StepBy(2).Any(gt0)
	fmt.Println(t)

	i := iter.New(list).StepBy(2)
	f := i.Any(ne2)
	fmt.Println(f)
	// Any stops at the first true, so there are still more elements
	fmt.Println(*i.Next())
	// Output:
	// true
	// true
	// 3
}

func ExampleStepped_Collect() {
	list := []int{1, 2, 3}
	s := iter.New(list).StepBy(2).Collect()
	fmt.Println(s)
	// Output:
	// [1 3]
}

func ExampleStepped_Last() {
	i := iter.New([]int{1, 2, 3, 4, 5}).StepBy(2)

	fmt.Println(*i.Last())
	// Output:
	// 5
}

package iter_test

import (
	"fmt"

	"github.com/partylich/go/iter"
)

func ExampleSkipped_Count() {
	i := iter.New([]int{1, 2, 3, 4, 5}).Skip(3)

	fmt.Println(i.Count())
	// Output:
	// 2
}

func ExampleSkipped_Filter() {
	list := []int{1, 2, 3, 4}
	f := iter.New(list).Skip(2).
		Filter(func(n int) bool { return n%2 == 0 })

	// alternatively,
	//  isEven := func(n int) bool { return n%2 == 0 }
	//  f := iter.New(list).Skip(2).Filter(isEven)

	for val := f.Next(); val != nil; val = f.Next() {
		fmt.Println(*val)
	}
	// Output:
	// 4
}

func ExampleSkipped_SkipWhile() {
	list := []int{-1, 2, -3, 4}
	isNeg := func(a int) bool { return a < 0 }

	i := iter.New(list).Skip(2).SkipWhile(isNeg)

	fmt.Println(*i.Next())
	fmt.Println(i.Next())
	// Output:
	// 4
	// <nil>
}

func ExampleSkipped_Partition() {
	list := []int{1, 2, 3, 4, 5, 6}
	isEven := func(a int) bool { return a%2 == 0 }

	a, b := iter.New(list).Skip(2).Partition(isEven)

	for _, v := range a {
		fmt.Println(v)
	}
	for _, v := range b {
		fmt.Println(v)
	}
	// Output:
	// 4
	// 6
	// 3
	// 5
}

func ExampleSkipped_Chain() {
	a1 := []int{1, 2, 3}
	a2 := []int{4, 5, 6}

	i := iter.New(a1).Skip(2).Chain(iter.New(a2))

	for val := i.Next(); val != nil; val = i.Next() {
		fmt.Println(*val)
	}
	// Output:
	// 3
	// 4
	// 5
	// 6
}

func ExampleSkipped_StepBy() {
	list := []int{1, 2, 3, 4, 5, 6}
	i := iter.New(list).Skip(2).StepBy(2)

	for val := i.Next(); val != nil; val = i.Next() {
		fmt.Println(*val)
	}
	// Output:
	// 3
	// 5
}

func ExampleSkipped_TakeWhile() {
	list := []int{-1, -2, 6, 7, -3, 4}
	isPos := func(a int) bool { return a > 0 }

	i := iter.New(list).Skip(2).TakeWhile(isPos)
	fmt.Println(*i.Next())
	fmt.Println(*i.Next())
	fmt.Println(i.Next())
	// Output:
	// 6
	// 7
	// <nil>
}

func ExampleSkipped_Take() {
	list := []int{-1, -2, -3, 4}

	i := iter.New(list).Skip(2).Take(1)

	fmt.Println(*i.Next())
	fmt.Println(i.Next())
	// Output:
	// -3
	// <nil>
}

func ExampleSkipped_ForEach() {
	list := []int{1, 2, 3, 4}

	iter.New(list).Skip(2).
		ForEach(func(i int) { fmt.Println(i) })
	// Output:
	// 3
	// 4
}

func ExampleSkipped_Nth() {
	list := []int{-1, -2, -3, 4}

	i := iter.New(list).Skip(2).Nth(1)

	fmt.Println(*i)
	// Output:
	// 4
}

func ExampleSkipped_All() {
	gt0 := func(a int) bool { return a > 0 }
	gt2 := func(a int) bool { return a > 2 }
	list := []int{1, 2, 3}

	t := iter.New(list).Skip(1).All(gt0)
	fmt.Println(t)

	i := iter.New(list).Skip(1)
	f := i.All(gt2)
	fmt.Println(f)
	// All stops at the first false, so there are still more elements
	fmt.Println(*i.Next())
	// Output:
	// true
	// false
	// 3
}

func ExampleSkipped_Any() {
	gt0 := func(a int) bool { return a > 0 }
	ne2 := func(a int) bool { return a != 2 }
	list := []int{1, 2, 3, 4}

	t := iter.New(list).Skip(1).Any(gt0)
	fmt.Println(t)

	i := iter.New(list).Skip(1)
	f := i.Any(ne2)
	fmt.Println(f)
	// Any stops at the first true, so there are still more elements
	fmt.Println(*i.Next())
	// Output:
	// true
	// true
	// 4
}

func ExampleSkipped_Collect() {
	list := []int{1, 2, 3, 4}

	s := iter.New(list).Skip(1).Collect()
	fmt.Println(s)
	// Output:
	// [2 3 4]
}

func ExampleSkipped_Last() {
	i := iter.New([]int{1, 2, 3, 4, 5}).Skip(3)

	fmt.Println(*i.Last())
	// Output:
	// 5
}

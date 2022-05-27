package iter_test

import (
	"fmt"

	"github.com/partylich/go/iter"
)

func ExampleSkipWhileT_Count() {
	isNeg := func(a int) bool { return a < 0 }
	list := []int{-1, -2, 3, 4}
	i := iter.New(list).SkipWhile(isNeg)

	fmt.Println(i.Count())
	// Output:
	// 2
}

func ExampleSkipWhileT_Filter() {
	isNeg := func(a int) bool { return a < 0 }
	list := []int{-1, -2, 3, 4}

	i := iter.New(list).SkipWhile(isNeg).
		Filter(func(n int) bool { return n%2 == 0 })

	// alternatively,
	//  isEven := func(n int) bool { return n%2 == 0 }
	//  i := iter.New(list).SkipWhile(isNeg).Filter(isEven)

	for val := i.Next(); val != nil; val = i.Next() {
		fmt.Println(*val)
	}
	// Output:
	// 4
}

func ExampleSkipWhileT_Partition() {
	isNeg := func(a int) bool { return a < 0 }
	isOdd := func(a int) bool { return a%2 != 0 }

	list := []int{-1, -2, 3, 4}
	a, b := iter.New(list).SkipWhile(isNeg).Partition(isOdd)

	fmt.Println(a)
	fmt.Println(b)
	// Output:
	// [3]
	// [4]
}

func ExampleSkipWhileT_Chain() {
	isNeg := func(a int) bool { return a < 0 }
	list := []int{-1, -2, 3, 4}
	a := iter.New(list).SkipWhile(isNeg)
	b := iter.New([]int{5, 6})

	i := a.Chain(b)

	for val := i.Next(); val != nil; val = i.Next() {
		fmt.Println(*val)
	}
	// Output:
	// 3
	// 4
	// 5
	// 6
}

func ExampleSkipWhileT_StepBy() {
	isNeg := func(a int) bool { return a < 0 }
	list := []int{-1, -2, 3, 4}
	i := iter.New(list).SkipWhile(isNeg).StepBy(2)

	for val := i.Next(); val != nil; val = i.Next() {
		fmt.Println(*val)
	}
	// Output:
	// 3
}

func ExampleSkipWhileT_TakeWhile() {
	isNeg := func(a int) bool { return a < 0 }
	isOdd := func(a int) bool { return a%2 != 0 }
	list := []int{-1, -2, 3, 4}
	i := iter.New(list).SkipWhile(isNeg).TakeWhile(isOdd)

	for val := i.Next(); val != nil; val = i.Next() {
		fmt.Println(*val)
	}
	// Output:
	// 3
}

func ExampleSkipWhileT_SkipWhile() {
	isNeg := func(a int) bool { return a < 0 }
	isOdd := func(a int) bool { return a%2 != 0 }
	list := []int{-1, -2, 3, 4}
	i := iter.New(list).SkipWhile(isNeg).SkipWhile(isOdd)

	for val := i.Next(); val != nil; val = i.Next() {
		fmt.Println(*val)
	}
	// Output:
	// 4
}

func ExampleSkipWhileT_Skip() {
	isNeg := func(a int) bool { return a < 0 }
	list := []int{-1, -2, 3, 4}
	i := iter.New(list).SkipWhile(isNeg).Skip(1)

	for val := i.Next(); val != nil; val = i.Next() {
		fmt.Println(*val)
	}
	// Output:
	// 4
}

func ExampleSkipWhileT_Take() {
	isNeg := func(a int) bool { return a < 0 }
	list := []int{-1, -2, 3, 4}
	i := iter.New(list).SkipWhile(isNeg).Take(2)

	for val := i.Next(); val != nil; val = i.Next() {
		fmt.Println(*val)
	}
	// Output:
	// 3
	// 4
}

func ExampleSkipWhileT_ForEach() {
	isNeg := func(a int) bool { return a < 0 }
	list := []int{-1, -2, 3, 4}

	iter.New(list).SkipWhile(isNeg).
		ForEach(func(i int) { fmt.Println(i) })
	// Output:
	// 3
	// 4
}

func ExampleSkipWhileT_Nth() {
	isNeg := func(a int) bool { return a < 0 }
	list := []int{-1, -2, 3, 4}
	i := iter.New(list).SkipWhile(isNeg).Nth(1)

	fmt.Println(*i)
	// Output:
	// 4
}

func ExampleSkipWhileT_All() {
	isNeg := func(a int) bool { return a < 0 }
	list := []int{-1, -2, 3, 4}

	ne2 := func(a int) bool { return a != 2 }
	t := iter.New(list).SkipWhile(isNeg).All(ne2)
	fmt.Println(t)

	gt5 := func(a int) bool { return a > 5 }
	i := iter.New(list).SkipWhile(isNeg)
	f := i.All(gt5)
	fmt.Println(f)
	// All stops at the first false, so there are still more elements
	fmt.Println(*i.Next())
	// Output:
	// true
	// false
	// 4
}

func ExampleSkipWhileT_Any() {
	isNeg := func(a int) bool { return a < 0 }
	list := []int{-1, -2, 3, 4}

	ne2 := func(a int) bool { return a != 2 }
	t := iter.New(list).SkipWhile(isNeg).Any(ne2)
	fmt.Println(t)

	gt0 := func(a int) bool { return a > 0 }
	i := iter.New(list).SkipWhile(isNeg)
	f := i.Any(gt0)
	fmt.Println(f)

	// Any stops at the first true, so there are still more elements
	fmt.Println(*i.Next())
	// Output:
	// true
	// true
	// 4
}

func ExampleSkipWhileT_Collect() {
	isNeg := func(a int) bool { return a < 0 }
	list := []int{-1, -2, 3, 4}
	s := iter.New(list).SkipWhile(isNeg).Collect()

	fmt.Println(s)
	// Output:
	// [3 4]
}

func ExampleSkipWhileT_Last() {
	isNeg := func(a int) bool { return a < 0 }
	list := []int{-1, -2, 3, 4}
	i := iter.New(list).SkipWhile(isNeg)

	fmt.Println(*i.Last())
	// Output:
	// 4
}

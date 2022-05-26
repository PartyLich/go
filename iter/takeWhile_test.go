package iter_test

import (
	"fmt"

	"github.com/partylich/go/iter"
)

func ExampleTakeWhileT_Count() {
	isNeg := func(a int) bool { return a < 0 }
	list := []int{-1, -2, 3, 4}
	i := iter.New(list).TakeWhile(isNeg)

	fmt.Println(i.Count())
	// Output:
	// 2
}

func ExampleTakeWhileT_Filter() {
	isNeg := func(a int) bool { return a < 0 }
	list := []int{-1, -2, 3, 4}

	i := iter.New(list).TakeWhile(isNeg).
		Filter(func(n int) bool { return n%2 == 0 })

	// alternatively,
	//  isEven := func(n int) bool { return n%2 == 0 }
	//  i := iter.New(a1).Chain(iter.New(a2)).Filter(isEven)

	for val := i.Next(); val != nil; val = i.Next() {
		fmt.Println(*val)
	}
	// Output:
	// -2
}

func ExampleTakeWhileT_Partition() {
	isNeg := func(a int) bool { return a < 0 }
	isOdd := func(a int) bool { return a%2 != 0 }

	list := []int{-1, -2, 3, 4}
	a, b := iter.New(list).TakeWhile(isNeg).Partition(isOdd)

	fmt.Println(a)
	fmt.Println(b)
	// Output:
	// [-1]
	// [-2]
}

func ExampleTakeWhileT_Chain() {
	isNeg := func(a int) bool { return a < 0 }
	list := []int{-1, -2, 3, 4}
	a := iter.New(list).TakeWhile(isNeg)
	b := iter.New([]int{5, 6})

	i := a.Chain(b)

	for val := i.Next(); val != nil; val = i.Next() {
		fmt.Println(*val)
	}
	// Output:
	// -1
	// -2
	// 5
	// 6
}

func ExampleTakeWhileT_StepBy() {
	isNeg := func(a int) bool { return a < 0 }
	list := []int{-1, -2, 3, 4}
	i := iter.New(list).TakeWhile(isNeg).StepBy(2)

	for val := i.Next(); val != nil; val = i.Next() {
		fmt.Println(*val)
	}
	// Output:
	// -1
}

func ExampleTakeWhileT_TakeWhile() {
	isNeg := func(a int) bool { return a < 0 }
	isOdd := func(a int) bool { return a%2 != 0 }
	list := []int{-1, -2, 3, 4}
	i := iter.New(list).TakeWhile(isNeg).TakeWhile(isOdd)

	for val := i.Next(); val != nil; val = i.Next() {
		fmt.Println(*val)
	}
	// Output:
	// -1
}

func ExampleTakeWhileT_SkipWhile() {
	isNeg := func(a int) bool { return a < 0 }
	isOdd := func(a int) bool { return a%2 != 0 }
	list := []int{-1, -2, 3, 4}
	i := iter.New(list).TakeWhile(isNeg).SkipWhile(isOdd)

	for val := i.Next(); val != nil; val = i.Next() {
		fmt.Println(*val)
	}
	// Output:
	// -2
}

func ExampleTakeWhileT_Skip() {
	isNeg := func(a int) bool { return a < 0 }
	list := []int{-1, -2, 3, 4}
	i := iter.New(list).TakeWhile(isNeg).Skip(1)

	for val := i.Next(); val != nil; val = i.Next() {
		fmt.Println(*val)
	}
	// Output:
	// -2
}

func ExampleTakeWhileT_Take() {
	isNeg := func(a int) bool { return a < 0 }
	list := []int{-1, -2, 3, 4}
	i := iter.New(list).TakeWhile(isNeg).Take(2)

	for val := i.Next(); val != nil; val = i.Next() {
		fmt.Println(*val)
	}
	// Output:
	// -1
	// -2
}

func ExampleTakeWhileT_ForEach() {
	isNeg := func(a int) bool { return a < 0 }
	list := []int{-1, -2, 3, 4}

	iter.New(list).TakeWhile(isNeg).
		ForEach(func(i int) { fmt.Println(i) })
	// Output:
	// -1
	// -2
}

func ExampleTakeWhileT_Nth() {
	isNeg := func(a int) bool { return a < 0 }
	list := []int{-1, -2, 3, 4}
	i := iter.New(list).TakeWhile(isNeg).Nth(1)

	fmt.Println(*i)
	// Output:
	// -2
}

func ExampleTakeWhileT_All() {
	isNeg := func(a int) bool { return a < 0 }
	list := []int{-1, -2, 3, 4}

	ne2 := func(a int) bool { return a != 2 }
	t := iter.New(list).TakeWhile(isNeg).All(ne2)
	fmt.Println(t)

	gt2 := func(a int) bool { return a > 2 }
	i := iter.New(list).TakeWhile(isNeg)
	f := i.All(gt2)
	fmt.Println(f)

	// All stops at the first false, so there are still more elements
	fmt.Println(*i.Next())
	// Output:
	// true
	// false
	// -2
}

func ExampleTakeWhileT_Any() {
	isNeg := func(a int) bool { return a < 0 }
	list := []int{-1, -2, 3, 4}

	isOdd := func(a int) bool { return a%2 != 0 }
	t := iter.New(list).TakeWhile(isNeg).Any(isOdd)
	fmt.Println(t)

	ne2 := func(a int) bool { return a != 2 }
	i := iter.New(list).TakeWhile(isNeg)
	f := i.Any(ne2)
	fmt.Println(f)

	// Any stops at the first true, so there are still more elements
	fmt.Println(*i.Next())
	// Output:
	// true
	// true
	// -2
}

func ExampleTakeWhileT_Collect() {
	isNeg := func(a int) bool { return a < 0 }
	list := []int{-1, -2, 3, 4}
	s := iter.New(list).TakeWhile(isNeg).Collect()

	fmt.Println(s)
	// Output:
	// [-1 -2]
}

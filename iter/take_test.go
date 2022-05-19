package iter_test

import (
	"fmt"

	"github.com/PartyLich/go/iter"
)

func ExampleTaken_Count() {
	i := iter.New([]int{1, 2, 3, 4, 5}).Take(3)

	fmt.Println(i.Count())

	// Output:
	// 3
}

func ExampleTaken_Filter() {
	list := []int{1, 2, 3, 4}
	f := iter.New(list).Take(2).
		Filter(func(n int) bool { return n%2 == 0 })

	// alternatively,
	//  isEven := func(n int) bool { return n%2 == 0 }
	//  f := iter.New(list).Take(2).Filter(isEven)

	for val := f.Next(); val != nil; val = f.Next() {
		fmt.Println(*val)
	}

	// Output:
	// 2
}

func ExampleTaken_Partition() {
	list := []int{1, 2, 3, 4, 5, 6}
	isEven := func(a int) bool { return a%2 == 0 }

	a, b := iter.New(list).Take(2).Partition(isEven)

	for _, v := range a {
		fmt.Println(v)
	}
	for _, v := range b {
		fmt.Println(v)
	}

	// Output:
	// 2
	// 1
}

func ExampleTaken_Chain() {
	a1 := []int{1, 2, 3}
	a2 := []int{4, 5, 6}

	i := iter.New(a1).Take(2).Chain(iter.New(a2))

	for val := i.Next(); val != nil; val = i.Next() {
		fmt.Println(*val)
	}

	// Output:
	// 1
	// 2
	// 4
	// 5
	// 6
}

func ExampleTaken_StepBy() {
	list := []int{1, 2, 3, 4, 5, 6}
	i := iter.New(list).Take(2).StepBy(2)

	for val := i.Next(); val != nil; val = i.Next() {
		fmt.Println(*val)
	}

	// Output:
	// 1
}

func ExampleTaken_TakeWhile() {
	list := []int{1, -2, 6, 7, -3, 4}
	isPos := func(a int) bool { return a > 0 }

	i := iter.New(list).Take(2).TakeWhile(isPos)
	fmt.Println(*i.Next())
	fmt.Println(i.Next())

	// Output:
	// 1
	// <nil>
}

func ExampleTaken_SkipWhile() {
	list := []int{-1, 2, -3, 4}
	isNeg := func(a int) bool { return a < 0 }

	i := iter.New(list).Take(2).SkipWhile(isNeg)

	fmt.Println(*i.Next())
	fmt.Println(i.Next())

	// Output:
	// 2
	// <nil>
}

func ExampleTaken_Skip() {
	list := []int{-1, -2, -3, 4}

	i := iter.New(list).Take(3).Skip(1)

	fmt.Println(*i.Next())
	fmt.Println(*i.Next())
	fmt.Println(i.Next())

	// Output:
	// -2
	// -3
	// <nil>
}

func ExampleTaken_ForEach() {
	list := []int{1, 2, 3, 4}

	iter.New(list).Take(3).
		ForEach(func(i int) { fmt.Println(i) })

	// Output:
	// 1
	// 2
	// 3
}

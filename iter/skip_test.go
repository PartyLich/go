package iter_test

import (
	"fmt"

	"github.com/PartyLich/go/iter"
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
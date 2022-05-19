package iter_test

import (
	"fmt"

	"github.com/PartyLich/go/iter"
)

func ExampleStepped_Count() {
	i := iter.New([]int{1, 2, 3, 4, 5}).StepBy(2)

	fmt.Println(i.Count())

	// Output:
	// 2
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
}

func ExampleStepped_Partition() {
	list := []int{1, 2, 3, 4, 5, 6}
	isEven := func(a int) bool { return a%2 == 0 }
	gt5 := func(a int) bool { return a > 5 }

	a, b := iter.New(list).StepBy(2).Partition(gt5)

	for _, v := range a {
		fmt.Println(v)
	}
	for _, v := range b {
		fmt.Println(v)
	}

	// Output:
	// 6
	// 2
	// 4
}

func ExampleStepped_Chain() {
	isEven := func(a int) bool { return a%2 == 0 }
	a1 := []int{1, 2, 3}
	a2 := []int{4, 5, 6}

	i := iter.New(a1).StepBy(2).Chain(iter.New(a2))

	for val := i.Next(); val != nil; val = i.Next() {
		fmt.Println(*val)
	}

	// Output:
	// 2
	// 4
	// 5
	// 6
}

func ExampleStepped_TakeWhile() {
	isEven := func(a int) bool { return a%2 == 0 }
	list := []int{1, 2, 6, -7, -3, -4}
	isPos := func(a int) bool { return a > 0 }

	i := iter.New(list).StepBy(2).TakeWhile(isPos)
	fmt.Println(*i.Next())
	fmt.Println(*i.Next())
	fmt.Println(i.Next())

	// Output:
	// 2
	// 6
	// <nil>
}

func ExampleStepped_SkipWhile() {
	isEven := func(a int) bool { return a%2 == 0 }
	list := []int{-1, -2, -3, 4}
	isNeg := func(a int) bool { return a < 0 }

	i := iter.New(list).StepBy(2).SkipWhile(isNeg)

	fmt.Println(*i.Next())
	fmt.Println(i.Next())

	// Output:
	// 4
	// <nil>
}

func ExampleStepped_Skip() {
	isEven := func(a int) bool { return a%2 == 0 }
	list := []int{-1, -2, -3, 4}

	i := iter.New(list).StepBy(2).Skip(1)

	fmt.Println(*i.Next())
	fmt.Println(i.Next())

	// Output:
	// 4
	// <nil>
}

func ExampleStepped_Take() {
	isEven := func(a int) bool { return a%2 == 0 }
	list := []int{-1, -2, -3, 4}

	i := iter.New(list).StepBy(2).Take(1)

	fmt.Println(*i.Next())
	fmt.Println(i.Next())

	// Output:
	// -2
	// <nil>
}

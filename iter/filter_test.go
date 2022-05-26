package iter_test

import (
	"fmt"

	"github.com/partylich/go/iter"
)

func ExampleFiltered_Count() {
	isEven := func(i int) bool { return i%2 == 0 }
	i := iter.New([]int{1, 2, 3, 4, 5}).Filter(isEven)

	fmt.Println(i.Count())
	// Output:
	// 2
}

// func ExampleFiltered_Filter() {
// 	list := []int{1, 2, 3, 4}
// 	f := iter.New(list).Take(2).
// 		Filter(func(n int) bool { return n%2 == 0 })

// 	// alternatively,
// 	//  isEven := func(n int) bool { return n%2 == 0 }
// 	//  f := iter.New(list).Take(2).Filter(isEven)

// 	for val := f.Next(); val != nil; val = f.Next() {
// 		fmt.Println(*val)
// 	}

// 	// Output:
// 	// 2
// }

func ExampleFiltered_Partition() {
	list := []int{1, 2, 3, 4, 5, 6}
	isEven := func(a int) bool { return a%2 == 0 }
	gt5 := func(a int) bool { return a > 5 }

	a, b := iter.New(list).Filter(isEven).Partition(gt5)

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

func ExampleFiltered_Chain() {
	isEven := func(a int) bool { return a%2 == 0 }
	a1 := []int{1, 2, 3}
	a2 := []int{4, 5, 6}

	i := iter.New(a1).Filter(isEven).Chain(iter.New(a2))

	for val := i.Next(); val != nil; val = i.Next() {
		fmt.Println(*val)
	}
	// Output:
	// 2
	// 4
	// 5
	// 6
}

func ExampleFiltered_StepBy() {
	isEven := func(a int) bool { return a%2 == 0 }
	list := []int{1, 2, 3, 4, 5, 6}
	i := iter.New(list).Filter(isEven).StepBy(2)

	for val := i.Next(); val != nil; val = i.Next() {
		fmt.Println(*val)
	}
	// Output:
	// 2
	// 6
}

func ExampleFiltered_TakeWhile() {
	isEven := func(a int) bool { return a%2 == 0 }
	list := []int{1, 2, 6, -7, -3, -4}
	isPos := func(a int) bool { return a > 0 }

	i := iter.New(list).Filter(isEven).TakeWhile(isPos)
	fmt.Println(*i.Next())
	fmt.Println(*i.Next())
	fmt.Println(i.Next())
	// Output:
	// 2
	// 6
	// <nil>
}

func ExampleFiltered_SkipWhile() {
	isEven := func(a int) bool { return a%2 == 0 }
	list := []int{-1, -2, -3, 4}
	isNeg := func(a int) bool { return a < 0 }

	i := iter.New(list).Filter(isEven).SkipWhile(isNeg)

	fmt.Println(*i.Next())
	fmt.Println(i.Next())
	// Output:
	// 4
	// <nil>
}

func ExampleFiltered_Skip() {
	isEven := func(a int) bool { return a%2 == 0 }
	list := []int{-1, -2, -3, 4}

	i := iter.New(list).Filter(isEven).Skip(1)

	fmt.Println(*i.Next())
	fmt.Println(i.Next())
	// Output:
	// 4
	// <nil>
}

func ExampleFiltered_Take() {
	isEven := func(a int) bool { return a%2 == 0 }
	list := []int{-1, -2, -3, 4}

	i := iter.New(list).Filter(isEven).Take(1)

	fmt.Println(*i.Next())
	fmt.Println(i.Next())
	// Output:
	// -2
	// <nil>
}

func ExampleFiltered_ForEach() {
	isEven := func(a int) bool { return a%2 == 0 }
	list := []int{1, 2, 3, 4}

	iter.New(list).Filter(isEven).
		ForEach(func(i int) { fmt.Println(i) })
	// Output:
	// 2
	// 4
}

func ExampleFiltered_Nth() {
	isEven := func(a int) bool { return a%2 == 0 }
	list := []int{-1, -2, -3, 4}

	i := iter.New(list).Filter(isEven).Nth(1)

	fmt.Println(*i)
	// Output:
	// 4
}

func ExampleFiltered_All() {
	gt0 := func(a int) bool { return a > 0 }
	gt2 := func(a int) bool { return a > 2 }
	list := []int{1, 2, 3, 4}

	t := iter.New(list).Filter(gt2).All(gt0)
	fmt.Println(t)

	i := iter.New(list).Filter(gt0)
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

func ExampleFiltered_Any() {
	gt0 := func(a int) bool { return a > 0 }
	gt2 := func(a int) bool { return a > 2 }
	list := []int{1, 2, 3, 4}

	t := iter.New(list).Filter(gt2).Any(gt0)
	fmt.Println(t)

	i := iter.New(list).Filter(gt0)
	f := i.Any(gt2)
	fmt.Println(f)
	// Any stops at the first true, so there are still more elements
	fmt.Println(*i.Next())
	// Output:
	// true
	// true
	// 4
}

func ExampleFiltered_Last() {
	isEven := func(i int) bool { return i%2 == 0 }
	i := iter.New([]int{1, 2, 3, 4, 5}).Filter(isEven)

	fmt.Println(*i.Last())
	// Output:
	// 4
}

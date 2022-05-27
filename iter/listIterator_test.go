package iter_test

import (
	"container/list"
	"fmt"

	"github.com/partylich/go/iter"
)

func ExampleListIterator_Next() {
	l := list.New()
	// fill linked list with 1 to 4
	for i := 1; i <= 4; i++ {
		l.PushBack(i)
	}

	i := iter.FromList[int](l)

	fmt.Println(*i.Next())
	fmt.Println(*i.Next())
	fmt.Println(*i.Next())
	fmt.Println(*i.Next())
	fmt.Println(i.Next())
	// Output:
	// 1
	// 2
	// 3
	// 4
	// <nil>
}

func ExampleListIterator_Find() {
	isTwo := func(i int) bool { return i == 2 }
	l := list.New()
	// fill linked list with 1 to 4
	for i := 1; i <= 4; i++ {
		l.PushBack(i)
	}

	i := iter.FromList[int](l)

	fmt.Println(*i.Find(isTwo))
	fmt.Println(i.Find(isTwo))
	// Output:
	// 2
	// <nil>
}

func ExampleListIterator_Count() {
	l := list.New()
	// fill linked list with 1 to 4
	for i := 1; i <= 4; i++ {
		l.PushBack(i)
	}

	i := iter.FromList[int](l)

	fmt.Println(i.Count())
	// Output:
	// 4
}

func ExampleListIterator_Filter() {
	l := list.New()
	// fill linked list with 1 to 4
	for i := 1; i <= 4; i++ {
		l.PushBack(i)
	}

	i := iter.FromList[int](l)

	f := i.Filter(func(n int) bool { return n > 2 })

	// alternatively,
	//  gt2 := func(n int) bool { return n > 2 }
	//  f := i.Filter(gt2)

	for val := f.Next(); val != nil; val = f.Next() {
		fmt.Println(*val)
	}
	// Output:
	// 3
	// 4
}

func ExampleListIterator_SkipWhile() {
	l := list.New()
	// fill linked list with -1 to 2
	for i := -1; i <= 2; i++ {
		l.PushBack(i)
	}

	isNeg := func(a int) bool { return a < 0 }

	i := iter.FromList[int](l).SkipWhile(isNeg)

	fmt.Println(*i.Next())
	fmt.Println(*i.Next())
	fmt.Println(*i.Next())
	fmt.Println(i.Next())
	// Output:
	// 0
	// 1
	// 2
	// <nil>
}

func ExampleListIterator_Partition() {
	l := list.New()
	// fill linked list with 1 to 4
	for i := 1; i <= 4; i++ {
		l.PushBack(i)
	}

	isEven := func(a int) bool { return a%2 == 0 }

	a, b := iter.FromList[int](l).Partition(isEven)

	for _, v := range a {
		fmt.Println(v)
	}
	for _, v := range b {
		fmt.Println(v)
	}
	// Output:
	// 2
	// 4
	// 1
	// 3
}

func ExampleListIterator_Chain() {
	l1, l2 := list.New(), list.New()
	// fill linked lists
	for i := 1; i <= 2; i++ {
		l1.PushBack(i)
	}
	for i := 3; i <= 4; i++ {
		l1.PushBack(i)
	}

	i := iter.FromList[int](l1).Chain(iter.FromList[int](l2))

	for val := i.Next(); val != nil; val = i.Next() {
		fmt.Println(*val)
	}
	// Output:
	// 1
	// 2
	// 3
	// 4
}

func ExampleListIterator_StepBy() {
	l := list.New()
	// fill linked list with 1 to 6
	for i := 1; i <= 6; i++ {
		l.PushBack(i)
	}

	i := iter.FromList[int](l).StepBy(2)

	for val := i.Next(); val != nil; val = i.Next() {
		fmt.Println(*val)
	}
	// Output:
	// 1
	// 3
	// 5
}

func ExampleListIterator_TakeWhile() {
	l := list.New()
	// fill linked list with 1 to 4
	for i := -1; i <= 1; i++ {
		l.PushBack(i)
	}

	isNeg := func(a int) bool { return a < 0 }

	i := iter.FromList[int](l).TakeWhile(isNeg)
	fmt.Println(*i.Next())
	fmt.Println(i.Next())
	// Output:
	// -1
	// <nil>
}

func ExampleListIterator_Skip() {
	l := list.New()
	// fill linked list with 1 to 4
	for i := 1; i <= 4; i++ {
		l.PushBack(i)
	}

	i := iter.FromList[int](l).Skip(1)

	for val := i.Next(); val != nil; val = i.Next() {
		fmt.Println(*val)
	}
	// Output:
	// 2
	// 3
	// 4
}

func ExampleListIterator_Take() {
	l := list.New()
	// fill linked list with 1 to 4
	for i := 1; i <= 4; i++ {
		l.PushBack(i)
	}

	i := iter.FromList[int](l).Take(1)

	for val := i.Next(); val != nil; val = i.Next() {
		fmt.Println(*val)
	}
	// Output:
	// 1
}

func ExampleListIterator_Collect() {
	l := list.New()
	// fill linked list with 1 to 4
	for i := 1; i <= 4; i++ {
		l.PushBack(i)
	}

	i := iter.FromList[int](l)

	copied := i.Collect()
	fmt.Println(copied)

	// Collected slice should contain copies
	copied[0] = 42
	fmt.Println(l.Front().Value)
	fmt.Println(copied)
	// Output:
	// [1 2 3 4]
	// 1
	// [42 2 3 4]
}

func ExampleListIterator_ForEach() {
	l := list.New()
	// fill linked list with 1 to 4
	for i := 1; i <= 4; i++ {
		l.PushBack(i)
	}

	iter.FromList[int](l).
		ForEach(func(i int) { fmt.Println(i) })
	// Output:
	// 1
	// 2
	// 3
	// 4
}

func ExampleListIterator_Nth() {
	l := list.New()
	// fill linked list with 1 to 4
	for i := 1; i <= 4; i++ {
		l.PushBack(i)
	}

	i := iter.FromList[int](l).Nth(1)

	fmt.Println(*i)
	// Output:
	// 2
}

func ExampleListIterator_All() {
	l := list.New()
	// fill linked list with 1 to 4
	for i := 1; i <= 4; i++ {
		l.PushBack(i)
	}

	gt0 := func(a int) bool { return a > 0 }
	gt2 := func(a int) bool { return a > 2 }

	t := iter.FromList[int](l).All(gt0)
	fmt.Println(t)

	i := iter.FromList[int](l)
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

func ExampleListIterator_Any() {
	l := list.New()
	// fill linked list with 1 to 4
	for i := 1; i <= 4; i++ {
		l.PushBack(i)
	}

	gt0 := func(a int) bool { return a > 0 }
	ne2 := func(a int) bool { return a != 2 }

	t := iter.FromList[int](l).Any(gt0)
	fmt.Println(t)

	i := iter.FromList[int](l)
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

func ExampleListIterator_Last() {
	l := list.New()
	// fill linked list with 1 to 4
	for i := 1; i <= 4; i++ {
		l.PushBack(i)
	}

	i := iter.FromList[int](l)

	fmt.Println(*i.Last())
	// Output:
	// 4
}

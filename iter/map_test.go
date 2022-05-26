package iter_test

import (
	"fmt"

	"github.com/partylich/go/iter"
)

func ExampleMapped_Count() {
	ident := func(i int) int { return i }
	i := iter.New([]int{1, 2, 3, 4, 5})

	fmt.Println(iter.Map[int, int](i, ident).Count())
	// Output:
	// 5
}

func ExampleMapped_Filter() {
	ident := func(i int) int { return i }
	i := iter.New([]int{1, 2, 3, 4})

	f := iter.Map[int, int](i, ident).Filter(func(n int) bool { return n > 2 })

	// alternatively,
	//  gt2 := func(n int) bool { return n > 2 }
	//  f := iter.Map(i, ident).Filter(gt5)

	for val := f.Next(); val != nil; val = f.Next() {
		fmt.Println(*val)
	}
	// Output:
	// 3
	// 4
}

func ExampleMapped_SkipWhile() {
	list := []int{-1, 2, 3, 4}
	isNeg := func(a int) bool { return a < 0 }
	ident := func(i int) int { return i }

	i := iter.Map[int, int](iter.New(list), ident).SkipWhile(isNeg)

	fmt.Println(*i.Next())
	fmt.Println(*i.Next())
	fmt.Println(*i.Next())
	fmt.Println(i.Next())
	// Output:
	// 2
	// 3
	// 4
	// <nil>
}

func ExampleMapped_Partition() {
	list := []int{1, 2, 3, 4}
	isEven := func(a int) bool { return a%2 == 0 }
	ident := func(i int) int { return i }

	i := iter.New(list)
	a, b := iter.Map[int, int](i, ident).Partition(isEven)

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

func ExampleMapped_Chain() {
	ident := func(i int) int { return i }

	a1 := iter.New([]int{1, 2, 3})
	a2 := iter.New([]int{4, 5, 6})

	i := iter.Map[int, int](a1, ident).Chain(a2)

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

func ExampleMapped_StepBy() {
	ident := func(i int) int { return i }
	list := []int{1, 2, 3, 4, 5, 6}

	i := iter.New(list)
	m := iter.Map[int, int](i, ident).StepBy(2)

	for val := m.Next(); val != nil; val = m.Next() {
		fmt.Println(*val)
	}
	// Output:
	// 1
	// 3
	// 5
}

func ExampleMapped_TakeWhile() {
	list := []int{-1, -2, 3, 4}
	ident := func(i int) int { return i }
	isNeg := func(a int) bool { return a < 0 }

	i := iter.New(list)
	m := iter.Map[int, int](i, ident).TakeWhile(isNeg)

	fmt.Println(*m.Next())
	fmt.Println(*m.Next())
	fmt.Println(m.Next())
	// Output:
	// -1
	// -2
	// <nil>
}

func ExampleMapped_Skip() {
	list := []int{-1, 2, 3, 4}
	ident := func(i int) int { return i }

	i := iter.Map[int, int](iter.New(list), ident).Skip(2)

	fmt.Println(*i.Next())
	fmt.Println(*i.Next())
	fmt.Println(i.Next())
	// Output:
	// 3
	// 4
	// <nil>
}

func ExampleMapped_Take() {
	list := []int{-1, -2, 3, 4}
	ident := func(i int) int { return i }

	i := iter.New(list)
	m := iter.Map[int, int](i, ident).Take(2)

	fmt.Println(*m.Next())
	fmt.Println(*m.Next())
	fmt.Println(m.Next())
	// Output:
	// -1
	// -2
	// <nil>
}

func ExampleMapped_ForEach() {
	list := []int{1, 2, 3, 4}
	ident := func(i int) int { return i }

	i := iter.New(list)
	iter.Map[int, int](i, ident).
		ForEach(func(i int) { fmt.Println(i) })
	// Output:
	// 1
	// 2
	// 3
	// 4
}

func ExampleMapped_Nth() {
	list := []int{-1, -2, 3, 4}
	ident := func(i int) int { return i }

	i := iter.New(list)
	m := iter.Map[int, int](i, ident).Nth(2)

	fmt.Println(*m)
	// Output:
	// 3
}

func ExampleMapped_All() {
	gt0 := func(a int) bool { return a > 0 }
	gt2 := func(a int) bool { return a > 2 }
	ident := func(i int) int { return i }
	list := []int{1, 2, 3}

	i := iter.New(list)
	t := iter.Map[int, int](i, ident).All(gt0)
	fmt.Println(t)

	i = iter.New(list)
	m := iter.Map[int, int](i, ident)
	f := m.All(gt2)
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

func ExampleMapped_Any() {
	gt0 := func(a int) bool { return a > 0 }
	ne2 := func(a int) bool { return a != 2 }
	ident := func(i int) int { return i }
	list := []int{1, 2, 3}

	i := iter.New(list)
	t := iter.Map[int, int](i, ident).Any(gt0)
	fmt.Println(t)

	i = iter.New(list)
	m := iter.Map[int, int](i, ident)
	f := m.Any(ne2)
	fmt.Println(f)
	// Any stops at the first true, so there are still more elements
	fmt.Println(*i.Next())
	// Output:
	// true
	// true
	// 2
}

func ExampleMapped_Collect() {
	ident := func(i int) int { return i }
	i := iter.New([]int{1, 2, 3})

	s := iter.Map[int](i, ident).Collect()
	fmt.Println(s)
	// Output:
	// [1 2 3]
}

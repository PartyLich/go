package iter_test

import (
	"fmt"

	"github.com/PartyLich/go/iter"
)

func ExampleNew() {
	slice := []int{1, 2, 3}
	i := iter.New(slice)

	fmt.Println(*i.Next())
	fmt.Println(*i.Next())
	fmt.Println(*i.Next())
	fmt.Println(i.Next())

	// Output:
	// 1
	// 2
	// 3
	// <nil>
}

func ExampleIterable_Next() {
	slice := []int{1, 2, 3}
	i := iter.New(slice)

	fmt.Println(*i.Next())
	fmt.Println(*i.Next())
	fmt.Println(*i.Next())
	fmt.Println(i.Next())

	// Output:
	// 1
	// 2
	// 3
	// <nil>
}

func ExampleIterable_Find() {
	list := []int{1, 2, 3, 4}
	isTwo := func(i int) bool { return i == 2 }
	i := iter.New(list)

	fmt.Println(*i.Find(isTwo))
	fmt.Println(i.Find(isTwo))

	// Output:
	// 2
	// <nil>
}

func ExampleIterator_Rev() {
	slice := []int{1, 2, 3}
	i := iter.New(slice).Rev()

	fmt.Println(*i.Next())
	fmt.Println(*i.Next())
	fmt.Println(*i.Next())
	fmt.Println(i.Next())

	// Output:
	// 3
	// 2
	// 1
	// <nil>
}

func ExampleMap() {
	list := []int{1, 2, 3, 4}
	i := iter.New(list)

	doubled := iter.Map[int, int](i, func(n int) int { return n * 2 })

	for val := doubled.Next(); val != nil; val = doubled.Next() {
		fmt.Println(*val)
	}

	// Output:
	// 2
	// 4
	// 6
	// 8
}

func Example_compose() {
	list := []int{1, 2, 3, 4}
	i := iter.New(list)

	doubled := iter.Map[int, int](i, func(n int) int { return n * 2 })
	f := iter.Filter[int](doubled, func(n int) bool { return n > 5 })

	for val := f.Next(); val != nil; val = f.Next() {
		fmt.Println(*val)
	}

	// Output:
	// 6
	// 8
}

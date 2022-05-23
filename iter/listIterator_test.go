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

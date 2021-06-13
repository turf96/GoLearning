package main

import (
	"container/list"
	"fmt"
)


func main()  {
	l := list.New()

	//insert elements from list
	//1.insert elements from back
	l.PushBack("canon")

	//2.insert elements from front
	l.PushFront(67)

	//traverse the list
	for i := l.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
	fmt.Println("**************")
	//3.insert a new element from back and save it
	element := l.PushBack("first")

	//4.insert a new element after "first"
	l.InsertAfter("high", element)

	//5.insert a new element before "first"
	l.InsertBefore("noon", element)

	for i := l.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}

	fmt.Println("**************")
	//remove element from the list
	l.Remove(element)

	for i := l.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
}


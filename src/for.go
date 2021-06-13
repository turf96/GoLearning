package main

import (
	"fmt"
)

func traverseSlice()  {
	for key, value := range []int{1,2,5,3,6} {
		fmt.Printf("key: %d, value: %d\n", key, value)
	}
}

func traverseString()  {
	var str = "helloWorld"
	for key, value := range str{
		fmt.Printf("key: %d, value: %c\n", key, value)
	}
}

func traverseMap()  {
	m := map[string]int{
		"hello": 100,
		"world": 200,
	}

	for key, value := range m{
		fmt.Printf("key: %s, value: %d\n", key, value)
	}
}

func traverseChannel()  {
	c := make(chan int)

	go func() {
		c <- 1
		c <- 2
		c <- 3
		close(c)
	}()
	for v := range c{
		fmt.Println(v)
	}
}

func main()  {
	traverseSlice()
	traverseString()
	traverseMap()
	traverseChannel()
}

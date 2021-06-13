package main

import (
	"fmt"
	"sync"
)

func main()  {
	var scene sync.Map

	// save key-value to sync.Map
	scene.Store("Monday", 1)
	scene.Store("Tuesday", 2)
	scene.Store("Wednesday", 3)

	//get value with key from sync.Map
	fmt.Println(scene.Load("Monday"))

	//delete key-value from sync.Map
	scene.Delete("Monday")

	//traverse all the key-value from sync.Map
	scene.Range(func(key, value interface{}) bool {
		fmt.Println("iterate:", key, value)
		return true
	})

}
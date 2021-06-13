package main

import "fmt"

type Profile struct {
	Name 	string
	Age 	int
	Married	bool
}

type queryKey struct {
	Name	string
	Age 	int
}

var mapper = make(map[queryKey]*Profile)

func buildIndex(list []*Profile)  {
	//traverse all
	for _,profile := range list{
		//construct query key
		key := queryKey{
			Name: profile.Name,
			Age: profile.Age,
		}
		//save key
		mapper[key] = profile
	}
}

func queryData(name string, age int)  {
	//construct key based on query condition
	key := queryKey{name, age}

	//query data base on key
	result, ok := mapper[key]

	//print the result
	if ok {
		fmt.Println(result)
	} else {
		fmt.Println("not found")
	}
}

func main()  {
	list := []*Profile {
		{Name: "张三丰", Age: 99, Married: false},
		{Name: "黄飞鸿", Age: 30, Married: true},
		{Name: "杨过", Age: 20, Married: false},
	}

	buildIndex(list)
	queryData("张三丰", 99)
}
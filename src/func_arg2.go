package main

//Go语言中传入与返回参数在调用和返回时都使用值传递，
//这里需要注意的是指针、切片和 map 等引用型对象在参数传递中不会发生复制，
//而是将指针进行复制，类似于创建一次引用。

import "fmt"

type Data struct {
	complax []int       //测试切片在参数传递的效果
	instance InnerData  //实例分配的 innerData
	ptr *InnerData		//将 ptr 声明为 InnerData 的指针类型
}

type InnerData struct {
	a int
}

func main()  {
	in := Data{
		complax: []int{1,2,3},
		instance: InnerData{
			4,
		},
		ptr: &InnerData{1},
	}

	//输出结构成员状况
	fmt.Printf("in value: %+v \n", in)
	//output pointer
	fmt.Printf("in value ptr: %p\n", &in)
	//传入结构体，返回同类型的结构体
	out := passByValue(in)
	//输出结构成员情况
	fmt.Printf("out value %+v \n", out)
	//输出结构成员指针情况
	fmt.Printf("out value指针值： %p\n", &out)
}

func passByValue(inFunc Data) Data  {
	//输出参数成员情况
	fmt.Printf("inFunc value:%+v \n", inFunc)
	//打印inFunc 指针
	fmt.Printf("inFunc value:%p \n", &inFunc)
	return inFunc
}
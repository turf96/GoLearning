package main

import "fmt"

func main() {
	slice := []int{1,2,3,4}
	fmt.Printf("切片的原始值: %v\n", slice)
	fmt.Printf("main中slice的内存地址是: %p\n", &slice)
	fmt.Printf("main中slice指向数据的内存地址是: %p\n", slice)

	testChange(slice)
	fmt.Println("切片被修改为: ", slice)
}

func testChange(slice []int){
	fmt.Printf("testChange中slice的内存地址是：%p\n",&slice)
	fmt.Printf("testChange中slice指向数据的内存地址是：%p\n",slice)
	slice[0] = 500}

//运行结果
//切片的原始值: [1 2 3 4]
//main中slice的内存地址是: 0x140c040
//main中slice指向数据的内存地址是: 0x1410080
//testChange中slice的内存地址是：0x140c070
//testChange中slice指向数据的内存地址是：0x1410080
//切片被修改为:  [500 2 3 4]

//原理：
//通过参数传递，main 中的 slice 变量指向的切片地址值被拷贝，所以 testChange 中 slice 的变量本身地址发生变化，
//而存储的切片地址并没有发生变化
//这就是为什么修改切片会影响原值
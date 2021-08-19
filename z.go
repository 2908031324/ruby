package main

import (
	"fmt"
)

func main() {
	a := []int{76, 77, 78, 79, 80}

	for i:=0; i<=cap(a); i++{
		fmt.Printf("第%d次,值：%v\n", i, a[0:i])
	}

	fmt.Println("-----------------------------")
	b := a[1:2]
	fmt.Printf("%v", b)
	fmt.Printf("%v", cap(a))
	printArr(a)
	/*var numbers = make([]int,3,5)
	printArr(numbers)*/
}

func printData(param []struct{}){
	fmt.Printf("结构体：%v", param)

}

func printArr(param1 []int){
	fmt.Printf("array数组为：%v\n", param1)
}
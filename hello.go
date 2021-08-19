package main

import (
	"fmt"
	"strings"
)

var xxx, yyy = "xxx", "yyy"


//https://www.runoob.com/go/go-structures.html 菜鸟教程

func main() {


	//fmt.Printf("%s\n", "aaa打发阿斯顿发斯蒂芬俺的沙发")
	//
	//a := [...]float64{67.7, 89.8, 21, 78}
	//for i := 0; i < len(a); i++ { //looping from 0 to the length of the array
	//	fmt.Printf("%d th element of a is %.2f\n", i, a[i])
	//}
	//
	//var arr2   = [...] string{"第一个", "第二个", "第N个..."}
	//var curLen = len(arr2)
	//
	//for i :=0; i< curLen; i++{
	//	fmt.Printf("%d the element of a is %s\n", i, arr2[i])
	//}



	//a := [...]float64{67.7, 89.8, 21, 78}
	//sum := float64(0)
	//fmt.Println(sum)
	//for i, v := range a {//range returns both the index and value
	//	fmt.Printf("%d the element of a is %.2f\n", i, v)
	//	sum += v
	//}
	//fmt.Println("\nsum of all elements of a",sum)


	//ary := [...] string{"金牌", "银盘", "铜牌"}
	//for ii := 0; ii< len(ary); ii++{
	//	fmt.Printf("key值：%d, value值：%s\n", ii, ary[ii])
	//}
	//
	//
	//for i, v := range ary{
	//	fmt.Printf("key = %d, value = %s\n", i, v)
	//}
	//
	//for _, vvv := range ary{
	//	fmt.Printf("%s\n", vvv)
	//}


	/*if nn := 10; nn%2 == 0{
		fmt.Println("整除", nn)
	}else{
		fmt.Println("不属于", nn)
	}

	if nu := 9; nu % 3 == 0{
		fmt.Println(nu, "整除了")
	}else{
		fmt.Println(nu, "不能被3整除")
	}


	fmt.Println(xxx, yyy)


	switch param := "BBB"; param {
	case "AAA":
		fmt.Println("等于AAA")
	case "BBB":
		fmt.Println("等于BBB")
	default:
		fmt.Println("都不成立")
	}

	var i int
	for i=1; i<=10; i++{
		if i >= 8 {
			break
		}

		if i%3 == 0 {
			continue
		}

		if i%2 == 0{
			fmt.Println(i, "i被整除了")
		}else{
			fmt.Println(i, "i无法被整除")
		}
	}

	var a int = 10

	LOOP: for a < 15{
		if a%3 == 0{
			a = a+1
			goto LOOP
		}

		fmt.Printf("只为 : %d\n", a)
		a++
	}

	var k int = 2
	KKK: for k < 30{
		if k % 3 == 0{
			fmt.Println("k取3余数为：", k)
			k += 5
			goto KKK
		}else{
			fmt.Println("有余数", k)
			k+=10
		}

	}*/

	//var thisArray = [...][3]string{
	//	{"A", "B", "C"},
	//	{"第一个", "第二个", "第三个"},
	//}

	//for k := 0; k< len(thisArray); k++{
	//	curArr := thisArray[k]
	//	for k2 := 0; k2 < len(curArr); k2++ {
	//		fmt.Println("for得出结果", k2, curArr[k2])
	//	}
	//}

	//for _, val1 := range thisArray{
	//	for key2, val2 := range val1{
	//		fmt.Println("range 结果：", key2, val2)
	//	}
	//
	//
	//}

	//fmt.Println(thisArray)
	//fmt.Printf("thisArray地址为：%p", &thisArray)

	//a := [...]string{"USA", "China", "India", "Germany", "France"}
	//a[0] = "ccc"
	//b := a // a copy of a is assigned to b
	//b[0] = "Singapore"
	//fmt.Println("a is ", a)
	//fmt.Println("b is ", b)



	//var sar     = []int{67, 6, 12, 33, 22, 78}
	//var b []int = sar[0:4]
	//fmt.Println(sar)
	//fmt.Println(b)

	//a := [...]int{76, 77, 78, 79, 80}
	//var b []int = a[2:4] //creates a slice from a[1] to a[3]
	//fmt.Println(b)
	//
	//numa := [3]int{78, 79 ,80}
	//nums1 := numa[:] //creates a slice which contains all elements of the array
	//nums2 := numa[:]
	//fmt.Println("array before change 1",numa)
	//nums1[0] = 100
	//fmt.Println("array after modification to slice nums1", numa)
	//nums2[1] = 101
	//fmt.Println("array after modification to slice nums2", numa)


	//aar := [][]int{
	//	{121, 122, 123, 124, 125},
	//	{221, 222, 223, 224, 225},
	//	{331, 332, 333, 334, 335},
	//	{441, 442, 443, 444, 445},
	//	{551, 552, 553, 554, 555},
	//
	//}
	//
	//for _, val := range aar{
	//	ss1 := val[:]
	//	for k2, v2 := range ss1{
	//		fmt.Println(v2)
	//		ss1[k2]++
	//	}
	//	fmt.Println(ss1)
	//	fmt.Println("-----------------------------------------")
	//}



	//const STR     = "字符串"
	//const WIDTH   = 20
	//const LENGTH  = 30
	//
	//
	//const(
	//	a = iota   //0
	//	b          //1
	//	c          //2
	//	d = "ha"   //独立值，iota += 1
	//	e          //"ha"   iota += 1
	//	q = 100    //iota +=1
	//	g          //100  iota +=1
	//	h          //7,恢复计数
	//	i          //8
	//)
	//fmt.Println(a,b,c,d,e,q,g,h,i)

	//var numbers = make([]int,3,5)
	//
	//fmt.Println(numbers)
	//var aa int = 100
	//fmt.Println(aa)
	//var arr = []int{12}
	//aaa := []int{1,2,3}
	//fmt.Println(arr)
	//fmt.Println(aaa)
	//
	//var privateParam999 int = 999
	//fmt.Println(privateParam999)



	//var str string = "a,b,c"
	//fmt.Println(str)
	//var ss1 = strings.Split(str, ",")
	//fmt.Println(ss1)
	//var toUpper = strings.ToUpper(str)
	//var toLower = strings.ToLower(toUpper)
	//fmt.Println(toUpper)
	//fmt.Println(toLower)
	//
	//var arr = [][]string{
	//	{"1数组-1", "1数组-2"},
	//	{"2数组-1", "2数组-2"},
	//}
	//fmt.Println(arr)
	//
	////var st string = "hello"
	//gl := map[int]int{
	//	99: 99,
	//	97: 97,
	//}
	//fmt.Println(gl)
	//
	//glArr := map[string]string{
	//	"id":   "1",
	//	"name": "张三",
	//	"adr":  "广州市天河区",
	//}
	//fmt.Println(glArr)
	//
	//for k, v := range glArr{
	//	fmt.Println(k, "=>", v)
	//}
	//
	//
	//paramDemo := map[string]string{
	//	"a": "a",
	//}
	//fmt.Println("main.paramDemo 2", paramDemo)


	var abc string = "string"
	var toUpper = strings.ToUpper(abc)
	fmt.Println(toUpper)

}

func init() {

	fmt.Println("初始化方法：func init")
}

type StructDemo struct {
	private string
	PublicParam int32
}

func (demo *StructDemo) privateFun() error {
	fmt.Println("abc")
	return nil
}



func privateFunc(demo *StructDemo) error{
	fmt.Println("私有方法")
	return nil
}

func PublicFunc(demo *StructDemo) error{
	return nil
}


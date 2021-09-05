package arrays

import "fmt"

var Test [4]int

func ReverseArray(a []int32) []int32 {
	retArr := make([]int32, len(a))
	for i := 0; i < len(a); i++ {
		retArr[i] = a[len(a)-(i+1)]
	}
	fmt.Println(retArr)
	fmt.Println(Test)
	return retArr
}

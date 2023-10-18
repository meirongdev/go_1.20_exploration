package main

import "fmt"

func main() {
	// arr to slice
	arr1 := [5]int{1, 2, 3, 4, 5}
	s1 := arr1[:]
	fmt.Printf("%v %p\n", arr1, &arr1)
	fmt.Printf("%v %p\n", s1, &s1)

	// slice to array
	s := []int{1, 2, 3, 4, 5}
	fmt.Printf("%v %p\n", s, &s)
	arr := [5]int(s) // Go1.19 cannot convert s (variable of type []int) to type [5]int
	fmt.Printf("%v %p\n", arr, &arr)
	// slice to array pointer
	arrp := (*[5]int)(s)
	fmt.Printf("%v %p\n", arrp, &arrp)
	// Go1.19中，我们可以通过获取底层数组指针的方式，将切片转换为数组
	arr2 := *arrp
	fmt.Printf("%v %p\n", arr2, &arr2)
	s[0] = 100
	fmt.Printf("%v %p\n", s, &s) // [100 2 3 4 5]
	// arr和arr2都不会受到原先的切片s变化的影响
	fmt.Printf("%v %p\n", arr, &arr)   // [1 2 3 4 5]
	fmt.Printf("%v %p\n", arrp, &arrp) // &[100 2 3 4 5]
	fmt.Printf("%v %p\n", arr2, &arr2) // [1 2 3 4 5]
}

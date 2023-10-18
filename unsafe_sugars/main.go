package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println("unsafe.StringData with []byte")
	var arr = [5]byte{'a', 'b', 'c', 'd', 'e'}
	s := unsafe.String(&arr[0], 5)
	fmt.Printf("%v %p\n", s, &s) // abcde 0xc00000e1e0

	sd := unsafe.StringData(s)
	s1 := unsafe.String(sd, len(s))
	fmt.Printf("%v %p\n", s1, &s1) // abcde 0xc00000e200

	fmt.Println("unsafe.StringData with string")
	s2 := "abcde"
	fmt.Printf("%v %p\n", s2, &s2) // abcde 0xc00000e1e0
	sd1 := unsafe.StringData(s1)
	s3 := unsafe.String(sd1, len(s1))
	fmt.Printf("%v %p\n", s3, &s3) // abcde 0xc00000e220

	fmt.Println("unsafe.SliceData")
	sl := arr[:]
	fmt.Printf("%v %p\n", sl, &sl) // [97 98 99 100 101] 0xc00000e210
	r := unsafe.SliceData(sl)
	// Go1.17已经引入了unsafe.Slice，可以直接将底层存储的指针转为切片
	sl2 := unsafe.Slice(r, len(arr))
	fmt.Printf("%v %p\n", sl2, &sl2) // [97 98 99 100 101] 0xc00000e220
}

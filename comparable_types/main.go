package main

import "fmt"

func keys[K comparable, V any](m map[K]V) []K {
	var keys []K
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

type item struct {
	name  string
	price int
}

func main() {
	m := make(map[any]any) // ok
	// m:=make(map[interface{}]interface{}) // 这里any和interface{}是等价的
	// compiler error (Go 1.19): any does not implement comparable
	fmt.Printf("%v %d\n", keys(m), len(m)) // [] 0
	m[1] = 1
	m["1"] = "1"
	m[true] = true
	// compiler error (Go 1.19): any does not implement comparable
	fmt.Printf("%v %d\n", keys(m), len(m)) // [1 1 true] 3
	// func() {} 不是 comparable，但它是any，运行时会报错 hash of unhashable type func()
	// k := func() {}
	// m[k] = 1 // panic: runtime error: hash of unhashable type func()

	// struct as key
	item1 := item{"item1", 1}
	item2 := item{"item2", 2}
	m[item1] = 1
	m[item2] = 2
	fmt.Printf("%v %d\n", keys(m), len(m)) // [1 1 true {item1 1} {item2 2}] 5
	item3 := item{"item1", 1}
	m[item3] = 3
	fmt.Printf("%v %d\n", keys(m), len(m)) // [1 1 true {item1 1} {item2 2}] 5
	item4 := item{price: 1, name: "item1"}
	m[item4] = 4
	fmt.Printf("%v %d\n", keys(m), len(m)) // [1 true {item1 1} {item2 2} 1] 5

	// array as key
	// 相同长度的数组，元素类型相同，元素值相同，那么它们的哈希值也是相同的
	arr1 := [5]int{1, 2, 3, 4, 5}
	arr2 := [5]int{1, 2, 3, 4, 5}
	m[arr1] = 1
	m[arr2] = 2
	fmt.Printf("%v %d\n", keys(m), len(m)) // [{item2 2} [1 2 3 4 5] 1 1 true {item1 1}] 6
	// 相同长度，不同顺序
	arr3 := [5]int{1, 2, 3, 5, 4}
	m[arr3] = 3
	fmt.Printf("%v %d\n", keys(m), len(m)) // [{item1 1} {item2 2} [1 2 3 4 5] [1 2 3 5 4] 1 1 true] 7
	// slice as key
	// s1 := []int{1, 2, 3, 4, 5}
	// s2 := []int{1, 2, 3, 4, 5}
	// // panic: runtime error: hash of unhashable type []int
	// m[s1] = 1
	// m[s2] = 2
	// fmt.Printf("%v %d\n", keys(m), len(m))
}

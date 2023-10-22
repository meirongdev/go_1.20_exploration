package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	str := "Hello, World!"
	firstPrefix := "Hello,"
	secondPrefix := "World!"
	thirdPrefix := ""

	fmt.Println(strings.CutPrefix(str, firstPrefix))
	fmt.Println(strings.CutPrefix(str, secondPrefix))
	fmt.Println(strings.CutPrefix(str, thirdPrefix))
	// 输出
	//  World! true
	// Hello, World! false
	// Hello, World! true

	suffix1 := "World!"
	suffix2 := "Hello,"
	suffix3 := ""
	fmt.Println(strings.CutSuffix(str, suffix1))
	fmt.Println(strings.CutSuffix(str, suffix2))
	fmt.Println(strings.CutSuffix(str, suffix3))
	// 输出
	//  Hello,  true
	// Hello, World! false
	// Hello, World! true

	// Clone for bytes, 对clone的slice修改不会影响原slice
	b := []byte("Hello, World!")
	fmt.Printf("b: %s %p\n", b, b)
	bCloned := bytes.Clone(b)
	fmt.Printf("bCloned: %s %p\n", bCloned, bCloned)
	// 输出
	// b: Hello, World! 0xc000184000
	// bCloned: Hello, World! 0xc000184010
	bCloned[0] = 'h'
	fmt.Printf("b: %s %p\n", b, b)
	fmt.Printf("bCloned: %s %p\n", bCloned, bCloned)
	// 输出
	// b: Hello, World! 0xc000184000
	// bCloned: hello, World! 0xc000184010
}

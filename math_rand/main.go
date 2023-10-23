package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// Go1.20之后，每次执行都会输出不同的随机数
	// Go1.20之前，每次执行都会输出相同的随机数
	// rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Int())
}

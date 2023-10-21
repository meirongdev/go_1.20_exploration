package main

import (
	"errors"
	"fmt"
)

func main() {
	err0 := errors.New("err0")
	wErr0 := fmt.Errorf("werr0: %w", err0)
	fmt.Printf("errors.Unwrap before 1.20(fmt.Errorf): %v\n", wErr0)
	// 输出
	// errors.Unwrap before 1.20(fmt.Errorf): werr0: err0

	// 创建两个error
	err1 := errors.New("err1")
	err2 := errors.New("err2")

	// Join 函数将多个错误连接起
	err3 := errors.Join(err1, err2)
	fmt.Printf("output err3:\n%v\n", err3)
	// 输出
	// err1
	// err2
	if errors.Is(err3, err1) && errors.Is(err3, err2) {
		fmt.Println("err3 contains err1 and err2")
	}
	fmt.Printf("errors.Unwrap err3(Join):%v\n", errors.Unwrap(err3))
	// 输出
	// errors.Unwrap err3(Join):<nil>

	err4 := fmt.Errorf("err4 wrap err3\n%w", err3)
	fmt.Printf("output err4:\n%v\n", err4)
	// 输出
	// err4 wrap err3
	// err1
	// err2
	if errors.Is(err4, err1) && errors.Is(err4, err2) {
		fmt.Println("err4 contains err1 and err2")
	}
	if errors.Is(err4, err3) {
		fmt.Println("err4 contains err3")
	}
	fmt.Printf("errors.Unwrap err4(fmt.Errorf single %cw):%v\n", '%', errors.Unwrap(err4))
	// 输出
	// errors.Unwrap err4(fmt.Errorf single %w):err1

	err5 := errors.New("err5")
	err6 := errors.Join(err4, err5)
	fmt.Printf("output err6:\n%v\n", err6)
	// 输出
	// err4 wrap err3
	// err1
	// err2
	// err5
	if errors.Is(err6, err1) && errors.Is(err6, err2) {
		fmt.Println("err6 contains err1 and err2")
	}
	if errors.Is(err6, err3) {
		fmt.Println("err6 contains err3")
	}
	if errors.Is(err6, err4) {
		fmt.Println("err6 contains err4")
	}
	if errors.Is(err6, err5) {
		fmt.Println("err6 contains err5")
	}

	// fmt.Errorf multiple errors
	err7 := fmt.Errorf("%w\n%w", err4, err5)
	fmt.Printf("output err7:\n%v\n", err7)

	fmt.Printf("errors.Unwrap err7(fmt.Errorf multiple %cw):%v\n", '%', errors.Unwrap(err7))
	// 输出
	// errors.Unwrap err7(fmt.Errorf multiple %w):<nil>
}

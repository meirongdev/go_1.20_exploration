package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println("Without Format - ", now)
	// 输出
	// 2023-10-22 22:21:19.870200423 +0800 +08 m=+0.000034223
	formatted := now.Format("2006-01-02 15:04:05")
	fmt.Println("Format 2006-01-02 15:04:05 -", formatted)
	// 输出
	// 2023-10-22 22:21:19

	// 2006-01-02 15:04:05 是固定的格式，不可更改，如果更改，会出现奇怪的结果
	formatted = now.Format("2007-01-02 15:04:05")
	fmt.Println("Format 2007-01-02 15:04:05 -", formatted)
	// 输出
	// 22007-10-22  22:21:19

	formatted = now.Format(time.DateTime)
	fmt.Println(formatted)
	// 输出
	// 2023-10-22 22:21:19

	formatted = now.Format(time.DateOnly)
	fmt.Println(formatted)
	// 输出
	// 2023-10-22

	formatted = now.Format(time.TimeOnly)
	fmt.Println(formatted)
	// 输出
	// 22:21:19
}

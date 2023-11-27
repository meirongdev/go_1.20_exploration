package main

import (
	"fmt"
	"time"
)

func main() {
	// time.Parse忽略sub-nanosecond部分
	const in = "2021-09-29T16:04:33.0000000000Z"
	fmt.Println(time.Parse(time.RFC3339, in))
	fmt.Println(time.Parse(time.RFC3339Nano, in))

	// Go1.19.13
	// 2021-09-29 16:04:33 +0000 UTC <nil>
	// 0001-01-01 00:00:00 +0000 UTC parsing time "2021-09-29T16:04:33.0000000000Z" as "2006-01-02T15:04:05.999999999Z07:00": cannot parse "0Z" as "Z07:00"
	// d

	// Go1.20
	// 2021-09-29 16:04:33 +0000 UTC <nil>
	// 2021-09-29 16:04:33 +0000 UTC <nil>

	// time.Parse对RFC3339校验更多严格，不允许出现多余的符号
	ts, err := time.Parse(time.RFC3339, "2020-04-14T08:13:26--4:00")
	fmt.Printf("%v , %v\n", ts, err)

	// Go1.19
	// 2020-04-14 08:13:26 +0400 +0400 , <nil>

	// Go1.20
	// 0001-01-01 00:00:00 +0000 UTC , parsing time "2020-04-14T08:13:26--4:00" as "2006-01-02T15:04:05Z07:00": cannot parse "--4:00" as "Z07:00"
}

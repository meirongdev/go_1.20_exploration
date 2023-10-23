package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func main() {
	withTimeoutDemo()
	withCancelCauseDemo()
}

func withTimeoutDemo() {
	// 传递一个带Timeout的context，在ctx.Done()阻塞后，超时抛出异常
	timeoutDuration := 3 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDuration)
	defer cancel()

	<-ctx.Done()

	switch ctx.Err() {
	case context.DeadlineExceeded:
		fmt.Println("context timeout exceeded")
	case context.Canceled:
		fmt.Println("context cancelled by force")
	}

	// 输出:
	// context timeout exceeded
}

var ErrTemporarilyUnavailable = fmt.Errorf("service temporarily unavailable")

func withCancelCauseDemo() {
	ctx, cancel := context.WithCancelCause(context.Background())
	// 主动取消服务，并传递取消原因
	cancel(ErrTemporarilyUnavailable)

	// error的类型是*cancelError
	switch ctx.Err() {
	case context.Canceled:
		fmt.Println("context cancelled by force")
	}

	// 获取取消原因
	err := context.Cause(ctx)

	if errors.Is(err, ErrTemporarilyUnavailable) {
		fmt.Printf("cancellation reason: %s\n", err)
	}
	// 输出:
	// cancellation reason: service temporarily unavailable
}

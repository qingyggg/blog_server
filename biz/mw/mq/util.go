package mq

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"time"
)

func failOnError(err error, msg string) {
	if err != nil {
		hlog.Fatalf("mq==>%s: %s", msg, err)
	}
}
func newContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 5*time.Second)
}

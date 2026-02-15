package model

import (
	"fmt"
)

type Context interface {
	Value(key any) any
}

func requestUUID(ctx Context) string {
	return fmt.Sprintf("%v", ctx.Value("requestUUID"))
}

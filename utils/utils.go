package utils

import (
	"context"
)

func GetStringValueFromCtx(ctx context.Context, key string) (value string) {
	// 从 context 中获取值
	if inf := ctx.Value(key); inf != nil {
		value, _ = inf.(string)
	}
	return value
}

func StringValue(v *string) string {
	if v == nil {
		return ""
	}
	return *v
}

package hook

import (
	"context"
)

type Hook interface {
	Init(params map[string]string) error
	VerifyAPIKey(key string) (string, error)
	Mock(ctx context.Context, req interface{}, fullMethod string) (bool, interface{}, error)
	Before(ctx context.Context, req interface{}, fullMethod string) (context.Context, error)
	After(ctx context.Context, result interface{}, err error, fullMethod string) error
	Release()
}

type LogInfo interface {
	GetFields(fields []string) map[string]any
}

type Extension interface {
	Report(info any) int
	ReportRefused(ctx context.Context, req interface{}, resp interface{}, err error, fullMethod string) error
}

type HookContextKeyType string

const GinParamsKey = HookContextKeyType("gin_params")

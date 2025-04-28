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

type Cipher interface {
	Init(params map[string]string) error

	GetEncryptor(ezID, collectionID int64) (encryptor Encryptor, safeKey []byte, err error)
	GetDecryptor(ezID, collectionID int64, safeKey []byte) (Decryptor, error)
	GetUnsafeKey(ezID, collectionID int64) []byte
}

// Encryptor uses one DEK in the life cycle
type Encryptor interface {
	Encrypt(plainText []byte) (cipherText []byte, err error)
}

// Decryptor uses one DEK in the life cycle
type Decryptor interface {
	Decrypt(cipherText []byte) (plainText []byte, err error)
}

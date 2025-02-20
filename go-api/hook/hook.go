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

// Requirements
// Design V1:
//  1. Memory Issues
//  2. Lack of Flexibility for Plugin Users
//  3. Associate cannot be shared.
type Associate struct {
	ID        int64
	Verifier  []byte
	UnsafeKey []byte
}
type CipherV1 interface {
	Init(params map[string]string) error
	Encrypt(data []byte, associate Associate) (cipherText []byte, err error)
	Decrypt(data []byte, associate Associate) (plainText []byte, err error)

	// For CPP Tasks
	GetEZK(associate Associate) []byte
}

// Design V2
type CipherV2 interface {
	Init(params map[string]string) error

	GetEncryptor(ID int64) (Encryptor, error)
	GetDecryptor(ID int64, key []byte) (Decryptor, error)
	GetUnsafeKey(ID int64) []byte
}

// Encryptor uses one DEK in the life cycle
type Encryptor interface {
	Encrypt(plainText []byte) (cipherText []byte, err error)
	GetSafeKey() (key []byte, err error)
}

// Decryptor uses one DEK in the life cycle
type Decryptor interface {
	Decrypt(cipherText []byte) (plainText []byte, err error)
}

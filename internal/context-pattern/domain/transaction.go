package domain

import "context"

type TxFunction func(ctx context.Context) error

type TxManager interface {
	DoInTx(ctx context.Context, fn TxFunction) error
}

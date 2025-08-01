package domain

import "context"

type Tx any

type TxFunction func(ctx context.Context, tx Tx) error

type TxManager interface {
	DoInTx(ctx context.Context, fn TxFunction) error
}

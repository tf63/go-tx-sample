package domain

import "context"

type AccountRepository interface {
	FindByID(ctx context.Context, id string) (*Account, error)
	FindByIDWithTx(ctx context.Context, id string, tx Tx) (*Account, error)
	Save(ctx context.Context, account Account) error
	SaveWithTx(ctx context.Context, account Account, tx Tx) error
}

package domain

import "context"

type AccountRepository interface {
	FindByID(ctx context.Context, id string) (*Account, error)
	Save(ctx context.Context, account Account) error
}

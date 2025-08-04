package domain

import "context"

type RepositoryManager interface {
	AccountRepository() AccountRepository
}

type UnitOfWork interface {
	DoInTx(
		ctx context.Context,
		fn func(ctx context.Context, rpManager RepositoryManager) error,
	) error
}

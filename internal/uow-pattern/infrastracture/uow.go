package infrastracture

import (
	"context"
	"database/sql"
)

// UnitOfWorkの実装
type UnitOfWork interface {
	DoInTx(
		ctx context.Context,
		fn func(ctx context.Context, uowRepoManager RepositoryManager) error,
	) error
}

type unitOfWork struct {
	db *sql.DB
}

func NewUnitOfWork(db *sql.DB) UnitOfWork {
	return &unitOfWork{db: db}
}

func (u *unitOfWork) DoInTx(
	ctx context.Context,
	fn func(ctx context.Context, uowRepoManager RepositoryManager) error,
) error {
	tx, err := u.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	repoManager := NewRepositoryManager(
		NewAccountRepository(tx),
	)

	if err = fn(ctx, repoManager); err != nil {
		if txErr := tx.Rollback(); txErr != nil {
			return err
		}
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}
	return nil
}

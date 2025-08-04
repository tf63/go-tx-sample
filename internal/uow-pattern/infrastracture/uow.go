package infrastracture

import (
	"context"
	"database/sql"

	"github.com/tf63/go-tx-sample/internal/uow-pattern/domain"
)

// UnitOfWorkの実装

type unitOfWork struct {
	db *sql.DB
}

func NewUnitOfWork(db *sql.DB) domain.UnitOfWork {
	return &unitOfWork{db: db}
}

func (u *unitOfWork) DoInTx(
	ctx context.Context,
	fn func(ctx context.Context, rpManager domain.RepositoryManager) error,
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

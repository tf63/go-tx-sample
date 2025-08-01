package infrastracture

import (
	"context"
	"database/sql"

	"github.com/tf63/go-tx-sample/internal/context-pattern/db/xcontext"
	"github.com/tf63/go-tx-sample/internal/context-pattern/domain"
)

type txManager struct {
	db *sql.DB
}

func NewTransactionManager(db *sql.DB) domain.TxManager {
	return &txManager{
		db,
	}
}

func (tm *txManager) DoInTx(
	ctx context.Context,
	fn domain.TxFunction,
) error {
	tx, err := tm.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	ctxWithTx := xcontext.WithTx(ctx, tx)

	err = fn(ctxWithTx)
	if err != nil {
		return err
	}

	return nil
}

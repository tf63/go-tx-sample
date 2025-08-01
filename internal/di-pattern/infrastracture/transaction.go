package infrastracture

import (
	"context"
	"database/sql"
	"errors"

	"github.com/tf63/go-tx-sample/internal/di-pattern/domain"
)

type Tx interface {
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
}

type tx struct {
	*sql.Tx
}

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
	_tx, err := tm.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			_tx.Rollback()
		} else {
			err = _tx.Commit()
		}
	}()

	err = fn(ctx, &tx{_tx})
	if err != nil {
		return err
	}

	return nil
}

func ExtractTx(_tx domain.Tx) (Tx, error) {
	tx, ok := _tx.(*tx)
	if !ok {
		return nil, errors.New("mysql Tx is invalid")
	}
	return tx, nil
}

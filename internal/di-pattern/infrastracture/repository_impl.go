package infrastracture

import (
	"context"
	"database/sql"

	"github.com/tf63/go-tx-sample/internal/di-pattern/domain"
)

type accountRepositoryImpl struct {
	DB *sql.DB
}

func NewAccountRepository(db *sql.DB) domain.AccountRepository {
	return &accountRepositoryImpl{DB: db}
}

func (r *accountRepositoryImpl) FindByID(ctx context.Context, id string) (*domain.Account, error) {
	// (IDで口座を検索する)
	return &domain.Account{}, nil
}

func (r *accountRepositoryImpl) FindByIDWithTx(
	ctx context.Context,
	id string,
	_tx domain.Tx,
) (*domain.Account, error) {
	tx, err := ExtractTx(_tx)
	if err != nil {
		return nil, err
	}

	tx.QueryRowContext(
		ctx,
		"SELECT id, balance FROM accounts WHERE id = ?",
		id,
	)

	// (トランザクションを使用してIDで口座を検索する)
	return &domain.Account{}, nil
}

func (r *accountRepositoryImpl) Save(ctx context.Context, account domain.Account) error {
	// (永続化する)
	return nil
}

func (r *accountRepositoryImpl) SaveWithTx(
	ctx context.Context,
	account domain.Account,
	_tx domain.Tx,
) error {
	tx, err := ExtractTx(_tx)
	if err != nil {
		return err
	}

	// ここでトランザクションを使用して永続化する
	tx.ExecContext(
		ctx,
		"INSERT INTO accounts (id, balance) VALUES (?, ?)",
		account.ID,
		account.Balance,
	)

	// (トランザクションを使用して永続化する)
	return nil
}

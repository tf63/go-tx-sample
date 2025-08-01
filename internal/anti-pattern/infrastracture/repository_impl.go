package infrastracture

import (
	"context"
	"database/sql"

	"github.com/tf63/go-tx-sample/internal/anti-pattern/domain"
)

type accountRepositoryImpl struct {
	DB *sql.DB
}

func NewAccountRepository(db *sql.DB) domain.AccountRepository {
	return &accountRepositoryImpl{DB: db}
}

func (r *accountRepositoryImpl) Save(ctx context.Context, account domain.Account) error {
	// (永続化する)
	return nil
}

func (r *accountRepositoryImpl) FindByID(ctx context.Context, id string) (*domain.Account, error) {
	// (IDで口座を検索する)
	return &domain.Account{}, nil
}

// アンチパターン: Repositoryにトランザクションを含める
func (a *accountRepositoryImpl) Transfer(
	ctx context.Context,
	fromID, toID string,
	amount int,
) error {
	tx, err := a.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		} else if err != nil {
			tx.Rollback() // エラーがあればロールバック
		} else {
			err = tx.Commit() // 成功すればコミット
		}
	}()

	fromAccount, err := a.FindByID(ctx, fromID)
	if err != nil {
		return err
	}

	toAccount, err := a.FindByID(ctx, toID)
	if err != nil {
		return err
	}

	if err := fromAccount.Withdraw(amount); err != nil {
		return err
	}

	toAccount.Deposit(amount)

	if err := a.Save(ctx, *fromAccount); err != nil {
		return err
	}

	if err := a.Save(ctx, *toAccount); err != nil {
		return err
	}

	return nil
}

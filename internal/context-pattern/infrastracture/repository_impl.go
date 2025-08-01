package infrastracture

import (
	"context"
	"database/sql"

	"github.com/tf63/go-tx-sample/internal/context-pattern/domain"
)

type accountRepositoryImpl struct {
	BaseRepository
}

func NewAccountRepository(db *sql.DB) domain.AccountRepository {
	return &accountRepositoryImpl{
		BaseRepository: NewBaseRepository(db),
	}
}

func (r *accountRepositoryImpl) FindByID(ctx context.Context, id string) (*domain.Account, error) {
	// (IDで口座を検索する)
	return &domain.Account{}, nil
}

func (r *accountRepositoryImpl) Save(ctx context.Context, account domain.Account) error {
	// (永続化する)
	return nil
}

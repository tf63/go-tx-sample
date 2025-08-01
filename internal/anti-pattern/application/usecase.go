package application

import (
	"context"

	"github.com/tf63/go-tx-sample/internal/anti-pattern/domain"
)

type AccountUsecase struct {
	ar domain.AccountRepository
}

type AccountUsecaseInterface interface {
	Transfer(ctx context.Context, fromID, toID string, amount int) error
}

func NewAccountUsecase(ar domain.AccountRepository) *AccountUsecase {
	return &AccountUsecase{ar: ar}
}

func (a *AccountUsecase) Transfer(ctx context.Context, fromID, toID string, amount int) error {
	// アンチパターン: ビジネスロジックがサービス層ではなくリポジトリ層に含まれている
	return a.ar.Transfer(ctx, fromID, toID, amount)
}

package application

import (
	"context"

	"github.com/tf63/go-tx-sample/internal/uow-pattern/domain"
)

type AccountUsecase struct {
	ar  domain.AccountRepository
	uow domain.UnitOfWork
}

type AccountUsecaseInterface interface {
	Transfer(ctx context.Context, fromID, toID string, amount int) error
}

func NewAccountUsecase(
	ar domain.AccountRepository,
	uow domain.UnitOfWork,
) *AccountUsecase {
	return &AccountUsecase{ar: ar, uow: uow}
}

func (a *AccountUsecase) Transfer(ctx context.Context, fromID, toID string, amount int) error {
	return a.uow.DoInTx(
		ctx,
		func(ctx context.Context, rpManager domain.RepositoryManager) error {
			fromAcc, err := rpManager.AccountRepository().FindByID(ctx, fromID)
			if err != nil {
				return err
			}
			toAcc, err := rpManager.AccountRepository().FindByID(ctx, toID)
			if err != nil {
				return err
			}

			if err = fromAcc.Withdraw(amount); err != nil {
				return err
			}
			toAcc.Deposit(amount)

			if err = rpManager.AccountRepository().Save(ctx, *fromAcc); err != nil {
				return err
			}
			if err = rpManager.AccountRepository().Save(ctx, *toAcc); err != nil {
				return err
			}
			return nil
		},
	)
}

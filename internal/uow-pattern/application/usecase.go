package application

import (
	"context"

	"github.com/tf63/go-tx-sample/internal/context-pattern/domain"
)

type AccountUsecase struct {
	ar        domain.AccountRepository
	txManager domain.TxManager
}

type AccountUsecaseInterface interface {
	Transfer(ctx context.Context, fromID, toID string, amount int) error
}

func NewAccountUsecase(
	ar domain.AccountRepository,
	txManager domain.TxManager,
) *AccountUsecase {
	return &AccountUsecase{ar: ar, txManager: txManager}
}

func (a *AccountUsecase) Transfer(ctx context.Context, fromID, toID string, amount int) error {
	return a.txManager.DoInTx(ctx, func(ctx context.Context) error {
		fromAcc, err := a.ar.FindByID(ctx, fromID)
		if err != nil {
			return err
		}
		toAcc, err := a.ar.FindByID(ctx, toID)
		if err != nil {
			return err
		}

		if err = fromAcc.Withdraw(amount); err != nil {
			return err
		}
		toAcc.Deposit(amount)

		if err = a.ar.Save(ctx, *fromAcc); err != nil {
			return err
		}
		if err = a.ar.Save(ctx, *toAcc); err != nil {
			return err
		}
		return nil
	})
}

package domain

import "errors"

// 口座を表すエンティティ
type Account struct {
	ID      string
	Balance int
}

// 口座に入金する
func (a *Account) Deposit(amount int) {
	a.Balance += amount
}

// 口座から引き出す
func (a *Account) Withdraw(amount int) error {
	if a.Balance < amount {
		return errors.New("残高不足")
	}
	a.Balance -= amount
	return nil
}

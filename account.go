package main

import (
	"fmt"
	"time"
)

type Transaction struct {
	Type      string // "deposit", "withdraw", "transfer_out", "transfer_in"
	Amount    float64
	Timestamp time.Time
	ToID      string // используется только для transfer_out
	FromID    string // используется только для transfer_in
}

type Account struct {
	ID           string
	OwnerName    string
	Balance      float64
	Transactions []Transaction
}

func NewAccount(id, ownerName string) *Account {
	return &Account{
		ID:        id,
		OwnerName: ownerName,
		Balance:   0.0,
	}
}

func (a *Account) Deposit(amount float64) error {
	if amount <= 0 {
		return ErrInvalidAmount
	}
	a.Balance += amount
	a.Transactions = append(a.Transactions, Transaction{
		Type:      "deposit",
		Amount:    amount,
		Timestamp: time.Now(),
	})
	return nil
}

func (a *Account) Withdraw(amount float64) error {
	if amount <= 0 {
		return ErrInvalidAmount
	}
	if a.Balance < amount {
		return ErrInsufficientFunds
	}
	a.Balance -= amount
	a.Transactions = append(a.Transactions, Transaction{
		Type:      "withdraw",
		Amount:    amount,
		Timestamp: time.Now(),
	})
	return nil
}

func (a *Account) Transfer(to *Account, amount float64) error {
	if a == to {
		return ErrSameAccountTransfer
	}
	if amount <= 0 {
		return ErrInvalidAmount
	}
	if a.Balance < amount {
		return ErrInsufficientFunds
	}
	a.Balance -= amount
	to.Balance += amount

	now := time.Now()
	a.Transactions = append(a.Transactions, Transaction{
		Type:      "transfer_out",
		Amount:    amount,
		Timestamp: now,
		ToID:      to.ID,
	})

	to.Transactions = append(to.Transactions, Transaction{
		Type:      "transfer_in",
		Amount:    amount,
		Timestamp: now,
		FromID:    a.ID,
	})

	return nil
}

func (a *Account) GetBalance() float64 {
	return a.Balance
}

func (a *Account) GetStatement() string {
	if len(a.Transactions) == 0 {
		return "Нет транзакций."
	}
	statement := "История транзакций:\n"
	for _, t := range a.Transactions {
		var desc string
		switch t.Type {
		case "deposit":
			desc = fmt.Sprintf("Пополнение: %+v", t.Amount)
		case "withdraw":
			desc = fmt.Sprintf("Снятие: %+v", t.Amount)
		case "transfer_out":
			desc = fmt.Sprintf("Перевод на %s: %+v", t.ToID, t.Amount)
		case "transfer_in":
			desc = fmt.Sprintf("Перевод от %s: %+v", t.FromID, t.Amount)
		}
		statement += fmt.Sprintf("[%s] %s\n", t.Timestamp.Format("2006-01-02 15:04"), desc)
	}
	return statement
}

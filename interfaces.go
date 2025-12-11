package main

// AccountService — основной интерфейс для работы со счетом.
type AccountService interface {
	Deposit(amount float64) error
	Withdraw(amount float64) error
	Transfer(to *Account, amount float64) error
	GetBalance() float64
	GetStatement() string
}

// Storage — интерфейс для работы с хранилищем данных.
type Storage interface {
	SaveAccount(account *Account) error
	LoadAccount(accountID string) (*Account, error)
	GetAllAccounts() ([]*Account, error)
}

package main

import "errors"

var (
	ErrInsufficientFunds    = errors.New("недостаточно средств на счете")
	ErrInvalidAmount        = errors.New("некорректная сумма (отрицательная или нулевая)")
	ErrAccountNotFound      = errors.New("счет не найден")
	ErrSameAccountTransfer  = errors.New("нельзя переводить на тот же счёт")
)
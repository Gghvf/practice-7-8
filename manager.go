package main

import (
	"bufio"
	"fmt"
	"log"
)

type AccountManager struct {
	storage Storage
}

func NewAccountManager(storage Storage) *AccountManager {
	return &AccountManager{storage: storage}
}

func (am *AccountManager) createAccount(scanner *bufio.Scanner) {
	fmt.Print("Введите имя владельца счета: ")
	scanner.Scan()
	ownerName := scanner.Text()

	// Генерируем ID
	id := fmt.Sprintf("acc_%d", len(am.getAllAccountIDs())+1)

	account := NewAccount(id, ownerName)
	err := am.storage.SaveAccount(account)
	if err != nil {
		log.Printf("Ошибка сохранения счета: %v", err)
		return
	}
	fmt.Printf("Счет создан с ID: %s\n", id)
}

func (am *AccountManager) getAllAccountIDs() []string {
	accounts, _ := am.storage.GetAllAccounts()
	ids := make([]string, 0, len(accounts))
	for _, acc := range accounts {
		ids = append(ids, acc.ID)
	}
	return ids
}

func (am *AccountManager) selectAndOperate(scanner *bufio.Scanner, operation func(*Account, Scanner)) {
	fmt.Print("Введите ID счета: ")
	scanner.Scan()
	id := scanner.Text()

	account, err := am.storage.LoadAccount(id)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
		return
	}

	operation(account, scanner)
	am.storage.SaveAccount(account) // Сохраняем изменения
}

func (am *AccountManager) transferBetweenAccounts(scanner *bufio.Scanner) {
	fmt.Print("Введите ID вашего счета: ")
	scanner.Scan()
	fromID := scanner.Text()

	fromAcc, err := am.storage.LoadAccount(fromID)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
		return
	}

	fmt.Print("Введите ID счета получателя: ")
	scanner.Scan()
	toID := scanner.Text()

	if fromID == toID {
		fmt.Println("Ошибка:", ErrSameAccountTransfer.Error())
		return
	}

	toAcc, err := am.storage.LoadAccount(toID)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
		return
	}

	fmt.Print("Введите сумму перевода: ")
	scanner.Scan()
	amount, err := parseFloat(scanner.Text())
	if err != nil {
		fmt.Println("Некорректная сумма.")
		return
	}

	err = fromAcc.Transfer(toAcc, amount)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
		return
	}

	am.storage.SaveAccount(fromAcc)
	am.storage.SaveAccount(toAcc)
	fmt.Println("Перевод успешно выполнен.")
}
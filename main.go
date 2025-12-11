package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	storage := NewInMemoryStorage()
	accountManager := NewAccountManager(storage)

	scanner := bufio.NewScanner(os.Stdin)
	for {
		printMenu()
		scanner.Scan()
		choice := strings.TrimSpace(scanner.Text())

		switch choice {
		case "1":
			accountManager.createAccount(scanner)
		case "2":
			accountManager.selectAndOperate(scanner, func(acc *Account, s Scanner) {
				fmt.Print("Введите сумму для пополнения: ")
				s.Scan()
				amount, err := parseFloat(s.Text())
				if err != nil {
					fmt.Println("Некорректная сумма.")
					return
				}
				err = acc.Deposit(amount)
				if err != nil {
					fmt.Printf("Ошибка: %v\n", err)
				} else {
					fmt.Println("Счет успешно пополнен.")
				}
			})
		case "3":
			accountManager.selectAndOperate(scanner, func(acc *Account, s Scanner) {
				fmt.Print("Введите сумму для снятия: ")
				s.Scan()
				amount, err := parseFloat(s.Text())
				if err != nil {
					fmt.Println("Некорректная сумма.")
					return
				}
				err = acc.Withdraw(amount)
				if err != nil {
					fmt.Printf("Ошибка: %v\n", err)
				} else {
					fmt.Println("Средства успешно сняты.")
				}
			})
		case "4":
			accountManager.transferBetweenAccounts(scanner)
		case "5":
			accountManager.selectAndOperate(scanner, func(acc *Account, s Scanner) {
				balance := acc.GetBalance()
				fmt.Printf("Баланс: %.2f\n", balance)
			})
		case "6":
			accountManager.selectAndOperate(scanner, func(acc *Account, s Scanner) {
				statement := acc.GetStatement()
				fmt.Println("Выписка:")
				fmt.Println(statement)
			})
		case "7":
			fmt.Println("Выход из приложения.")
			return
		default:
			fmt.Println("Некорректный выбор. Попробуйте снова.")
		}
	}
}

func printMenu() {
	fmt.Println("\n--- Меню ---")
	fmt.Println("1. Создать счет")
	fmt.Println("2. Пополнить счет")
	fmt.Println("3. Снять средства")
	fmt.Println("4. Перевести другому счету")
	fmt.Println("5. Просмотреть баланс")
	fmt.Println("6. Получить выписку")
	fmt.Println("7. Выйти")
	fmt.Print("Выберите действие: ")
}

func parseFloat(input string) (float64, error) {
	return strconv.ParseFloat(strings.TrimSpace(input), 64)
}

type Scanner interface {
	Scan() bool
	Text() string
}
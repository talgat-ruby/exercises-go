package problem7

import (
	"fmt"
)

// Структура BankAccount (Банковский счёт)
type BankAccount struct {
	name    string  // Имя владельца счёта
	balance float64 // Баланс счёта
}

// Метод для снятия денег
func (a *BankAccount) Withdraw(amount float64) bool {
	if amount > 0 && amount <= a.balance {
		a.balance -= amount
		fmt.Printf("%s withdrew %.2f. New balance: %.2f\n", a.name, amount, a.balance)
		return true
	}
	fmt.Printf("%s failed to withdraw %.2f. Insufficient funds.\n", a.name, amount)
	return false
}

// Структура FedexAccount (Счёт FedEx)
type FedexAccount struct {
	name     string   // Имя владельца счёта
	packages []string // Срез для хранения отправленных пакетов
}

// Метод для отправки пакетов
func (a *FedexAccount) SendPackage(recipient string) {
	a.packages = append(a.packages, recipient)
	fmt.Printf("%s sent a package to %s\n", a.name, recipient)
}

// Структура KazPostAccount (Счёт Казпочты)
type KazPostAccount struct {
	name     string   // Имя владельца счёта
	balance  float64  // Баланс счёта
	packages []string // Срез для хранения отправленных пакетов
}

// Метод для снятия денег
func (a *KazPostAccount) Withdraw(amount float64) bool {
	if amount > 0 && amount <= a.balance {
		a.balance -= amount
		fmt.Printf("%s withdrew %.2f. New balance: %.2f\n", a.name, amount, a.balance)
		return true
	}
	fmt.Printf("%s failed to withdraw %.2f. Insufficient funds.\n", a.name, amount)
	return false
}

// Метод для отправки пакетов
func (a *KazPostAccount) SendPackage(recipient string) {
	a.packages = append(a.packages, recipient)
	fmt.Printf("%s sent a package to %s\n", a.name, recipient)
}

// Функция для снятия денег с нескольких счётов
func withdrawMoney(amount float64, accounts ...interface{}) {
	for _, account := range accounts {
		switch acc := account.(type) {
		case *BankAccount:
			acc.Withdraw(amount)
		case *KazPostAccount:
			acc.Withdraw(amount)
		default:
			fmt.Println("Unsupported account type for withdrawal")
		}
	}
}

// Функция для отправки пакетов с нескольких счётов
func sendPackagesTo(recipient string, accounts ...interface{}) {
	for _, account := range accounts {
		switch acc := account.(type) {
		case *FedexAccount:
			acc.SendPackage(recipient)
		case *KazPostAccount:
			acc.SendPackage(recipient)
		default:
			fmt.Println("Unsupported account type for sending packages")
		}
	}
}

func main() {
	// Создание экземпляров
	normanOsborne := &BankAccount{
		name:    "Norman Osborne",
		balance: 1_000_000,
	}
	peterParker := &FedexAccount{
		name:     "Peter Parker",
		packages: make([]string, 0),
	}
	auntMay := &KazPostAccount{
		name:     "Aunt May",
		balance:  200,
		packages: make([]string, 0),
	}

	// Снятие денег
	withdrawMoney(10.0, normanOsborne, auntMay) // можно снимать с нескольких счётов

	// Отправка пакетов
	sendPackagesTo("Mary Jane  ", peterParker, auntMay) // можно отправлять с нескольких счётов
}

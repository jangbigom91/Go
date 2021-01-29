package accounts

import "errors"

// Account struct
type Account struct {
	owner   string
	balance int
}

var errNoMoney = errors.New("Can't withdraw")

// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit x amount on your account
// Go에서 method -> 작성 시 struct의 이름인 Account의 맨앞글자의 소문자a를 먼저쓰고 Account라고 선언
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// Balance of your account
// method
func (a Account) Balance() int {
	return a.balance
}

// Withdraw x amount from your account
// Go에는 Exception이 없고 직접 error코드를 직접 써야됨
// error에는 두가지 value가 있음 - error, nil
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil
}

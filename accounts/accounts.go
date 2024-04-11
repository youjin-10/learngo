package accounts

type bankAccount struct {
	owner string
	balance int
}

func NewAccount(owner string) *bankAccount {
	account := bankAccount{owner: owner, balance: 0}
	return &account
}

// value receiver
// receives a copy of the type
func (a bankAccount) Owner() string {
	return a.owner
}

func (a bankAccount) Balance() int {
	return a.balance
}

// pointer receiver
// receives a memory address of the type (an actual value, not a copy)
func (a *bankAccount) Deposit(amount int) {
	a.balance += amount
}

package accounts

type bankAccount struct {
	owner string
	balance int
}

func NewAccount(owner string) *bankAccount {
	account := bankAccount{owner: owner, balance: 0}
	return &account
}
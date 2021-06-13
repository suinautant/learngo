package accounts

// Account struct - nico
type Account struct {
	owner   string
	balance int
}

// NewAccount creates Acccount
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

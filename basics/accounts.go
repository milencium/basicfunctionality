package accounts

//Account struct that creates our structs
type Account struct {
	owner string
	balance int
}

func NewAccount(owner string) *Account{
	account := Account{owner:owner, balance:0}
	return &account //we are returning new object, not copy
}

func (a *Account) Deposit(ammount int)int{
	a.balance += ammount
}

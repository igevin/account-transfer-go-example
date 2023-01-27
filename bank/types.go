package bank

type Account struct {
	Id      int64
	Balance int64
}

type Accountable interface {
	Transfer(to Accountable, amount int64)
}

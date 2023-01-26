package bank

import "testing"

func BenchmarkAccount_Transfer(b *testing.B) {
	x, y, z := newAccount(1), newAccount(2), newAccount(3)
	transferTo := func(from, to *Account) {
		from.Transfer(to, 1)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		go transferTo(x, y)
		go transferTo(y, z)
		go transferTo(z, x)
	}
}

func BenchmarkBank_Transfer(b *testing.B) {
	x, y, z := newAccount(1), newAccount(2), newAccount(3)
	bank := NewBank()
	transferTo := func(bank *Bank, from, to *Account) {
		bank.Transfer(from, to, 1)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		go transferTo(bank, x, y)
		go transferTo(bank, y, z)
		go transferTo(bank, z, x)
	}
}

func BenchmarkBank_TransferAsync(b *testing.B) {
	x, y, z := newAccount(1), newAccount(2), newAccount(3)
	bank := NewBank()
	transferTo := func(bank *Bank, from, to *Account) {
		bank.TransferAsync(from, to, 1)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		go transferTo(bank, x, y)
		go transferTo(bank, y, z)
		go transferTo(bank, z, x)
	}
}

func newAccount(id int64) *Account {
	return &Account{
		Id:      id,
		Balance: 1000,
	}
}

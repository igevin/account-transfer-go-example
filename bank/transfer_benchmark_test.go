package bank

import "testing"

func BenchmarkAccount_Transfer(b *testing.B) {
	x, y, z := newAccountV1(1), newAccountV1(2), newAccountV1(3)
	transferTo := func(from, to *AccountV1) {
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
	x, y, z := newAccountV1(1), newAccountV1(2), newAccountV1(3)
	bank := NewBank()
	transferTo := func(bank *Bank, from, to *AccountV1) {
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
	x, y, z := newAccountV1(1), newAccountV1(2), newAccountV1(3)
	bank := NewBank()
	transferTo := func(bank *Bank, from, to *AccountV1) {
		bank.TransferAsync(from, to, 1)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		go transferTo(bank, x, y)
		go transferTo(bank, y, z)
		go transferTo(bank, z, x)
	}
}

func newAccountV1(id int64) *AccountV1 {
	return &AccountV1{
		Account: Account{
			Id:      id,
			Balance: 1000,
		},
	}
}

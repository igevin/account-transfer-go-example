package bank

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
	"time"
)

//func TestBank_Transfer(t *testing.T) {
//	a := newAccountV1(1)
//	b := newAccountV1(2)
//	c := newAccountV1(3)
//	bank := NewBank()
//
//	//t.Logf("Before -----> a: %d, b: %d, c: %d", a.Balance, b.Balance, c.Balance)
//	wg := sync.WaitGroup{}
//
//	start := time.Now()
//
//	concurrentBankTransfer(bank, c, a, 1, &wg)
//	concurrentBankTransfer(bank, a, b, 1, &wg)
//	concurrentBankTransfer(bank, b, c, 1, &wg)
//
//	wg.Wait()
//	//t.Logf("After -----> a: %d, b: %d, c: %d", a.Balance, b.Balance, c.Balance)
//	t.Logf("time: %v", time.Since(start))
//	assert.Equal(t, defaultBalance, a.Balance)
//	assert.Equal(t, defaultBalance, b.Balance)
//	assert.Equal(t, defaultBalance, c.Balance)
//	bank.Close()
//}

//func TestBank_TransferAsync(t *testing.T) {
//	a := newAccountV1(1)
//	b := newAccountV1(2)
//	c := newAccountV1(3)
//	bank := NewBank()
//
//	//t.Logf("Before -----> a: %d, b: %d, c: %d", a.Balance, b.Balance, c.Balance)
//
//	//start := time.Now()
//
//	concurrentBankTransferAsync(bank, c, a, 1)
//	concurrentBankTransferAsync(bank, a, b, 1)
//	concurrentBankTransferAsync(bank, b, c, 1)
//
//	time.Sleep(time.Millisecond * 2)
//	//t.Logf("After -----> a: %d, b: %d, c: %d", a.Balance, b.Balance, c.Balance)
//	//t.Logf("time: %v", time.Since(start))
//
//	assert.Equal(t, defaultBalance, a.Balance)
//	assert.Equal(t, defaultBalance, b.Balance)
//	assert.Equal(t, defaultBalance, c.Balance)
//	bank.Close()
//}
//

func TestAccountV1Transfer(t *testing.T) {
	a := newAccountV1(1)
	b := newAccountV1(2)
	c := newAccountV1(3)
	testAccountableTransfer(t, a, b, c)
}

func TestAccountV2Transfer(t *testing.T) {
	a := newAccountV2(1)
	b := newAccountV2(2)
	c := newAccountV2(3)
	testAccountableTransfer(t, a, b, c)
}

func testAccountableTransfer(t *testing.T, a, b, c Accountable) {
	wg := sync.WaitGroup{}

	start := time.Now()

	concurrentAccountableTransfer(c, a, 1, &wg)
	concurrentAccountableTransfer(a, b, 1, &wg)
	concurrentAccountableTransfer(b, c, 1, &wg)

	wg.Wait()
	t.Logf("time: %v", time.Since(start))

	assert.Equal(t, defaultBalance, a.GetBalance())
	assert.Equal(t, defaultBalance, b.GetBalance())
	assert.Equal(t, defaultBalance, c.GetBalance())
}

func concurrentAccountableTransfer(from, to Accountable, amount int64, wg *sync.WaitGroup) {
	for i := 0; i < 500; i++ {
		wg.Add(1)
		go func() {
			from.Transfer(to, amount)
			wg.Done()
		}()
	}
}

//func concurrentBankTransfer(bank *Bank, from *AccountV1, to *AccountV1, amount int64, wg *sync.WaitGroup) {
//	for i := 0; i < 500; i++ {
//		wg.Add(1)
//		go func() {
//			bank.Transfer(from, to, amount)
//			wg.Done()
//		}()
//	}
//}
//func concurrentBankTransferAsync(bank *Bank, from *AccountV1, to *AccountV1, amount int64) {
//	for i := 0; i < 500; i++ {
//		go func() {
//			bank.TransferAsync(from, to, amount)
//		}()
//	}
//}

func newAccountV2(id int64) *AccountV2 {
	return &AccountV2{
		Account: Account{
			Id:      id,
			Balance: defaultBalance,
		},
	}
}

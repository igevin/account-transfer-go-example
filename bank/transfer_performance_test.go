package bank

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
	"time"
)

func TestBank_Transfer(t *testing.T) {
	a := newAccountV1(1)
	b := newAccountV1(2)
	c := newAccountV1(3)
	bank := NewBank()

	//t.Logf("Before -----> a: %d, b: %d, c: %d", a.Balance, b.Balance, c.Balance)
	wg := sync.WaitGroup{}

	start := time.Now()

	concurrentBankTransfer(bank, c, a, 1, &wg)
	concurrentBankTransfer(bank, a, b, 1, &wg)
	concurrentBankTransfer(bank, b, c, 1, &wg)

	wg.Wait()
	//t.Logf("After -----> a: %d, b: %d, c: %d", a.Balance, b.Balance, c.Balance)
	t.Logf("time: %v", time.Since(start))
	assert.Equal(t, defaultBalance, a.Balance)
	assert.Equal(t, defaultBalance, b.Balance)
	assert.Equal(t, defaultBalance, c.Balance)
	bank.Close()
}

func TestBank_TransferAsync(t *testing.T) {
	a := newAccountV1(1)
	b := newAccountV1(2)
	c := newAccountV1(3)
	bank := NewBank()

	//t.Logf("Before -----> a: %d, b: %d, c: %d", a.Balance, b.Balance, c.Balance)

	//start := time.Now()

	concurrentBankTransferAsync(bank, c, a, 1)
	concurrentBankTransferAsync(bank, a, b, 1)
	concurrentBankTransferAsync(bank, b, c, 1)

	time.Sleep(time.Millisecond * 2)
	//t.Logf("After -----> a: %d, b: %d, c: %d", a.Balance, b.Balance, c.Balance)
	//t.Logf("time: %v", time.Since(start))

	assert.Equal(t, defaultBalance, a.Balance)
	assert.Equal(t, defaultBalance, b.Balance)
	assert.Equal(t, defaultBalance, c.Balance)
	bank.Close()
}

//func TestAccountTransfer(t *testing.T) {
//	a := newAccountV1(1)
//	b := newAccountV1(2)
//	c := newAccountV1(3)
//	//t.Logf("Before -----> a: %d, b: %d, c: %d", a.Balance, b.Balance, c.Balance)
//	wg := sync.WaitGroup{}
//
//	start := time.Now()
//
//	concurrentAccountTransfer(c, a, 1, &wg)
//	concurrentAccountTransfer(a, b, 1, &wg)
//	concurrentAccountTransfer(b, c, 1, &wg)
//
//	wg.Wait()
//	//t.Logf("After -----> a: %d, b: %d, c: %d", a.Balance, b.Balance, c.Balance)
//	t.Logf("time: %v", time.Since(start))
//
//	assert.Equal(t, defaultBalance, a.Balance)
//	assert.Equal(t, defaultBalance, b.Balance)
//	assert.Equal(t, defaultBalance, c.Balance)
//}

//func concurrentAccountTransfer(from *AccountV1, to *AccountV1, amount int64, wg *sync.WaitGroup) {
//	for i := 0; i < 500; i++ {
//		wg.Add(1)
//		go func() {
//			from.Transfer(to, amount)
//			wg.Done()
//		}()
//	}
//}

func concurrentBankTransfer(bank *Bank, from *AccountV1, to *AccountV1, amount int64, wg *sync.WaitGroup) {
	for i := 0; i < 500; i++ {
		wg.Add(1)
		go func() {
			bank.Transfer(from, to, amount)
			wg.Done()
		}()
	}
}
func concurrentBankTransferAsync(bank *Bank, from *AccountV1, to *AccountV1, amount int64) {
	for i := 0; i < 500; i++ {
		go func() {
			bank.TransferAsync(from, to, amount)
		}()
	}
}

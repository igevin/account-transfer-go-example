package bank

import (
	"sync"
	"testing"
	"time"
)

func TestBank(t *testing.T) {
	a := account{Id: 1, Balance: 1000, lock: sync.Mutex{}}
	b := account{Id: 2, Balance: 1000, lock: sync.Mutex{}}
	c := account{Id: 3, Balance: 1000, lock: sync.Mutex{}}
	bank := Bank{lock: sync.Mutex{}, trans: make(chan TransferTask)}

	t.Logf("Before -----> a: %d, b: %d, c: %d", a.Balance, b.Balance, c.Balance)
	wg := sync.WaitGroup{}

	start := time.Now()

	concurrentBankTransfer(&bank, &c, &a, 1, &wg)
	concurrentBankTransfer(&bank, &a, &b, 1, &wg)
	concurrentBankTransfer(&bank, &b, &c, 1, &wg)

	wg.Wait()
	t.Logf("After -----> a: %d, b: %d, c: %d", a.Balance, b.Balance, c.Balance)
	t.Logf("time: %v", time.Since(start))
}

func TestBank2(t *testing.T) {
	a := account{Id: 1, Balance: 1000, lock: sync.Mutex{}}
	b := account{Id: 2, Balance: 1000, lock: sync.Mutex{}}
	c := account{Id: 3, Balance: 1000, lock: sync.Mutex{}}
	bank := Bank{lock: sync.Mutex{}, trans: make(chan TransferTask)}
	bank.TransferHandler()

	t.Logf("Before -----> a: %d, b: %d, c: %d", a.Balance, b.Balance, c.Balance)

	start := time.Now()

	concurrentBankTransfer2(&bank, &c, &a, 1)
	concurrentBankTransfer2(&bank, &a, &b, 1)
	concurrentBankTransfer2(&bank, &b, &c, 1)

	time.Sleep(time.Millisecond * 2)
	t.Logf("After -----> a: %d, b: %d, c: %d", a.Balance, b.Balance, c.Balance)
	t.Logf("time: %v", time.Since(start))
}

func TestAccountTransfer(t *testing.T) {
	a := account{Id: 1, Balance: 1000, lock: sync.Mutex{}}
	b := account{Id: 2, Balance: 1000, lock: sync.Mutex{}}
	c := account{Id: 3, Balance: 1000, lock: sync.Mutex{}}
	t.Logf("Before -----> a: %d, b: %d, c: %d", a.Balance, b.Balance, c.Balance)
	wg := sync.WaitGroup{}

	start := time.Now()

	concurrentAccountTransfer(&c, &a, 1, &wg)
	concurrentAccountTransfer(&a, &b, 1, &wg)
	concurrentAccountTransfer(&b, &c, 1, &wg)

	wg.Wait()
	t.Logf("After -----> a: %d, b: %d, c: %d", a.Balance, b.Balance, c.Balance)
	t.Logf("time: %v", time.Since(start))
}

func concurrentAccountTransfer(from *account, to *account, amount int64, wg *sync.WaitGroup) {
	for i := 0; i < 500; i++ {
		wg.Add(1)
		go func() {
			from.Transfer(to, amount)
			wg.Done()
		}()
	}
}

func concurrentBankTransfer(bank *Bank, from *account, to *account, amount int64, wg *sync.WaitGroup) {
	for i := 0; i < 500; i++ {
		wg.Add(1)
		go func() {
			bank.Transfer(from, to, amount)
			wg.Done()
		}()
	}
}
func concurrentBankTransfer2(bank *Bank, from *account, to *account, amount int64) {
	for i := 0; i < 500; i++ {
		go func() {
			bank.TransferAsync(from, to, amount)
		}()
	}
}

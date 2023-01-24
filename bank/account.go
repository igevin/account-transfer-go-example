package bank

import (
	"runtime"
	"sync"
)

type Account struct {
	Id      int64
	Balance int64
	lock    sync.Mutex
}

func (a *Account) Transfer(to *Account, amount int64) {
	left, right := a, to
	if left.Id > right.Id {
		left, right = to, a
	}

	defer left.lock.Unlock()
	left.lock.Lock()
	defer right.lock.Unlock()
	right.lock.Lock()
	a.TransferUnsafe(to, amount)
}

func (a *Account) TransferUnsafe(to *Account, amount int64) {
	a.Balance -= amount
	// Gosched yields the processor, allowing other goroutines to run. It does not
	// suspend the current goroutine, so execution resumes automatically.
	runtime.Gosched()
	to.Balance += amount
}

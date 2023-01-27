package bank

import (
	"runtime"
	"sync"
)

type AccountV1 struct {
	Account
	lock sync.Mutex
}

func (a *AccountV1) Transfer(to Accountable, amount int64) {
	var too = to.(*AccountV1)
	left, right := a, too
	if left.Id > right.Id {
		left, right = too, a
	}

	defer left.lock.Unlock()
	left.lock.Lock()
	defer right.lock.Unlock()
	right.lock.Lock()
	a.transfer(too, amount)
}

func (a *AccountV1) transfer(to *AccountV1, amount int64) {
	a.Balance -= amount
	// Gosched yields the processor, allowing other goroutines to run. It does not
	// suspend the current goroutine, so execution resumes automatically.
	runtime.Gosched()
	to.Balance += amount
}

package bank

import (
	"runtime"
	"sync"
)

type account struct {
	Id      int64
	Balance int64
	lock    sync.Mutex
}

func (from *account) Transfer(to *account, amount int64) {
	left, right := from, to
	if left.Id > right.Id {
		left, right = to, from
	}

	defer left.lock.Unlock()
	left.lock.Lock()
	defer right.lock.Unlock()
	right.lock.Lock()
	from.Balance -= amount
	runtime.Gosched()
	to.Balance += amount
}

func (from *account) TransferUnsafe(to *account, amount int64) {
	from.Balance -= amount
	runtime.Gosched()
	to.Balance += amount
}

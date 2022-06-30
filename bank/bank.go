package bank

import (
	"runtime"
	"sync"
	"time"
)

type AccountTransfer interface {
	Transfer(to *account, amount int64)
}

type InstituteTransfer interface {
	Transfer(from *account, to *account, amount int64)
}

type Bank struct {
	lock  sync.Mutex
	trans chan TransferTask
}

type TransferTask struct {
	from   *account
	to     *account
	amount int64
}

func (bank *Bank) Transfer(from *account, to *account, amount int64) {
	defer bank.lock.Unlock()
	bank.lock.Lock()
	from.Balance -= amount
	runtime.Gosched()
	to.Balance += amount
}

func (bank *Bank) TransferAsync(from *account, to *account, amount int64) {
	task := TransferTask{from: from, to: to, amount: amount}
	bank.trans <- task
}

func (bank *Bank) TransferHandler() {
	go func() {
		for {
			select {
			case task := <-bank.trans:
				task.from.TransferUnsafe(task.to, task.amount)
			case <-time.After(time.Millisecond * 10):
				close(bank.trans)
				break
			}

		}
	}()
}

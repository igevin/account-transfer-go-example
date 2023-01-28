package bank

import (
	"runtime"
	"sync"
)

type Bank struct {
	lock        sync.Mutex
	trans       chan TransferTask
	closeOnce   sync.Once
	closeSignal chan struct{}
}

func NewBank() *Bank {
	b := &Bank{trans: make(chan TransferTask), closeSignal: make(chan struct{}, 1)}
	b.listenToTransfer()
	return b
}

type TransferTask struct {
	from   Accountable
	to     Accountable
	amount int64
}

func (bank *Bank) Transfer(from Accountable, to Accountable, amount int64) {
	defer bank.lock.Unlock()
	bank.lock.Lock()

	from.SetBalance(from.GetBalance() - amount)
	runtime.Gosched()
	to.SetBalance(to.GetBalance() + amount)
}

func (bank *Bank) TransferAsync(from Accountable, to Accountable, amount int64) {
	task := TransferTask{from: from, to: to, amount: amount}
	bank.trans <- task
}

func (bank *Bank) listenToTransfer() {
	go func() {
		for {
			select {
			case task := <-bank.trans:
				if task.from.GetId() == 0 {
					bank.Close()
					continue
				}
				task.from.Transfer(task.to, task.amount)
			case <-bank.closeSignal:
				close(bank.trans)
				close(bank.closeSignal)
				return
			}

		}
	}()
}

func (bank *Bank) Close() {
	bank.closeOnce.Do(func() {
		bank.closeSignal <- struct{}{}
	})
}

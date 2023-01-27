package bank

import (
	"runtime"
	"sync"
)

type AccountTransfer interface {
	Transfer(to *AccountV1, amount int64)
}

type InstituteTransfer interface {
	Transfer(from *AccountV1, to *AccountV1, amount int64)
}

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
	from   *AccountV1
	to     *AccountV1
	amount int64
}

func (bank *Bank) Transfer(from *AccountV1, to *AccountV1, amount int64) {
	defer bank.lock.Unlock()
	bank.lock.Lock()
	from.Balance -= amount
	runtime.Gosched()
	to.Balance += amount
}

func (bank *Bank) TransferAsync(from *AccountV1, to *AccountV1, amount int64) {
	task := TransferTask{from: from, to: to, amount: amount}
	bank.trans <- task
}

func (bank *Bank) listenToTransfer() {
	go func() {
		for {
			select {
			case task := <-bank.trans:
				task.from.transfer(task.to, task.amount)
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

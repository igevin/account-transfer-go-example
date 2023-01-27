package bank

import (
	"sync"
)

type AccountV1 struct {
	Account
	lock sync.Mutex
}

var _ Accountable = &AccountV1{}

func (a *AccountV1) Transfer(to Accountable, amount int64) {
	too, ok := to.(*AccountV1)
	if !ok {
		return
	}
	left, right := a, too
	if left.Id > right.Id {
		left, right = too, a
	}

	left.lock.Lock()
	defer left.lock.Unlock()
	right.lock.Lock()
	defer right.lock.Unlock()
	a.transfer(to, amount)
}

type AccountV2 struct {
	Account
}

var _ Accountable = &AccountV2{}

func (a *AccountV2) Transfer(to Accountable, amount int64) {

}

type Allocator struct {
}

func (a *Allocator) Apply(from, to any) {

}

func (a *Allocator) Free(from, to any) {

}
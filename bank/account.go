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
	al := getAllocator()
	al.Apply(a, to)
	a.transfer(to, amount)
	al.Free(a, to)
}

var allocator *Allocator
var allocatorOnce sync.Once

type Allocator struct {
	occupied map[any]bool
	mu       *sync.Mutex
	cond     *sync.Cond
}

func (a *Allocator) Apply(from, to any) {
	a.mu.Lock()
	defer a.mu.Unlock()
	for a.occupied[from] || a.occupied[to] {
		a.cond.Wait()
	}
	a.occupied[from] = true
	a.occupied[to] = true
}

func (a *Allocator) Free(from, to any) {
	a.mu.Lock()
	defer a.mu.Unlock()
	delete(a.occupied, from)
	delete(a.occupied, to)
	a.cond.Broadcast()
}

func getAllocator() *Allocator {
	allocatorOnce.Do(func() {
		mu := &sync.Mutex{}
		allocator = &Allocator{
			occupied: make(map[any]bool, 1000),
			mu:       mu,
			cond:     sync.NewCond(mu),
		}
	})
	return allocator
}

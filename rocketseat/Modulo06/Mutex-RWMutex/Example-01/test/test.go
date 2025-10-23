package test

import "sync"

type Account struct {
	mu      sync.RWMutex
	balance int
}

var MyAccount Account

func (a *Account) Deposit(amount int) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.balance += amount
}

func (a *Account) AquireValue() int {
	a.mu.RLock()
	defer a.mu.RUnlock()
	return a.balance
}

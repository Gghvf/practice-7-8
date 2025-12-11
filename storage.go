package main

import (
	"sync"
)

type InMemoryStorage struct {
	mu       sync.RWMutex
	accounts map[string]*Account
}

func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{
		accounts: make(map[string]*Account),
	}
}

func (s *InMemoryStorage) SaveAccount(account *Account) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.accounts[account.ID] = account
	return nil
}

func (s *InMemoryStorage) LoadAccount(accountID string) (*Account, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	account, exists := s.accounts[accountID]
	if !exists {
		return nil, ErrAccountNotFound
	}
	return account, nil
}

func (s *InMemoryStorage) GetAllAccounts() ([]*Account, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	accounts := make([]*Account, 0, len(s.accounts))
	for _, acc := range s.accounts {
		accounts = append(accounts, acc)
	}
	return accounts, nil
}

package domain

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID        string
	Name      string
	Email     string
	APIKey    string
	Balance   float64
	mu        sync.RWMutex
	CreatedAt time.Time
	UpdatedAt time.Time
}

func GenerateAPIKey() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return ""
	}
	return hex.EncodeToString(b)
}

func NewAccount(name, email string) *Account {
	return &Account{
		ID:        uuid.New().String(),
		Name:      name,
		Email:     email,
		Balance:   0,
		APIKey:    GenerateAPIKey(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (a *Account) AddBalance(amount float64) error {
	if amount < 0 {
		return fmt.Errorf("can't add negative balance")
	}
	a.mu.Lock()
	defer a.mu.Unlock()
	a.Balance += amount
	a.UpdatedAt = time.Now()
	return nil
}

func (a *Account) DebitBalance(amount float64) error {
	if a.Balance < amount {
		return fmt.Errorf("insufficient balance")
	}
	a.mu.Lock()
	a.Balance -= amount
	a.UpdatedAt = time.Now()
	a.mu.Unlock()
	return nil
}

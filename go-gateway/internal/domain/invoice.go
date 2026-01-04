package domain

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
)

type Status string

const (
	StatusPending  Status = "pending"
	StatusApproved Status = "approved"
	StatusRejected Status = "rejected"
)

type Invoice struct {
	ID             string
	AccountID      string
	Amount         float64
	Status         Status
	Description    string
	PaymentType    string
	CardLastDigits string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type CreditCard struct {
	Number          string
	CVV             string
	ExpirationMonth int
	ExpirationYear  int
	CardHolderName  string
}

func NewInvoice(accountID string, amount float64, description string, paymentType string, card *CreditCard) (*Invoice, error) {

	if amount <= 0 {
		return nil, ErrInvalidAmount
	}

	lastDigits := card.Number[len(card.Number)-4:]

	invoice := &Invoice{
		ID:             uuid.New().String(),
		AccountID:      accountID,
		Amount:         amount,
		Description:    description,
		PaymentType:    paymentType,
		CardLastDigits: lastDigits,
		Status:         StatusPending,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	return invoice, nil
}

func (i *Invoice) Process() error {
	if i.Amount > 10000 {
		return nil
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var newStatus Status
	if r.Float64() <= 0.7 {
		newStatus = StatusApproved
	} else {
		newStatus = StatusRejected
	}
	return i.UpdateStatus(newStatus)
}

func (i *Invoice) UpdateStatus(status Status) error {
	if i.Status != StatusPending {
		return ErrInvalidStatus
	}
	i.Status = status
	i.UpdatedAt = time.Now()
	return nil
}

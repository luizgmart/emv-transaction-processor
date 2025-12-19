package usecase

import (
    "emv-transaction-processor/internal/domain/valueobject"
)

type Gateway interface {
    Authorize(pan, expiry, cvm string) (bool, error)
}

type ProcessTransaction struct {
    gateway Gateway
}

func NewProcessTransaction(g Gateway) *ProcessTransaction {
    return &ProcessTransaction{gateway: g}
}

type Result struct {
    Approved bool
}

func (uc *ProcessTransaction) Execute(pan, expiry, cvm string) (*Result, error) {
    if err := valueobject.ValidatePAN(pan); err != nil {
        return nil, err
    }
    if err := valueobject.ValidateExpiry(expiry); err != nil {
        return nil, err
    }
    if err := valueobject.ValidateCVM(cvm); err != nil {
        return nil, err
    }

    approved, err := uc.gateway.Authorize(pan, expiry, cvm)
    if err != nil {
        return nil, err
    }
    return &Result{Approved: approved}, nil
}

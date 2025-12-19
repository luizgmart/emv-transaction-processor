package repository

import "emv-transaction-processor/internal/domain/entity"

// TransactionRepository define como transações são persistidas
type TransactionRepository interface {
    Save(transaction entity.Transaction) error
}

package persistence

import (
	"emv-transaction-processor/internal/domain/entity"
	"emv-transaction-processor/internal/usecase"
	"time"
)

// Converte resultado do Use Case para Transaction
func TransactionFromResult(pan string, result *usecase.Result) entity.Transaction {
	return entity.Transaction{
		PAN:       pan,
		Approved:  result.Approved,
		CreatedAt: time.Now(),
	}
}

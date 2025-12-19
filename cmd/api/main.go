package main

import (
	"emv-transaction-processor/internal/adapter/persistence"
	"emv-transaction-processor/internal/adapter/gateway"
	"emv-transaction-processor/internal/usecase"
	"fmt"
)

func main() {
	// Criar o repositório JSON
	repo := &persistence.JSONTransactionRepository{
		File: "transactions.json",
	}

	// Criar o gateway mock
	mockGateway := &gateway.MockAuthorizer{
		Approve: true, // define se a transação será aprovada ou recusada
	}

	// Criar o Use Case, injetando dependências
	processTransaction := usecase.NewProcessTransaction(mockGateway)

	// Dados do cartão simulados
	pan := "4539682995824395"
	expiry := "301231" // YYMMDD
	cvm := "PIN"

	// Executar o Use Case
	result, err := processTransaction.Execute(pan, expiry, cvm)
	if err != nil {
		fmt.Println("Transação rejeitada:", err)
		return
	}

	// Salvar no repositório JSON
	tx := persistence.TransactionFromResult(pan, result)
	if err := repo.Save(tx); err != nil {
		fmt.Println("Erro ao salvar a transação:", err)
		return
	}

	// Exibir resultado
	fmt.Println("Transação aprovada:", result)
}

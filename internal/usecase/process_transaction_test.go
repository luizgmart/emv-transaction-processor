package usecase

import (
    "testing"
    "emv-transaction-processor/internal/adapter/gateway"
)

func TestProcessTransaction(t *testing.T) {
    useCase := NewProcessTransaction(&gateway.MockAuthorizer{Approve: true})

    result, err := useCase.Execute("4539578763621486", "301231", "PIN")
    if err != nil || !result.Approved {
        t.Errorf("Transação válida deveria ser aprovada")
    }

    _, err = useCase.Execute("1234", "301231", "PIN")
    if err == nil {
        t.Errorf("PAN inválido deveria falhar")
    }

    _, err = useCase.Execute("4539578763621486", "301231", "BIOMETRY")
    if err == nil {
        t.Errorf("CVM inválido deveria falhar")
    }

    useCase = NewProcessTransaction(&gateway.MockAuthorizer{Approve: false})
    result, err = useCase.Execute("4539578763621486", "301231", "PIN")
    if err != nil || result.Approved {
        t.Errorf("Transação deveria ser recusada")
    }
}

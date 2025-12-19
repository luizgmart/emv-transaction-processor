package persistence

import (
    "encoding/json"
    "emv-transaction-processor/internal/domain/entity"
    "io/ioutil"
    "os"
)

type JSONTransactionRepository struct {
    File string
}

func (r *JSONTransactionRepository) Save(tx entity.Transaction) error {
    var transactions []entity.Transaction

    if _, err := os.Stat(r.File); err == nil {
        data, _ := ioutil.ReadFile(r.File)
        json.Unmarshal(data, &transactions)
    }

    transactions = append(transactions, tx)
    data, _ := json.MarshalIndent(transactions, "", "  ")
    return ioutil.WriteFile(r.File, data, 0644)
}

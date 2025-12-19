package valueobject

import (
    "testing"
    "time"
)

func TestValidateExpiry(t *testing.T) {
    // Futuro
    future := time.Now().AddDate(1, 0, 0).Format("060102")
    if err := ValidateExpiry(future); err != nil {
        t.Errorf("Data futura válida retornou erro: %v", err)
    }

    // Passado
    past := time.Now().AddDate(-1, 0, 0).Format("060102")
    if err := ValidateExpiry(past); err == nil {
        t.Errorf("Data passada deveria retornar erro")
    }
}

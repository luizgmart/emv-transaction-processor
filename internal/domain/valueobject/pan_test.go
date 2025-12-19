package valueobject

import "testing"

func TestValidatePAN(t *testing.T) {
    valid := "4539578763621486"
    if err := ValidatePAN(valid); err != nil {
        t.Errorf("PAN válido retornou erro: %v", err)
    }

    invalid := "1234"
    if err := ValidatePAN(invalid); err == nil {
        t.Errorf("PAN inválido deveria retornar erro")
    }
}

package valueobject

import "testing"

func TestValidateCVM(t *testing.T) {
    supported := []string{"PIN", "SIGNATURE"}
    for _, cvm := range supported {
        if err := ValidateCVM(cvm); err != nil {
            t.Errorf("CVM suportado %s retornou erro: %v", cvm, err)
        }
    }

    if err := ValidateCVM("BIOMETRY"); err == nil {
        t.Errorf("CVM não suportado deveria falhar")
    }
}

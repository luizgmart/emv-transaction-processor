package valueobject

import "errors"

// ValidateCVM valida se o CVM é suportado
func ValidateCVM(cvm string) error {
    supported := map[string]bool{
        "PIN":       true,
        "SIGNATURE": true,
    }

    if !supported[cvm] {
        return errors.New("CVM não suportado")
    }
    return nil
}

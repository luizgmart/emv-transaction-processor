package valueobject

import (
    "errors"
    "time"
)

// ValidateExpiry valida a data de validade do cartão (YYMMDD)
func ValidateExpiry(expiry string) error {
    t, err := time.Parse("060102", expiry)
    if err != nil {
        return errors.New("formato de validade inválido")
    }
    if t.Before(time.Now()) {
        return errors.New("cartão expirado")
    }
    return nil
}

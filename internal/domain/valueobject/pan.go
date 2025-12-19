package valueobject

import "errors"

// ValidatePAN garante que o PAN é válido pelo algoritmo de Luhn
func ValidatePAN(pan string) error {
    if len(pan) < 13 || len(pan) > 19 {
        return errors.New("PAN deve ter entre 13 e 19 dígitos")
    }
    if !luhn(pan) {
        return errors.New("PAN inválido pelo algoritmo de Luhn")
    }
    return nil
}

func luhn(pan string) bool {
    sum := 0
    alternate := false
    for i := len(pan) - 1; i >= 0; i-- {
        n := int(pan[i] - '0')
        if alternate {
            n *= 2
            if n > 9 {
                n -= 9
            }
        }
        sum += n
        alternate = !alternate
    }
    return sum%10 == 0
}

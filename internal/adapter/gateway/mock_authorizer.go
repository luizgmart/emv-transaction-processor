package gateway

import "errors"

type MockAuthorizer struct {
    Approve bool
}

func (m *MockAuthorizer) Authorize(pan, expiry, cvm string) (bool, error) {
    if pan == "error" {
        return false, errors.New("erro simulado")
    }
    return m.Approve, nil
}

package entity

import "time"

type Transaction struct {
    PAN       string
    Approved  bool
    CreatedAt time.Time
}

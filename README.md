# 💳 EMV Transaction Processor (Clean Architecture)

A simple **EMV transaction processing simulator** implemented in **Go**, following **Clean Architecture** principles. The project validates card data, simulates authorization through a mock gateway, and persists transactions to a JSON file.

## 🏗 Project Structure

```text
emv-transaction-processor/
├── cmd/api                # Application entry point
├── internal/
│   ├── domain/
│   │   ├── valueobject    # PAN, Expiry, CVM
│   │   └── entity         # Domain entities
│   ├── usecase            # Business use cases
│   └── adapter/
│       ├── gateway        # Mock authorization gateway
│       └── persistence   # JSON-based persistence
├── go.mod
└── transactions.json      # Auto-generated
```

## 🔄 EMV Transaction Flow (Simulated)

1. Card data input:

   * PAN (Primary Account Number)
   * Expiry date
   * CVM (Cardholder Verification Method)
2. Business rule validation:

   * PAN length (13–19 digits) with **Luhn algorithm**
   * Expiry date must not be expired
   * Supported CVM (**PIN** or **SIGNATURE**)
3. Authorization via **mock gateway**
4. Transaction result (approved or declined)
5. Persistence of the transaction into a **JSON file**

## 🧪 Automated Tests

Unit tests cover the critical domain logic:

* PAN validation
* Expiry validation
* CVM validation
* Complete use case flow

### Run tests

```bash
go test ./...
```

Expected output:

```text
ok internal/domain/valueobject
ok internal/usecase
```

## ▶️ Running the Application

### Prerequisites

* Go **1.20+**

### Run

From the project root:

```bash
go run ./cmd/api
```

Expected output:

```text
Transaction approved: &{true}
```

The file `transactions.json` will be automatically created with the transaction record.

## 🧾 Persistence Example (JSON)

```json
[
  {
    "PAN": "4539682995824395",
    "Approved": true,
    "CreatedAt": "2025-01-18T20:45:12Z"
  }
]
```

## 🎯 Purpose

This project demonstrates:

* Clean Architecture in Go
* Strong domain validation using Value Objects
* Testable and decoupled business logic
* A simplified EMV transaction processing flow


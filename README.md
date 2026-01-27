# 💳 EMV Transaction Processor

A simple EMV transaction processing simulator built in Go, following Clean Architecture principles. The project validates card data, simulates authorization through a mock gateway, and persists transactions to a JSON file.

🏗 Project Structure
emv-transaction-processor/
├── cmd/api                # Application entry point
├── internal/
│   ├── domain/
│   │   ├── valueobject    # PAN, Expiry, and CVM
│   │   └── entity         # Domain entities
│   ├── usecase            # Use cases (business rules)
│   └── adapter/
│       ├── gateway        # Authorization gateway (mock)
│       └── persistence   # JSON persistence
├── go.mod
└── transactions.json      # Automatically generated

🔄 EMV Transaction Flow (Simulated)
Card data input

PAN (Primary Account Number)

Expiry date

CVM (Cardholder Verification Method)

Business rule validation

PAN between 13 and 19 digits, validated using the Luhn algorithm

Expiry date must not be expired

Supported CVM (PIN or SIGNATURE)

Processing steps

Transaction authorization via mock gateway

Result returned (approved or rejected)

Transaction persisted to a JSON file

🧪 Automated Tests

Unit tests cover critical domain components:

PAN validation

Expiry validation

CVM validation

Full use case flow

Run tests
go test ./...


Expected output:

ok internal/domain/valueobject
ok internal/usecase

▶️ Running the Application
Prerequisites

Go 1.20+

Run the project

From the project root:

go run ./cmd/api


Expected output:

Transaction approved: &{true}


The transactions.json file will be automatically created with the transaction record.

🧾 Persistence Example (JSON)
[
  {
    "PAN": "4539682995824395",
    "Approved": true,
    "CreatedAt": "2025-01-18T20:45:12Z"
  }
]

🎯 Project Purpose

This project demonstrates:

Clean Architecture implementation in Go

Use of Value Objects for strong domain validation

Decoupled and testable business rules

Simplified simulation of an EMV processing flow

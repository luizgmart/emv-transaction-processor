# 💳 EMV Transaction Processor

> Simulação de um módulo de processamento de transações EMV, desenvolvido em Go, seguindo princípios de Clean Architecture, com validações de domínio, mocks, persistência e testes automatizados.

---

## 🎯 Objetivo

Este projeto simula o fluxo de uma transação EMV entre um terminal de pagamento (POS) e um cartão com chip, com foco em:

- Processamento de dados EMV (TLV – Tag-Length-Value)
- Validação de regras essenciais do domínio
- Comunicação simulada com um gateway de pagamento
- Persistência de transações
- Código limpo, testável e bem estruturado

---

## 🧠 Conceitos e Boas Práticas

- Clean Architecture
- Separação de responsabilidades
- Value Objects para regras de domínio
- Use Cases como núcleo da aplicação
- Injeção de dependências
- Mocks para integrações externas
- Testes unitários focados em regras de negócio

---

## 🏗 Estrutura do Projeto

```text
emv-transaction-processor/
├── cmd/api                # Entry point da aplicação
├── internal/
│   ├── domain/
│   │   ├── valueobject    # PAN, Expiry e CVM
│   │   └── entity         # Entidades do domínio
│   ├── usecase            # Casos de uso
│   └── adapter/
│       ├── gateway        # Gateway de autorização (mock)
│       └── persistence    # Persistência em JSON
├── go.mod
└── transactions.json      # Gerado automaticamente

🔄 Fluxo da Transação EMV (Simulado)

Entrada dos dados do cartão:

PAN (Primary Account Number)

Expiry (Data de validade)

CVM (Cardholder Verification Method)

Validação das regras:

PAN entre 13 e 19 dígitos com algoritmo de Luhn

Data de validade não expirada

CVM suportado (PIN ou SIGNATURE)

Autorização via gateway mock

Retorno do resultado (aprovada ou rejeitada)

Persistência da transação em arquivo JSON

🧪 Testes Automatizados

Os testes unitários cobrem as partes críticas do domínio:

Validação de PAN

Validação de Expiry

Validação de CVM

Fluxo completo do caso de uso

Executar os testes
go test ./...


Saída esperada:

ok internal/domain/valueobject
ok internal/usecase

▶️ Executando a Aplicação
Pré-requisitos

Go 1.20+

Rodar o projeto

Na raiz do projeto:

go run ./cmd/api


Saída esperada:

Transação aprovada: &{true}


O arquivo transactions.json será criado automaticamente com o registro da transação.

🧾 Exemplo de Persistência (JSON)
[
  {
    "PAN": "4539682995824395",
    "Approved": true,
    "CreatedAt": "2025-01-18T20:45:12Z"
  }
]

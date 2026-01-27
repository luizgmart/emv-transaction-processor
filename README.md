# 💳 EMV Transaction Processor

Um **simulador simples de processamento de transações EMV**, desenvolvido em **Go**, seguindo os princípios de **Clean Architecture**. O projeto valida dados de cartão, simula uma autorização por meio de um gateway mock e persiste as transações em um arquivo JSON.

## 🏗 Estrutura do Projeto

```text
emv-transaction-processor/
├── cmd/api                # Ponto de entrada da aplicação
├── internal/
│   ├── domain/
│   │   ├── valueobject    # PAN, Expiry e CVM
│   │   └── entity         # Entidades do domínio
│   ├── usecase            # Casos de uso (regras de negócio)
│   └── adapter/
│       ├── gateway        # Gateway de autorização (mock)
│       └── persistence   # Persistência em JSON
├── go.mod
└── transactions.json      # Gerado automaticamente
```

## 🔄 Fluxo da Transação EMV (Simulado)

1. Entrada dos dados do cartão:

   * PAN (Primary Account Number)
   * Data de validade (Expiry)
   * CVM (Cardholder Verification Method)
2. Validação das regras de negócio:

   * PAN entre **13 e 19 dígitos**, validado pelo **algoritmo de Luhn**
   * Data de validade não expirada
   * CVM suportado (**PIN** ou **SIGNATURE**)
3. Autorização da transação via **gateway mock**
4. Retorno do resultado (aprovada ou rejeitada)
5. Persistência da transação em um **arquivo JSON**

## 🧪 Testes Automatizados

Os testes unitários cobrem as partes críticas do domínio:

* Validação de PAN
* Validação de Expiry
* Validação de CVM
* Fluxo completo do caso de uso

### Executar os testes

```bash
go test ./...
```

Saída esperada:

```text
ok internal/domain/valueobject
ok internal/usecase
```

## ▶️ Executando a Aplicação

### Pré-requisitos

* Go **1.20+**

### Rodar o projeto

Na raiz do projeto:

```bash
go run ./cmd/api
```

Saída esperada:

```text
Transação aprovada: &{true}
```

O arquivo `transactions.json` será criado automaticamente com o registro da transação.

## 🧾 Exemplo de Persistência (JSON)

```json
[
  {
    "PAN": "4539682995824395",
    "Approved": true,
    "CreatedAt": "2025-01-18T20:45:12Z"
  }
]
```

## 🎯 Objetivo do Projeto

Este projeto demonstra:

* Aplicação de **Clean Architecture** em Go
* Uso de **Value Objects** para validações fortes de domínio
* Regras de negócio desacopladas e testáveis
* Simulação simplificada de um fluxo de processamento EMV

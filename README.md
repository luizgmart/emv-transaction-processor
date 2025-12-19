EMV Transaction Processor

Simulação de um módulo básico de processamento de transações EMV, desenvolvido em Go, seguindo princípios de Clean Architecture, com validações de domínio, mocks, persistência e testes automatizados.

🎯 Objetivo do Projeto

Este projeto foi desenvolvido como desafio técnico para simular o fluxo de uma transação EMV entre um terminal de pagamento (POS) e um cartão com chip.

O foco está em:

Processamento de dados EMV (simulados via TLV)

Regras de negócio claras e testáveis

Arquitetura limpa e desacoplada

Testes unitários confiáveis

🧠 Conceitos Aplicados

Clean Architecture

Separação de responsabilidades

Value Objects para regras de domínio

Use Cases como núcleo da aplicação

Injeção de dependências

Mocks para simular integrações externas

Testes unitários focados em regras de negócio

🏗 Estrutura do Projeto
emv-transaction-processor/
├── cmd/api                # Entry point da aplicação (main)
├── internal/
│   ├── domain/
│   │   ├── valueobject    # Regras de negócio (PAN, Expiry, CVM)
│   │   └── entity         # Entidades do domínio
│   ├── usecase            # Casos de uso da aplicação
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

Validação dos dados segundo regras EMV:

PAN entre 13 e 19 dígitos + algoritmo de Luhn

Data de validade não expirada

CVM suportado (PIN ou SIGNATURE)

Comunicação com um gateway mock para autorização

Retorno do resultado da transação (aprovada ou rejeitada)

Persistência da transação em arquivo JSON

🧪 Testes Automatizados

Os testes unitários cobrem as partes críticas do domínio, garantindo previsibilidade e segurança:

Validação de PAN

Validação de data de validade

Validação de CVM

Fluxo completo do caso de uso de transação

Executar todos os testes
go test ./...


Saída esperada:

ok internal/domain/valueobject
ok internal/usecase

▶️ Como Executar o Projeto
Pré-requisitos

Go 1.20 ou superior

Executar a aplicação

Na raiz do projeto:

go run ./cmd/api


Saída esperada:

Transação aprovada: &{true}


Após a execução, um arquivo transactions.json será criado automaticamente contendo o registro da transação.

🧩 Exemplo de Saída (JSON)
[
  {
    "PAN": "4539682995824395",
    "Approved": true,
    "CreatedAt": "2025-01-18T20:45:12Z"
  }
]

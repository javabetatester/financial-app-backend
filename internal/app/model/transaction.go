package model

import "time"

type TransactionType string

const (
	Income  TransactionType = "income"
	Expense TransactionType = "expense"
)

// =============================================================================
// ENTIDADES DO BANCO DE DADOS (Database Entities)
// =============================================================================

// Transaction - Entidade principal da transação no banco de dados
// Tags GORM: para mapeamento ORM com o banco
// Tags JSON: para serialização quando retornar dados
type Transaction struct {
	ID          string          `json:"id" gorm:"primaryKey"`                    // Gerado automaticamente
	UserID      string          `json:"user_id" gorm:"not null"`                // Vem da autenticação
	Type        TransactionType `json:"type" gorm:"not null"`                   // income ou expense
	Value       float64         `json:"value" gorm:"not null"`                  // Valor da transação
	Category    string          `json:"category" gorm:"not null"`               // Categoria da transação
	Observation string          `json:"observation"`                            // Observação opcional
	Date        time.Time       `json:"date" gorm:"not null"`                   // Data da transação
	CreatedAt   time.Time       `json:"created_at" gorm:"autoCreateTime"`       // Gerado automaticamente
	UpdatedAt   time.Time       `json:"updated_at" gorm:"autoUpdateTime"`       // Atualizado automaticamente
}

// =============================================================================
// DTOs - DATA TRANSFER OBJECTS (Requisições HTTP)
// =============================================================================

// CreateTransactionRequest - DTO para CRIAR transação
// Vem do BODY da requisição HTTP POST
// Tags binding: para validação automática do Gin
type CreateTransactionRequest struct {
	Type        TransactionType `json:"type" binding:"required,oneof=income expense"`  // Obrigatório
	Value       float64         `json:"value" binding:"required,gt=0"`                 // Obrigatório e > 0
	Category    string          `json:"category" binding:"required"`                   // Obrigatório
	Observation string          `json:"observation"`                                   // Opcional
	Date        time.Time       `json:"date" binding:"required"`                       // Obrigatório
}

// UpdateTransactionRequest - DTO para ATUALIZAR transação
// Vem do BODY da requisição HTTP PUT/PATCH
// Todos os campos são ponteiros (*) = opcionais para update parcial
type UpdateTransactionRequest struct {
	Type        *TransactionType `json:"type,omitempty" binding:"omitempty,oneof=income expense"`
	Value       *float64         `json:"value,omitempty" binding:"omitempty,gt=0"`
	Category    *string          `json:"category,omitempty"`
	Observation *string          `json:"observation,omitempty"`
	Date        *time.Time       `json:"date,omitempty"`
}

// =============================================================================
// DTOs DE RESPOSTA (Response DTOs)
// =============================================================================

// TransactionResponse - DTO para RETORNAR transação (sem dados sensíveis)
type TransactionResponse struct {
	ID          string          `json:"id"`
	Type        TransactionType `json:"type"`
	Value       float64         `json:"value"`
	Category    string          `json:"category"`
	Observation string          `json:"observation"`
	Date        time.Time       `json:"date"`
	CreatedAt   time.Time       `json:"created_at"`
}
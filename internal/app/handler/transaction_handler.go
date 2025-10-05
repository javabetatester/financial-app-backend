package handler

import (
	"github.com/gin-gonic/gin"
	"financial-app-backend/internal/app/model"
	"financial-app-backend/internal/app/service"
	"financial-app-backend/pkg/response"
)

type TransactionHandler struct {
	transactionService service.TransactionService
}

func NewTransactionHandler(transactionService service.TransactionService) *TransactionHandler {
	return &TransactionHandler{
		transactionService: transactionService,
	}
}

// CreateTransaction - POST /transactions
// Body: CreateTransactionRequest (JSON)
// Header: Authorization: Bearer <jwt_token>
func (h *TransactionHandler) CreateTransaction(c *gin.Context) {
	// 1. Pegar o userID do contexto (vem do middleware de autenticação)
	userID, exists := c.Get("userID")
	if !exists {
		response.Unauthorized(c, "Token inválido", "UserID não encontrado no contexto")
		return
	}
	
	// 2. Converter para string (se necessário)
	userIDStr, ok := userID.(string)
	if !ok {
		response.InternalServerError(c, "Erro interno", "UserID inválido")
		return
	}
	
	// 3. Declarar o DTO da requisição
	var req model.CreateTransactionRequest
	
	// 4. Fazer bind do JSON do body para o DTO
	// ShouldBindJSON automaticamente valida as tags binding
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Dados inválidos", err.Error())
		return
	}
	
	// 5. Chamar o service passando o userID + DTO
	// TODO: Implementar no service
	// transaction, err := h.transactionService.CreateTransaction(userIDStr, req)
	// if err != nil {
	//     response.InternalServerError(c, "Erro ao criar transação", err.Error())
	//     return
	// }
	
	// 6. Retornar sucesso com os dados
	response.Success(c, "Transação criada com sucesso", map[string]interface{}{
		"user_id": userIDStr,
		"transaction": req,
	})
}

func (h *TransactionHandler) GetTransactions(c *gin.Context) {
	response.Success(c, "Listar transações", nil)
}

func (h *TransactionHandler) GetTransaction(c *gin.Context) {
	response.Success(c, "Buscar transação", nil)
}

func (h *TransactionHandler) UpdateTransaction(c *gin.Context) {
	response.Success(c, "Atualizar transação", nil)
}

func (h *TransactionHandler) DeleteTransaction(c *gin.Context) {
	response.Success(c, "Deletar transação", nil)
}
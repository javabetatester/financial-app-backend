package service

import (
	"financial-app-backend/internal/app/model"
	"financial-app-backend/internal/app/repository"
)

type TransactionService interface {
	CreateTransaction(userID string, req *model.CreateTransactionRequest) (*model.Transaction, error)
	GetUserTransactions(userID string, page, limit int) ([]*model.Transaction, error)
	GetTransactionByID(userID, transactionID string) (*model.Transaction, error)
	UpdateTransaction(userID, transactionID string, req *model.UpdateTransactionRequest) (*model.Transaction, error)
	DeleteTransaction(userID, transactionID string) error
}

type transactionService struct {
	transactionRepo repository.TransactionRepository
}

func NewTransactionService(transactionRepo repository.TransactionRepository) TransactionService {
	return &transactionService{
		transactionRepo: transactionRepo,
	}
}

func (s *transactionService) CreateTransaction(userID string, req *model.CreateTransactionRequest) (*model.Transaction, error) {
	// TODO: implementar lógica
	return nil, nil
}

func (s *transactionService) GetUserTransactions(userID string, page, limit int) ([]*model.Transaction, error) {
	// TODO: implementar lógica
	return nil, nil
}

func (s *transactionService) GetTransactionByID(userID, transactionID string) (*model.Transaction, error) {
	// TODO: implementar lógica
	return nil, nil
}

func (s *transactionService) UpdateTransaction(userID, transactionID string, req *model.UpdateTransactionRequest) (*model.Transaction, error) {
	// TODO: implementar lógica
	return nil, nil
}

func (s *transactionService) DeleteTransaction(userID, transactionID string) error {
	// TODO: implementar lógica
	return nil
}
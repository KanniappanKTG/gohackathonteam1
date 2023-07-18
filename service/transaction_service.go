package service

import (
	"context"
	"go-app/user-service/db"
	"go-app/user-service/model"
	"go-app/user-service/utils"
)

type transactionService struct {
	repo   db.Repo
	logger utils.LogService
}

func NewTransactionService(repo db.Repo, logger utils.LogService) Service {
	logger.Log().Info("Initiating TransactionService")
	return &transactionService{
		repo:   repo,
		logger: logger,
	}
}

func (tr transactionService) GetTransactions(ctx context.Context, obj model.RequestModel) (transactions []model.Transaction, err error) {
	return tr.repo.GetTransactions(ctx, obj), nil
}

func (tr transactionService) GetTransactionDetails(ctx context.Context, obj model.RequestModel) (transaction model.Transaction, err error) {
	return tr.repo.GetTransactionDetails(ctx, obj), nil
}

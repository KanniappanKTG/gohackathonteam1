package service

import (
	"context"
	"go-app/user-service/model"
)

type Service interface {
	GetTransactions(ctx context.Context, obj model.RequestModel) (transactions []model.Transaction, err error)
	GetTransactionDetails(ctx context.Context, obj model.RequestModel) (transaction model.Transaction, err error)
}

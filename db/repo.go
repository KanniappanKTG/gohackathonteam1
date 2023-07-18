package db

import (
	"context"
	"go-app/user-service/model"
)

type Repo interface {
	GetTransactions(ctx context.Context, req model.RequestModel) (transactions []model.Transaction)
	GetTransactionDetails(ctx context.Context, req model.RequestModel) (transaction model.Transaction)
}

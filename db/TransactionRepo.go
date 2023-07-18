package db

import (
	"context"
	"fmt"
	"go-app/user-service/model"
	"go-app/user-service/utils"
	"strconv"
	"time"

	//"github.com/jmoiron/sqlx"
	//_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	logger1 "gorm.io/gorm/logger"
)

type transactionRepo struct {
	dbConn *gorm.DB
	logger utils.LogService
}

var dbConString = fmt.Sprintf(
	"host=%s port=%s dbname=%s user=%s password='%s' sslmode=disable search_path=%s",
	utils.Cfg.DBHost,
	utils.Cfg.DBPort,
	utils.Cfg.DBName,
	utils.Cfg.DBUserName,
	utils.Cfg.DBPassword,
	utils.Cfg.DBSchema,
)

func NewTransactionRepo(logger utils.LogService) Repo {
	logger.Log().Info("Initiating Transaction Repo")
	logger.Log().Debug(dbConString)
	dbConn, err := gorm.Open(postgres.Open(dbConString), &gorm.Config{
		Logger: logger1.Default.LogMode(logger1.Silent),
	})
	if err != nil {
		logger.Log().Fatal("Unable to open DB Connection", err)
	} else {
		logger.Log().Info("DB connection Sucessful")
	}
	return &transactionRepo{
		dbConn: dbConn,
		logger: logger,
	}
}

func (repo *transactionRepo) GetTransactions(ctx context.Context, req model.RequestModel) (transactions []model.Transaction) {

	currentTime := time.Now()

	toDateTime := req.ToDate

	formattedCurrentDateTime := fmt.Sprintf("%d-%d-%d %d:%d:%d",
		currentTime.Year(),
		currentTime.Month(),
		currentTime.Day(),
		currentTime.Hour(),
		currentTime.Hour(),
		currentTime.Second())

	if len(toDateTime) <= 0 {
		toDateTime = formattedCurrentDateTime
	}

	dateTimeCondition := "tx_ts >= ?  AND tx_ts <= ?"
	fromDateTime := req.FromDate
	if len(fromDateTime) <= 0 {
		dateTimeCondition = "tx_ts <= ?  AND tx_ts <= ?"
		fromDateTime = formattedCurrentDateTime
	}

	pageSizeStr := req.PageSize
	pageSize := 50
	if len(pageSizeStr) > 0 {
		convPageSize, _ := strconv.Atoi(pageSizeStr)
		pageSize = convPageSize
	}
	pageStr := req.Page
	page, _ := strconv.Atoi(pageStr)
	offset := (page - 1) * pageSize
	repo.dbConn.Unscoped().Where(&model.Transaction{AccountId: req.AccId, Status: req.Status}).Where(dateTimeCondition, fromDateTime, toDateTime).Limit(pageSize).Offset(offset).Find(&transactions)

	return transactions
}

func (repo *transactionRepo) GetTransactionDetails(ctx context.Context, req model.RequestModel) (transaction model.Transaction) {
	repo.dbConn.Unscoped().Where(&model.Transaction{TransactionId: req.TransactionId}).Find(&transaction)
	return transaction
}

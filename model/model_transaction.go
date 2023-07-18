package model

import (
	"time"
)

type Transaction struct {
	TransactionId        string    `gorm:"column:tx_id" json:"transactionId,omitempty"`
	AccountId            string    `gorm:"column:acc_id" json:"accountID,omitempty"`
	TransactionTimeStamp time.Time `gorm:"column:tx_ts" json:"transactionTimeStamp,omitempty"`
	Status               string    `gorm:"column:status" json:"transactionStatus,omitempty"`
	Amount               uint      `gorm:"column:amount" json:"transactionAmt,omitempty"`
	MerchantName         string    `gorm:"column:merchantname" json:"merchantName,omitempty"`
	MerchantId           string    `gorm:"column:merchant_id" json:"merchantId,omitempty"`
	TransactionType      string    `gorm:"column:tx_type" json:"transactionType,omitempty"`
	TransactionDetails   string    `gorm:"column:tx_details" json:"transactionDetails,omitempty"`
}

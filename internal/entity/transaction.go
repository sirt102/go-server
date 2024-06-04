package entity

import commonentity "go-server/internal/common/cmentity"

type Transaction struct {
	commonentity.Entity `bson:",inline" json:",inline"`
	TransactionID       string `bson:"transaction_id" json:"transaction_id"`
}

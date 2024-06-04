package entity

import commonentity "go-server/internal/common/cmentity"

type Employee struct {
	commonentity.Entity `bson:",inline" json:",inline"`
	FirstName           string `bson:"first_name" json:"first_name"`
	LastName            string `bson:"last_name" json:"last_name"`
	Department          string `bson:"department" json:"department"`
	Position            string `bson:"position" json:"position"`
}

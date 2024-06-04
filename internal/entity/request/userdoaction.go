package request

import commonentity "go-server/internal/common/cmentity"

type UserDoActionUpdateRequest struct {
	Set         UserDoActionSetUpdateReq   `bson:"$set"`
	SetOnInsert UserDoActionSetOnInsertReq `bson:"$setOnInsert"`
}

type UserDoActionSetUpdateReq struct {
	UpdatedAt commonentity.TimeAt `bson:"updated_at"`
	Status    commonentity.Status `bson:"status,omitempty"`
	Type      *string             `bson:"type,omitempty"`
	MimeType  *string             `bson:"mime_type,omitempty"`
	Size      *int64              `bson:"size,omitempty"`
}

type UserDoActionSetOnInsertReq struct {
	CreatedAt commonentity.TimeAt `bson:"created_at"`
	Status    commonentity.Status `bson:"status"`
	Name      string              `bson:"name"`
	URL       string              `bson:"url"`
	OrgID     commonentity.ID     `bson:"org_id"`
	UserID    string              `bson:"user_id"`
}


package refresh_token_model

import (
	"parishioner_management/internal/common"
	"time"
)

type RefreshTokenModel struct {
	common.BaseModel `json:",inline" bson:",inline"`
	AccountID        int       `bson:"account_id" json:"account_id"`
	Expired          time.Time `bson:"expired" json:"expired"`
	Device           string    `bson:"device" json:"device"`
}

package refresh_token_database

import (
	"parishioner_management/internal/common"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Model struct {
	common.BaseModel `json:",inline" bson:",inline"`
	AccountID        primitive.ObjectID `bson:"account_id" json:"account_id"`
	Expired          time.Time          `bson:"expired" json:"expired"`
	Device           string             `bson:"device,omitempty" json:"device,omitempty"`
	RefreshToken     string             `bson:"refresh_token" json:"refresh_token"`
}

func (u *Model) MarshalBSON() ([]byte, error) {
	if u.CreatedAt.IsZero() {
		u.CreatedAt = time.Now()
		u.DeletedAt = time.Time{}
	}

	u.UpdatedAt = time.Now()

	type my Model
	return bson.Marshal((*my)(u))
}

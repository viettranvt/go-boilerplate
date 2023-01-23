package account_model

import (
	"parishioner_management/internal/common"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type Model struct {
	common.BaseModel `json:",inline" bson:",inline"`
	UserName         string `bson:"username" json:"username"`
	FullName         string `bson:"full_name" json:"full_name"`
	Phone            string `bson:"phone" json:"phone"`
	Email            string `bson:"email" json:"email"`
	Password         string `bson:"password" json:"password"`
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

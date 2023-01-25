package relationship_database

import (
	"parishioner_management/internal/common"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type Model struct {
	common.BaseModel   `json:",inline" bson:",inline"`
	ParishionerID      int `json:"parishioner_id" bson:"parishioner_id"`
	OtherParishionerID int `json:"other_parishioner_id" bson:"other_parishioner_id"`
	Relationship       int `json:"relationship" bson:"relationship"` // protector //father //mother //wife //husband
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

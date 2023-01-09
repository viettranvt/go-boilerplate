package relationship

import "parishioner_management/internal/common"

type Model struct {
	common.BaseModel   `json:",inline" bson:",inline"`
	ParishionerID      int `json:"parishioner_id" bson:"parishioner_id"`
	OtherParishionerID int `json:"other_parishioner_id" bson:"other_parishioner_id"`
	Relationship       int `json:"relationship" bson:"relationship"` // protector //father //mother //wife //husband
}

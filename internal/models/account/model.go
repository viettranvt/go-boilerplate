package account_model

import "parishioner_management/internal/common"

type Model struct {
	common.BaseModel `json:",inline" bson:",inline"`
	UserName         string `bson:"user_name" json:"user_name"`
	FullName         string `bson:"full_name" json:"full_name"`
	Phone            string `bson:"phone" json:"phone"`
	Email            string `bson:"email" json:"email"`
	Password         string `bson:"password" json:"password"`
}

package parishioner

import (
	"parishioner_management/internal/common"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type Model struct {
	common.BaseModel     `json:",inline" bson:",inline"`
	DateOfBirth          time.Time `json:"date_of_birth,omitempty" bson:"date_of_birth,omitempty"`                     // ngay sinh
	DateOfDeath          time.Time `json:"date_of_death,omitempty" bson:"date_of_death,omitempty"`                     // ngay mat
	DateOfWedding        time.Time `json:"date_of_wedding,omitempty" bson:"date_of_wedding,omitempty"`                 // ngay cuoi
	DateOfBaptism        time.Time `json:"date_of_baptism,omitempty" bson:"date_of_baptism,omitempty"`                 // ngay rua toi
	DateOfFirstCommunion time.Time `json:"date_of_first_communion,omitempty" bson:"date_of_first_communion,omitempty"` // ngay ruoc le lan dau
	DateOfConfirmation   time.Time `json:"date_of_confirmation,omitempty" bson:"date_of_confirmation,omitempty"`       // ngay them suc
	DateOfOath           time.Time `json:"date_of_oath,omitempty" bson:"date_of_oath,omitempty"`                       // ngay tuyen hua bao dong
	DateOfHolyOrder      time.Time `json:"date_of_holy_order,omitempty" bson:"date_of_holy_order,omitempty"`           // ngay truyen chuc thanh
	ParishName           string    `json:"parish_name" bson:"parish_name"`                                             // ten giao ho
	ChristianName        string    `json:"christian_name" bson:"christian_name"`                                       // ten thanh
	FullName             string    `json:"full_name" bson:"full_name"`
	Note                 string    `json:"note,omitempty" bson:"note,omitempty"`
	Gender               string    `json:"gender" bson:"gender"`
	Address              string    `json:"address,omitempty" bson:"address,omitempty"`
	WardID               int       `json:"ward_id" bson:"ward_id"`
	DistrictID           int       `json:"district_id" bson:"district_id"`
	ProvinceID           int       `json:"province_id" bson:"province_id"`
	Avatar               string    `json:"avatar,omitempty" bson:"avatar,omitempty"` // hinh anh

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

package parishioner

import (
	"parishioner_management/internal/common"
	"time"
)

type Model struct {
	common.BaseModel     `json:",inline" bson:",inline"`
	DateOfBirth          time.Time `json:"date_of_birth" bson:"date_of_birth"`                     // ngay sinh
	DateOfDeath          time.Time `json:"date_of_death" bson:"date_of_death"`                     // ngay mat
	DateOfWedding        time.Time `json:"date_of_wedding" bson:"date_of_wedding"`                 // ngay cuoi
	DateOfBaptism        time.Time `json:"date_of_baptism" bson:"date_of_baptism"`                 // ngay rua toi
	DateOfFirstCommunion time.Time `json:"date_of_first_communion" bson:"date_of_first_communion"` // ngay ruoc le lan dau
	DateOfConfirmation   time.Time `json:"date_of_confirmation" bson:"date_of_confirmation"`       // ngay them suc
	DateOfOath           time.Time `json:"date_of_oath" bson:"date_of_oath"`                       // ngay tuyen hua bao dong
	DateOfHolyOrder      time.Time `json:"date_of_holy_order" bson:"date_of_holy_order"`           // ngay truyen chuc thanh
	ParishName           string    `json:"parish_name" bson:"parish_name"`                         // ten giao ho
	ChristianName        string    `json:"christian_name" bson:"christian_name"`                   // ten thanh
	FullName             string    `json:"full_name" bson:"full_name"`
	Note                 string    `json:"note" bson:"note"`
	Gender               string    `json:"gender" bson:"gender"`
	Address              string    `json:"address" bson:"address"`
	WardID               int       `json:"ward_id" bson:"ward_id"`
	DistrictID           int       `json:"district_id" bson:"district_id"`
	ProvinceID           int       `json:"province_id" bson:"province_id"`
	Avatar               string    `json:"avatar" bson:"avatar"` // hinh anh

}

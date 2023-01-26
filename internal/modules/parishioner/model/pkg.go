package parishioner_model

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

type ParishionerCreation struct {
	DateOfBirth          *int64  `json:"date_of_birth"`           // ngay sinh, millisecond
	DateOfDeath          *int64  `json:"date_of_death"`           // ngay mat, millisecond
	DateOfWedding        *int64  `json:"date_of_wedding"`         // ngay cuoi, millisecond
	DateOfBaptism        *int64  `json:"date_of_baptism"`         // ngay rua toi, millisecond
	DateOfFirstCommunion *int64  `json:"date_of_first_communion"` // ngay ruoc le lan dau, millisecond
	DateOfConfirmation   *int64  `json:"date_of_confirmation"`    // ngay them suc, millisecond
	DateOfOath           *int64  `json:"date_of_oath"`            // ngay tuyen hua bao dong, millisecond
	DateOfHolyOrder      *int64  `json:"date_of_holy_order"`      // ngay truyen chuc thanh, millisecond
	ParishName           string  `json:"parish_name"`             // ten giao ho
	ChristianName        string  `json:"christian_name"`          // ten thanh
	Avatar               *string `json:"avatar"`                  // hinh anh
	FullName             string  `json:"full_name"`
	Note                 *string `json:"note"`
	Gender               string  `json:"gender"`
	Address              *string `json:"address"`

	// TODO add relation for parishioner
}

func (data ParishionerCreation) Validate() error {
	return validation.ValidateStruct(&data,
		validation.Field(&data.ParishName, validation.Required),
		validation.Field(&data.ChristianName, validation.Required),
		validation.Field(&data.FullName, validation.Required),
		validation.Field(&data.Gender, validation.Required),
	)
}

type ParishionerListResponse struct {
	FullName      string     `json:"full_name"`
	DateOfBirth   *time.Time `json:"date_of_birth,omitempty"` // ngay sinh
	ParishName    string     `json:"parish_name"`             // ten giao ho
	ChristianName string     `json:"christian_name"`          // ten thanh
	Gender        string     `json:"gender"`
}

package parishioner_business

import (
	"context"
	parishioner_database "parishioner_management/internal/databases/parishioner"
	parishioner_model "parishioner_management/internal/modules/parishioner/model"
	date_util "parishioner_management/internal/utils/date"
	string_util "parishioner_management/internal/utils/string"
)

type CreateStorage interface {
	InsertParishioner(
		ctx context.Context,
		model *parishioner_database.Model,
	) error
}

type createBusiness struct {
	store CreateStorage
}

func NewCreateBusiness(store CreateStorage) *createBusiness {
	return &createBusiness{store: store}
}

func (biz *createBusiness) Create(ctx context.Context, model *parishioner_database.Model) error {
	return biz.store.InsertParishioner(ctx, model)
}

// this func will convert input data to parishioner model
func (biz *createBusiness) ToModel(ctx context.Context, data *parishioner_model.ParishionerCreation) *parishioner_database.Model {
	model := &parishioner_database.Model{
		ParishName:    string_util.NormalizedData(data.ParishName, true),
		ChristianName: string_util.NormalizedData(data.ChristianName, true),
		Avatar:        data.Avatar,
		FullName:      data.FullName,
		Note:          data.Note,
		Gender:        data.Gender,
		Address:       data.Address,
	}

	if data.DateOfBirth != nil {
		date := date_util.ConvertMillisecondToTime(*data.DateOfBirth)
		model.DateOfBirth = &date
	}

	if data.DateOfDeath != nil {
		date := date_util.ConvertMillisecondToTime(*data.DateOfDeath)
		model.DateOfDeath = &date
	}

	if data.DateOfWedding != nil {
		date := date_util.ConvertMillisecondToTime(*data.DateOfWedding)
		model.DateOfWedding = &date
	}

	if data.DateOfBaptism != nil {
		date := date_util.ConvertMillisecondToTime(*data.DateOfBaptism)
		model.DateOfBaptism = &date
	}

	if data.DateOfFirstCommunion != nil {
		date := date_util.ConvertMillisecondToTime(*data.DateOfFirstCommunion)
		model.DateOfFirstCommunion = &date
	}

	if data.DateOfConfirmation != nil {
		date := date_util.ConvertMillisecondToTime(*data.DateOfConfirmation)
		model.DateOfConfirmation = &date
	}

	if data.DateOfOath != nil {
		date := date_util.ConvertMillisecondToTime(*data.DateOfOath)
		model.DateOfOath = &date
	}

	if data.DateOfHolyOrder != nil {
		date := date_util.ConvertMillisecondToTime(*data.DateOfHolyOrder)
		model.DateOfHolyOrder = &date
	}

	return model
}

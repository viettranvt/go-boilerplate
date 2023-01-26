package parishioner_business

import (
	"context"
	"parishioner_management/internal/common"
	parishioner_database "parishioner_management/internal/databases/parishioner"
	parishioner_model "parishioner_management/internal/modules/parishioner/model"
)

type GetListStorage interface {
	GetParishionerListHasPagination(
		ctx context.Context,
		filters map[string]interface{},
		paging *common.Paging,
		withDeleted bool,
		moreKeys ...interface{},
	) ([]parishioner_database.Model, error)
}

type getListBusiness struct {
	store GetListStorage
}

func NewGetListBusiness(store GetListStorage) *getListBusiness {
	return &getListBusiness{store: store}
}

// this func will get list of parishioner has pagination
func (biz *getListBusiness) GetList(ctx context.Context, paging *common.Paging) ([]parishioner_database.Model, error) {
	return biz.store.GetParishionerListHasPagination(ctx, nil, paging, false)
}

func (biz *getListBusiness) ToResponse(
	ctx context.Context, parishionerList []parishioner_database.Model,
) []parishioner_model.ParishionerListResponse {
	result := make([]parishioner_model.ParishionerListResponse, 0, len(parishionerList))

	for _, data := range parishionerList {
		parishionerData := parishioner_model.ParishionerListResponse{
			FullName:      data.FullName,
			DateOfBirth:   data.DateOfBirth,
			ParishName:    data.ParishName,
			ChristianName: data.ChristianName,
			Gender:        data.Gender,
		}

		result = append(result, parishionerData)
	}

	return result
}

package parishioner_business

import (
	"context"
	parishioner_database "parishioner_management/internal/databases/parishioner"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UpdateStorage interface {
	UpdateParishioner(
		ctx context.Context,
		id primitive.ObjectID,
		dataUpdate map[string]interface{},
	) error

	GetParishioner(
		ctx context.Context,
		filters map[string]interface{},
		withDeleted bool,
		moreKeys ...interface{},
	) (*parishioner_database.Model, error)
}

type updateBusiness struct {
	store UpdateStorage
}

func NewUpdateBusiness(store UpdateStorage) *updateBusiness {
	return &updateBusiness{store: store}
}

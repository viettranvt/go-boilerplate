package parishioner_business

import (
	"context"
	parishioner_database "parishioner_management/internal/databases/parishioner"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DeleteStorage interface {
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

type deleteBusiness struct {
	store UpdateStorage
}

func NewDeleteBusiness(store DeleteStorage) *deleteBusiness {
	return &deleteBusiness{store: store}
}

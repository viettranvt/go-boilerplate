package parishioner_business

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UpdateStorage interface {
	UpdateParishioner(
		ctx context.Context,
		id primitive.ObjectID,
		dataUpdate map[string]interface{},
	) error
}

type updateBusiness struct {
	store UpdateStorage
}

func NewUpdateBusiness(store UpdateStorage) *updateBusiness {
	return &updateBusiness{store: store}
}

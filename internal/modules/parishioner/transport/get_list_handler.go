package parishioner_transport

import (
	"net/http"
	"parishioner_management/internal/common"
	parishioner_business "parishioner_management/internal/modules/parishioner/business"
	parishioner_storage "parishioner_management/internal/modules/parishioner/storage"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetList(db *mongo.Client) func(*gin.Context) {
	return func(context *gin.Context) {
		var paging common.Paging

		if err := context.ShouldBind(&paging); err != nil {
			response := common.NewFullErrorResponse(http.StatusBadRequest, nil, err.Error(), err.Error(), err.Error())
			context.JSON(response.StatusCode, response)
			return
		}

		// normalize input data
		paging.Fulfill()

		ctx := context.Request.Context()
		store := parishioner_storage.NewMongoStore(db)
		business := parishioner_business.NewGetListBusiness(store)

		result, err := business.GetList(ctx, &paging)

		if err != nil {
			response := common.NewFullErrorResponse(
				http.StatusBadRequest, nil, err.Error(), err.Error(), err.Error(),
			)
			context.JSON(response.StatusCode, response)
			return
		}

		response := common.NewSuccessResponse(business.ToResponse(ctx, result), paging, nil)
		context.JSON(response.StatusCode, response)
	}
}

package parishioner_transport

import (
	"net/http"
	"parishioner_management/internal/common"
	parishioner_business "parishioner_management/internal/modules/parishioner/business"
	parishioner_model "parishioner_management/internal/modules/parishioner/model"
	parishioner_storage "parishioner_management/internal/modules/parishioner/storage"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func Create(db *mongo.Client) func(*gin.Context) {
	return func(context *gin.Context) {
		var inputData parishioner_model.ParishionerCreation

		if err := context.ShouldBind(&inputData); err != nil {
			response := common.NewFullErrorResponse(http.StatusBadRequest, nil, err.Error(), err.Error(), err.Error())
			context.JSON(response.StatusCode, response)
			return
		}

		if err := inputData.Validate(); err != nil {
			response := common.NewFullErrorResponse(http.StatusBadRequest, nil, err.Error(), err.Error(), err.Error())
			context.JSON(response.StatusCode, response)
			return
		}

		ctx := context.Request.Context()
		store := parishioner_storage.NewMongoStore(db)
		business := parishioner_business.NewCreateBusiness(store)

		// convert data to parishioner model
		model := business.ToModel(ctx, &inputData)

		if err := business.Create(ctx, model); err != nil {
			response := common.NewFullErrorResponse(http.StatusBadRequest, nil, err.Error(), err.Error(), err.Error())
			context.JSON(response.StatusCode, response)
			return
		}

		response := common.SimpleSuccessResponse(map[string]string{
			"message": "create parishioner successful",
		})
		context.JSON(response.StatusCode, response)
	}
}

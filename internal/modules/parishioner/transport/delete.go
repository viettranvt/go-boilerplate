package parishioner_transport

import (
	"net/http"
	"parishioner_management/internal/common"
	parishioner_model "parishioner_management/internal/modules/parishioner/model"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func Delete(db *mongo.Client) func(*gin.Context) {
	return func(context *gin.Context) {
		ids := common.BindIds(*context)

		if len(ids) < 1 {
			response := common.NewFullErrorResponse(
				http.StatusBadRequest, nil,
				parishioner_model.ErrInvalidId.Error(),
				parishioner_model.ErrInvalidId.Error(),
				parishioner_model.ErrInvalidId.Error(),
			)
			context.JSON(response.StatusCode, response)
			return
		}

		response := common.SimpleSuccessResponse(map[string]interface{}{
			"message": ids,
		})
		context.JSON(response.StatusCode, response)
	}
}

package auth_transport

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"

	"parishioner_management/internal/common"
	auth_business "parishioner_management/internal/modules/auth/business"
	auth_model "parishioner_management/internal/modules/auth/model"
	auth_storage "parishioner_management/internal/modules/auth/storage"
)

func Login(db *mongo.Client) func(*gin.Context) {
	return func(context *gin.Context) {
		var inputData auth_model.UserLogin

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
		store := auth_storage.NewMongoStore(db)
		business := auth_business.NewLoginBusiness(store)
		result, err := business.GetAccountByUserName(ctx, inputData.Username)

		if err != nil {
			response := common.NewFullErrorResponse(
				http.StatusBadRequest, nil,
				auth_model.ErrUsernameOrPasswordNotMatch.Error(),
				auth_model.ErrUsernameOrPasswordNotMatch.Error(),
				auth_model.ErrUsernameOrPasswordNotMatch.Error(),
			)
			context.JSON(response.StatusCode, response)
			return
		}

		if err := business.AuthenticatePassword(ctx, inputData.Password, result.Password); err != nil {
			response := common.NewFullErrorResponse(http.StatusBadRequest, nil, err.Error(), err.Error(), err.Error())
			context.JSON(response.StatusCode, response)
			return
		}

		if err != nil {
			response := common.NewFullErrorResponse(http.StatusBadRequest, nil, err.Error(), err.Error(), err.Error())
			context.JSON(response.StatusCode, response)
			return
		}

		token, err := business.GenAccessToken(ctx, result.ID.Hex())

		if err != nil {
			response := common.NewFullErrorResponse(http.StatusBadRequest, nil, err.Error(), err.Error(), err.Error())
			context.JSON(response.StatusCode, response)
			return
		}

		tokenExpirationTime, err := business.GetTokenExpirationTime(ctx, token)

		if err != nil {
			response := common.NewFullErrorResponse(http.StatusBadRequest, nil, err.Error(), err.Error(), err.Error())
			context.JSON(response.StatusCode, response)
			return
		}

		refreshToken, err := business.GenRefreshToken(ctx, result.ID, context.Request.UserAgent())

		if err != nil {
			response := common.NewFullErrorResponse(http.StatusBadRequest, nil, err.Error(), err.Error(), err.Error())
			context.JSON(response.StatusCode, response)
			return
		}

		response := common.SimpleSuccessResponse(business.ToResponse(ctx, *result, token, tokenExpirationTime, refreshToken.RefreshToken))
		context.JSON(response.StatusCode, response)
	}
}

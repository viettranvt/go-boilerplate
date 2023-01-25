package auth_transport

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"

	"parishioner_management/internal/common"
	auth_business "parishioner_management/internal/modules/auth/business"
	auth_model "parishioner_management/internal/modules/auth/model"
	auth_storage "parishioner_management/internal/modules/auth/storage"
)

func Refresh(db *mongo.Client) func(*gin.Context) {
	return func(context *gin.Context) {
		var inputData auth_model.UserRefreshToken

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
		business := auth_business.NewRefreshTokenBusiness(store)
		refreshTokenInfo, err := business.FindRefreshTokenInfo(ctx, inputData.RefreshToken)

		if err != nil {
			response := common.NewUnauthorized(nil, auth_model.ErrInvalidToken.Error(), auth_model.ErrInvalidToken.Error())
			context.JSON(response.StatusCode, response)
			return
		}

		accountInfo, err := business.FindAccountByID(ctx, refreshTokenInfo.AccountID)

		if err != nil {
			response := common.NewUnauthorized(nil, auth_model.ErrInvalidToken.Error(), auth_model.ErrInvalidToken.Error())
			context.JSON(response.StatusCode, response)
			return
		}

		if refreshTokenInfo.Expired.Before(time.Now()) {
			response := common.NewUnauthorized(nil, auth_model.ErrInvalidToken.Error(), auth_model.ErrInvalidToken.Error())
			context.JSON(response.StatusCode, response)
			return
		}

		token, err := business.GenAccessToken(ctx, accountInfo.ID.Hex())

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

		response := common.SimpleSuccessResponse(auth_model.RefreshTokenResponse{
			Token:               token,
			TokenExpirationTime: tokenExpirationTime,
		})
		context.JSON(response.StatusCode, response)
	}
}

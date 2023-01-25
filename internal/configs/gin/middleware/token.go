package gin_middleware

import (
	"net/http"
	"parishioner_management/internal/common"
	"parishioner_management/internal/components"
	database_field_const "parishioner_management/internal/constant/database/field"
	database_model_const "parishioner_management/internal/constant/database/model"
	token_const "parishioner_management/internal/constant/token"
	account_database "parishioner_management/internal/databases/account"
	database_util "parishioner_management/internal/utils/database"
	options_util "parishioner_management/internal/utils/options"
	token_util "parishioner_management/internal/utils/token"
	"strings"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func customToken(appContext components.AppContext, options options_util.Options) gin.HandlerFunc {
	return func(c *gin.Context) {
		// skip when api route is auth
		if DefaultAuthSkipper(c) {
			c.Next()
			return
		}

		// get token info from header
		authHeader := c.GetHeader(token_const.AuthHeaderKey)

		if authHeader == "" {
			response := common.NewFullErrorResponse(http.StatusBadRequest, nil, "invalid token", "invalid token", "invalid token")
			c.AbortWithStatusJSON(response.StatusCode, response)
			return
		}

		//split to get schema and token
		authHeaderSplit := strings.Split(authHeader, " ")

		if len(authHeaderSplit) != 2 {
			response := common.NewFullErrorResponse(http.StatusBadRequest, nil, "invalid token", "invalid token", "invalid token")
			c.AbortWithStatusJSON(response.StatusCode, response)
			return
		}

		schema := authHeaderSplit[0]
		token := authHeaderSplit[1]

		// verify schema
		if schema != token_const.Schema {
			response := common.NewFullErrorResponse(http.StatusBadRequest, nil, "invalid token", "invalid token", "invalid token")
			c.AbortWithStatusJSON(response.StatusCode, response)
			return
		}

		if token == "" {
			response := common.NewFullErrorResponse(http.StatusBadRequest, nil, "invalid token", "invalid token", "invalid token")
			c.AbortWithStatusJSON(response.StatusCode, response)
			return
		}

		ctx := c.Request.Context()
		// decode token
		data, err := token_util.Decode(ctx, token, options.JwtSecretKey)

		if err != nil {
			response := common.NewFullErrorResponse(http.StatusBadRequest, nil, err.Error(), err.Error(), err.Error())
			c.AbortWithStatusJSON(response.StatusCode, response)
			return
		}

		//get info in token
		accountID, found := data[token_const.UserIDKey]

		if !found {
			response := common.NewFullErrorResponse(http.StatusBadRequest, nil, "invalid token", "invalid token", "invalid token")
			c.AbortWithStatusJSON(response.StatusCode, response)
			return
		}

		expiredToken, found := data[token_const.ExpiredTokenKey]

		if !found {
			response := common.NewFullErrorResponse(http.StatusBadRequest, nil, "invalid token", "invalid token", "invalid token")
			c.AbortWithStatusJSON(response.StatusCode, response)
			return
		}

		// convert id to object id
		id, err := primitive.ObjectIDFromHex(accountID.(string))

		if err != nil {
			response := common.NewFullErrorResponse(http.StatusBadRequest, nil, "invalid token", "invalid token", "invalid token")
			c.AbortWithStatusJSON(response.StatusCode, response)
			return
		}

		// get account info in db
		mongoClient := appContext.GetMainMongoDBConnection()
		accountCollection := database_util.GetCollection(mongoClient, database_model_const.Account)
		filter := bson.M{database_field_const.ID: id}
		var accountInfo account_database.Model

		if err := accountCollection.FindOne(ctx, filter).Decode(&accountInfo); err != nil {
			response := common.NewFullErrorResponse(http.StatusBadRequest, nil, "invalid token", "invalid token", "invalid token")
			c.AbortWithStatusJSON(response.StatusCode, response)
			return
		}

		c.Set(token_const.BasicAccountInfoKey, accountInfo)
		c.Set(token_const.ExpiredTokenKey, expiredToken.(float64)*1000)

		c.Next()
	}
}

func tokenMiddleware(g *gin.Engine, appContext components.AppContext, options options_util.Options) {
	g.Use(customToken(appContext, options))
}

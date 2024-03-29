package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"parishioner_management/internal/common"
	"parishioner_management/internal/components"
	configs "parishioner_management/internal/configs"
	gin_middleware "parishioner_management/internal/configs/gin/middleware"
	gin_routes "parishioner_management/internal/routes/gin"
	options_util "parishioner_management/internal/utils/options"
)

type ConsoleParam struct {
	EnvPath string `json:"env_path"`
}

type server struct {
	options options_util.Options
}

func defaultOptions(envFile string) options_util.Options {
	if err := godotenv.Load(envFile); err != nil {
		log.Fatal("Error loading .env file")
	}

	return options_util.Options{
		MongoDBUrl:   os.Getenv(configs.EnvMongoDBUrl),
		Port:         os.Getenv(configs.EnvPort),
		JwtSecretKey: os.Getenv(configs.EnvJwtSecretKey),
		APIPrefix:    os.Getenv(configs.EnvAPIPrefix),
		Mode:         os.Getenv(configs.EnvMode),
	}
}

func parseParam() *ConsoleParam {
	param := &ConsoleParam{}
	flag.StringVar(&param.EnvPath, "env-path", ".env", "dotenv file path")
	flag.Parse()

	return param
}

func (s *server) start(appContext components.AppContext) error {
	gin, err := s.createAndConfigGin(appContext)

	if err != nil {
		return err
	}

	gin_routes.RegisterAllModules(gin, appContext, s.options.APIPrefix)
	address := fmt.Sprintf(":%v", s.options.Port)

	return gin.Run(address)
}

func (s *server) createAndConfigGin(appContext components.AppContext) (*gin.Engine, error) {
	if s.options.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	r.SetTrustedProxies(nil)
	// s.configLog(e)
	// s.configErrHandler(e)

	// setup response when routing not found
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, common.NewFullErrorResponse(http.StatusNotFound, nil, "page not found", "page not found", "PAGE_NOT_FOUND"))
	})

	if err := gin_middleware.RegisterMiddleware(r, appContext, s.options); err != nil {
		return nil, err
	}

	return r, nil
}

func main() {
	param := parseParam()
	optionsServer := defaultOptions(param.EnvPath)
	ctx := context.Background()
	mongoDB, err := mongo.Connect(ctx, options.Client().ApplyURI(optionsServer.MongoDBUrl))

	if err != nil {
		log.Fatalln(err)
	}

	// test connection of db
	if err := mongoDB.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatalln("database connection failed", err)
	}

	appContext := components.NewAppContext(mongoDB)
	server := &server{optionsServer}

	if err := server.start(appContext); err != nil {
		log.Fatalf("fail to start, err=%v", err)
	}
}

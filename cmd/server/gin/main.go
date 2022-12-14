package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

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
	r := gin.Default()

	// s.configLog(e)
	// s.configErrHandler(e)

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

	appContext := components.NewAppContext(mongoDB)
	server := &server{optionsServer}

	if err := server.start(appContext); err != nil {
		log.Fatalf("fail to start, err=%v", err)
	}
}

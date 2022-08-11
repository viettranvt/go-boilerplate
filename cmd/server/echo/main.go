package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"init_golang/internal/components"
	configs "init_golang/internal/configs"
	echo_middleware "init_golang/internal/configs/echo/middleware"
	echo_routes "init_golang/internal/routes/echo"
	options_util "init_golang/internal/utils/options"
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
		MySqlUrl:     os.Getenv(configs.EnvMySqlUrl),
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
	e, err := s.createAndConfigEcho(appContext)

	if err != nil {
		return err
	}

	echo_routes.RegisterAllModules(e, appContext, s.options.APIPrefix)
	address := fmt.Sprintf(":%v", s.options.Port)

	return e.Start(address)
}

func (s *server) createAndConfigEcho(appContext components.AppContext) (*echo.Echo, error) {
	r := echo.New()

	// s.configLog(e)
	// s.configErrHandler(e)

	if err := echo_middleware.RegisterMiddleware(r, appContext, s.options); err != nil {
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
	mySqlDB, err := gorm.Open(mysql.Open(optionsServer.MySqlUrl), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	appContext := components.NewAppContext(mySqlDB, mongoDB)
	server := &server{optionsServer}

	if err := server.start(appContext); err != nil {
		log.Fatalf("fail to start, err=%v", err)
	}
}

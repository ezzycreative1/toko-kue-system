package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"

	"github.com/ezzycreative1/toko-kue-system/pkg/constants"
	"github.com/ezzycreative1/toko-kue-system/pkg/glog"
)

type EnvironmentConfig struct {
	Env        string
	AppConfig  AppConfig
	Database   Database
	GrpcServer GrpcServer
	Log        glog.Logger
}

type AppConfig struct {
	Name           string
	Version        string
	Port           int
	MaxRequestTime int
}

type Log struct {
	Path      string
	Prefix    string
	Extension string
}

type GrpcServer struct {
	Port int
}

func LoadENVConfig() (config EnvironmentConfig, err error) {
	err = godotenv.Load(".env")
	if err != nil {
		err = fmt.Errorf(constants.ErrLoadENV, err)
		return EnvironmentConfig{}, err
	}

	port, err := strconv.Atoi(os.Getenv("APP_PORT"))
	if err != nil {
		err = fmt.Errorf(constants.ErrConvertStringToInt, err)
		return EnvironmentConfig{}, err
	}

	gRPCPort, err := strconv.Atoi(os.Getenv("GRPC_SERVER_PORT"))
	if err != nil {
		err = fmt.Errorf(constants.ErrConvertStringToInt, err)
		return EnvironmentConfig{}, err
	}

	config = EnvironmentConfig{
		Env: os.Getenv("ENV"),
		AppConfig: AppConfig{
			Name:           os.Getenv("APP_NAME"),
			Version:        os.Getenv("APP_VERSION"),
			Port:           port,
			MaxRequestTime: cast.ToInt(os.Getenv("APP_MAX_REQUEST_TIME")),
		},
		Database: Database{
			Engine:          os.Getenv("DATABASE_ENGINE"),
			Host:            os.Getenv("DATABASE_HOST"),
			Port:            cast.ToInt(os.Getenv("DATABASE_PORT")),
			Username:        os.Getenv("KBS_DATABASE_USERNAME"),
			Password:        os.Getenv("KBS_DATABASE_PASSWORD"),
			DBName:          os.Getenv("KBS_DATABASE_NAME"),
			Schema:          os.Getenv("KBS_DATABASE_SCHEMA"),
			MaxIdle:         cast.ToInt(os.Getenv("KBS_DATABASE_MAX_IDLE")),
			MaxConn:         cast.ToInt(os.Getenv("KBS_DATABASE_MAX_CONN")),
			ConnMaxLifetime: cast.ToInt(os.Getenv("KBS_DATABASE_CONN_LIFETIME")),
		},
		GrpcServer: GrpcServer{
			Port: gRPCPort,
		},
	}

	return
}

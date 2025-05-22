package deps

import (
	"log"

	config "github.com/ezzycreative1/toko-kue-system/configs"
	"github.com/ezzycreative1/toko-kue-system/internal/adapters/v1/repositories/ingredientRepo"
	"github.com/ezzycreative1/toko-kue-system/internal/core/services/v1/ingredientService"
	"github.com/ezzycreative1/toko-kue-system/internal/infra/db"
	"github.com/ezzycreative1/toko-kue-system/pkg/bvalidator"
	"github.com/ezzycreative1/toko-kue-system/pkg/glog"
)

const (
	keyTransaction = "kbs-ctx"
	timeout        = 60
)

type Dependency struct {
	Cfg               config.EnvironmentConfig
	IngredientService ingredientPort.IIngredientService
	Validator         bvalidator.Validator
	Logger            glog.Logger
}

func SetupDependencies() Dependency {
	validator := bvalidator.New()

	// init config
	config, err := config.LoadENVConfig()
	if err != nil {
		log.Panic(err)
	}

	// load logger
	logger := glog.New("info", "stdout")
	defer logger.Sync() // This script will be executed las`t
	defer logger.Info("Done cleanup tasks...")

	// BIG DEPENDENCY STAGE =======================================
	database, err := db.OpenPgsqlConnection(&config.Database)
	if err != nil {
		log.Panic(err)
	}

	// redis, err := db.RedisNewClient(&config.Cache)
	// if err != nil {
	// 	log.Panic(err)
	// }

	// BIG DEPENDENCY STAGE END =======================================

	// init repository
	ingredientRepository := ingredientRepo.NewIngredientRepository(database, keyTransaction, timeout)

	//init middleware

	// init service
	ingredientService := ingredientService.NewIngredientService(ingredientRepository, logger)

	return Dependency{
		Cfg:               config,
		IngredientService: ingredientService,
		Validator:         validator,
		Logger:            logger,
	}
}

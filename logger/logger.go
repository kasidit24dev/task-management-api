package logger

import (
	"go.uber.org/zap"
	"log"
	"task-management-api/config"
)

func InitLogger(cfg *config.Config) *zap.SugaredLogger {
	env := cfg.App.Env
	loggerCore, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("[ERROR] [InitLogger] init logger %s failed with error : %s", env, err.Error())
	}

	if env == "prod" {
		loggerCore, err = zap.NewProduction()
		if err != nil {
			log.Fatalf("[ERROR] [InitLogger] init logger %s failed with error : %s", env, err.Error())
		}
	}

	defer loggerCore.Sync()

	logger := loggerCore.Sugar()
	logger.Infof("logger init in %s", env)

	return logger

}

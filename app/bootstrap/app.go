package bootstrap

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Application struct {
	Config *Config
	DB     *gorm.DB
}

func App() (*Application, error) {
	app := &Application{}
	config, err := NewConfig(os.Getenv("GO_ENV"))
	if err != nil {
		return nil, err
	}
	app.Config = config
	db, err := gorm.Open(postgres.Open(config.GetDBUrl()), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	app.DB = db
	return app, nil
}

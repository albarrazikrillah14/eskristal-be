package pg

import (
	"fmt"
	"rania-eskristal/src/commons/config"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type pgConnection struct {
	Config *config.DBConfig
	Logger *logrus.Logger
}

func New(config *config.DBConfig, logger *logrus.Logger) *pgConnection {
	return &pgConnection{
		Config: config,
		Logger: logger,
	}
}

func (p *pgConnection) Connection() *gorm.DB {
	dsn := "host=%v user=%v password=%v dbname=%v port=%d sslmode=disable TimeZone=Asia/Shanghai"

	db, err := gorm.Open(postgres.Open(fmt.Sprintf(dsn, p.Config.Host, p.Config.User, p.Config.Password, p.Config.Name, p.Config.Port)), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		p.Logger.Error(fmt.Sprintf("database error %v", err))
	}

	return db
}

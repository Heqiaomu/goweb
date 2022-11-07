package mysql

import (
	"context"
	"github.com/sun-iot/goweb/config"
	"github.com/sun-iot/goweb/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"sync"
	"time"
)

var db *gorm.DB
var dbOnce sync.Once

func NewGormDB() error {
	m := config.GetConfig().Mysql
	if m.Dbname == "" {
		return global.ErrorNilDBName
	}
	mysqlConfig := mysql.Config{
		DSN:                       m.Dsn(), // DSN data source name
		DefaultStringSize:         191,     // string 类型字段的默认长度
		SkipInitializeWithVersion: false,   // 根据版本自动配置
	}
	var err error

	dbOnce.Do(func() {
		db, err = gorm.Open(mysql.New(mysqlConfig), GetConfig())
	})

	if err != nil {
		return nil
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(m.MaxIdleConns)
	sqlDB.SetMaxOpenConns(m.MaxOpenConns)
	return nil
}

type BASE interface {
	GetLogMode() string
}

func GetConfig() *gorm.Config {
	dbConfig := &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}
	_default := logger.New(NewWriter(log.New(os.Stdout, "\r\n", log.LstdFlags)), logger.Config{
		SlowThreshold: 200 * time.Millisecond,
		LogLevel:      logger.Warn,
		Colorful:      true,
	})
	logMode := &config.GetConfig().Mysql

	switch logMode.GetLogMode() {
	case "silent", "Silent":
		dbConfig.Logger = _default.LogMode(logger.Silent)
	case "error", "Error":
		dbConfig.Logger = _default.LogMode(logger.Error)
	case "warn", "Warn":
		dbConfig.Logger = _default.LogMode(logger.Warn)
	case "info", "Info":
		dbConfig.Logger = _default.LogMode(logger.Info)
	default:
		dbConfig.Logger = _default.LogMode(logger.Info)
	}
	return dbConfig
}
func GetDB(ctx context.Context) *gorm.DB {
	db = db.WithContext(ctx)
	return db
}

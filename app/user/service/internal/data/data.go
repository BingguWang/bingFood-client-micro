package data

import (
	"github.com/go-kratos/bingfood-client-micro/app/user/service/internal/conf"
    "gorm.io/gorm/logger"
    "os"
    "time"

    "github.com/go-kratos/kratos/v2/log"
    "github.com/google/wire"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    lg "log"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewDB, NewUserRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	db *gorm.DB
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, db *gorm.DB) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{db: db}, cleanup, nil
}

// *conf.Data就是数据的配置结构体
func NewDB(c *conf.Data) *gorm.DB {
	defer func() {
		if e := recover(); e != nil {
			lg.Printf("open mysql failed, err: %v", e)
		}
	}()

	newLogger := logger.New(
		lg.New(os.Stdout, "\r\n", lg.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,   // 慢 SQL 阈值
			LogLevel:      logger.Silent, // Log level
			Colorful:      true,          // 彩色打印
		},
	)
	//m := global.GVA_CONFIG.Mysql
	m := Mysql{
		Host:         "1.14.163.5",
		Port:         "3306",
		Username:     "root",
		Password:     "1234",
		Dbname:       "user",
		MaxIdleConns: 64,
		MaxOpenConns: 128,
	}

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       c.Database.Dsn, // DSN data source name
		DefaultStringSize:         191,            // string 类型字段的默认长度
		SkipInitializeWithVersion: true,           // 根据版本自动配置
	}), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		QueryFields:                              true,                           // 这样查询的时候是用字段名称而不是*
		Logger:                                   newLogger.LogMode(logger.Info), // 开启info级别的日志
	})
	if err != nil {
		panic(err)
	}

	sqlDB, _ := db.DB() // 取出成员SqlDB来配置
	sqlDB.SetMaxIdleConns(m.MaxIdleConns)
	sqlDB.SetMaxOpenConns(m.MaxOpenConns)

	lg.Printf("init mysql successful")

	return db
}

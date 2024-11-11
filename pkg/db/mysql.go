package db

import (
	"database/sql"
	"gin-django-example/pkg/log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func NewSql(conf *viper.Viper) *sql.DB {
	dsn := conf.GetString("his.db_url")
	myDB, err := sql.Open("mysql", dsn)
	myDB.Ping()
	if err != nil {
		panic(err)
	}
	return myDB
}

func NewDB(conf *viper.Viper, l *log.Logger) *gorm.DB {
	var (
		db  *gorm.DB
		err error
	)
	//logger := zapgorm2.New(l.Logger)
	dsn := conf.GetString("his.db_url")

	// GORM doc: https://gorm.io/docs/connecting_to_the_database.html
	//db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
	//	Logger: logger,
	//})
	db, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err)
	}
	db = db.Debug()

	// Connection Pool config
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	return db
}

package eye

import (
	"database/sql"
	"gin-django-example/pkg/log"
	"github.com/kataras/iris/v12/middleware/jwt/blocklist/redis"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type AppParam struct {
	Db  *gorm.DB
	Sql *sql.DB
	Env *viper.Viper
	Log *log.Logger
	Rds *redis.Client
}

package jwt

import (
	"github.com/spf13/viper"
)

var JwtKey string

func NewJwtKey(conf *viper.Viper) {
	JwtKey = conf.GetString("jwt.key")
}

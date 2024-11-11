package sentry

import (
	"fmt"
	"github.com/getsentry/sentry-go"
	"github.com/spf13/viper"
	"time"
)

var Client *sentry.Hub

func NewSentry(conf *viper.Viper) {
	DSN := conf.GetString("sentrydns")
	err := sentry.Init(sentry.ClientOptions{
		Dsn: DSN,
	})
	if err != nil {
		fmt.Println("sentry.Init: %s", err)
	}
	// Flush buffered events before the program terminates.
	defer sentry.Flush(2 * time.Second)
	Client = sentry.CurrentHub()
}

package sentry

import (
	"fmt"
	"os"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/sirupsen/logrus"

	"github.com/phanletrunghieu/pkg/logger"
)

func init() {
	sentryDSN := os.Getenv("SENTRY_DSN")
	env := os.Getenv("ENV")

	err := sentry.Init(sentry.ClientOptions{
		Dsn:              sentryDSN,
		Environment:      env,
		AttachStacktrace: true,
	})
	if err != nil {
		logger.NewLogger().Option(logger.WithLevel(logrus.ErrorLevel)).Log(fmt.Sprintf("sentry.Init: %s", err))

		return
	}

}

func Flush() {
	sentry.Flush(2 * time.Second)
}

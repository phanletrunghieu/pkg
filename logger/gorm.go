package logger

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
)

type GormLogger struct {
	logger                    *logrus.Logger
	level                     gormlogger.LogLevel
	IgnoreRecordNotFoundError bool
	SlowThreshold             time.Duration
}

func NewGorm() *GormLogger {
	l := logrus.New()
	l.SetFormatter(&logrus.JSONFormatter{})

	return &GormLogger{
		logger: l,
	}
}

func (l *GormLogger) LogMode(level gormlogger.LogLevel) gormlogger.Interface {
	l.level = level

	return l
}

func (l *GormLogger) Info(ctx context.Context, s string, args ...interface{}) {
	l.logger.WithContext(ctx).Infof(s, args...)
}

func (l *GormLogger) Warn(ctx context.Context, s string, args ...interface{}) {
	l.logger.WithContext(ctx).Warnf(s, args)
}

func (l *GormLogger) Error(ctx context.Context, s string, args ...interface{}) {
	l.logger.WithContext(ctx).Errorf(s, args)
}

func (l *GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.level <= gormlogger.Silent {
		return
	}

	logger := l.logger.WithContext(ctx)

	elapsed := time.Since(begin)
	switch {
	case err != nil && l.level >= gormlogger.Error && (!errors.Is(err, gorm.ErrRecordNotFound) || !l.IgnoreRecordNotFoundError):
		sql, rows := fc()
		if rows == -1 {
			logger.WithFields(logrus.Fields{
				"line":     utils.FileWithLineNum(),
				"sql":      sql,
				"duration": float64(elapsed.Nanoseconds()) / 1e6,
			}).Error(err)
		} else {
			logger.WithFields(logrus.Fields{
				"line":     utils.FileWithLineNum(),
				"sql":      sql,
				"rows":     rows,
				"duration": float64(elapsed.Nanoseconds()) / 1e6,
			}).Error(err)
		}
	case elapsed > l.SlowThreshold && l.SlowThreshold != 0 && l.level >= gormlogger.Warn:
		sql, rows := fc()
		slowLog := fmt.Sprintf("SLOW SQL >= %v", l.SlowThreshold)
		if rows == -1 {
			logger.WithFields(logrus.Fields{
				"line":     utils.FileWithLineNum(),
				"sql":      sql,
				"slowLog":  slowLog,
				"duration": float64(elapsed.Nanoseconds()) / 1e6,
			}).Warn(err)
		} else {
			logger.WithFields(logrus.Fields{
				"line":     utils.FileWithLineNum(),
				"sql":      sql,
				"slowLog":  slowLog,
				"rows":     rows,
				"duration": float64(elapsed.Nanoseconds()) / 1e6,
			}).Warn(err)
		}
	case l.level == gormlogger.Info:
		sql, rows := fc()
		if rows == -1 {
			logger.WithFields(logrus.Fields{
				"line":     utils.FileWithLineNum(),
				"sql":      sql,
				"duration": float64(elapsed.Nanoseconds()) / 1e6,
			}).Info(err)
		} else {
			logger.WithFields(logrus.Fields{
				"line":     utils.FileWithLineNum(),
				"sql":      sql,
				"rows":     rows,
				"duration": float64(elapsed.Nanoseconds()) / 1e6,
			}).Info(err)
		}
	}
}

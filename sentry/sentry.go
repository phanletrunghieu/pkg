package sentry

import (
	"github.com/getsentry/sentry-go"
)

type Sentry struct {
	// Identify Users
	userID    *string
	userEmail *string
	userName  *string

	level  sentry.Level
	fields map[string]string
}

type Option func(*Sentry)

func NewSentry() *Sentry {
	return &Sentry{
		level:  sentry.LevelError,
		fields: map[string]string{},
	}
}

func WithUserID(userID string) Option {
	return func(s *Sentry) {
		s.userID = &userID
	}
}

func WithUserEmail(userEmail string) Option {
	return func(s *Sentry) {
		s.userEmail = &userEmail
	}
}

func WithUserName(userName string) Option {
	return func(s *Sentry) {
		s.userName = &userName
	}
}

func WithLevel(level sentry.Level) Option {
	return func(s *Sentry) {
		s.level = level
	}
}

func WithFields(keyvals ...string) Option {
	return func(s *Sentry) {
		for i := 0; i < len(keyvals); i += 2 {
			if i+1 < len(keyvals) {
				s.fields[keyvals[i]] = keyvals[i+1]
			} else {
				s.fields[keyvals[i]] = "MISSING"
			}
		}
	}
}

func (s *Sentry) Option(options ...Option) *Sentry {
	for _, optFunc := range options {
		optFunc(s)
	}

	return s
}

func (s *Sentry) Log(err error) {
	localHub := sentry.CurrentHub().Clone()

	localHub.ConfigureScope(func(scope *sentry.Scope) {
		if s.userID != nil ||
			s.userEmail != nil ||
			s.userName != nil {
			user := sentry.User{}

			if s.userID != nil {
				user.ID = *s.userID
			}

			if s.userEmail != nil {
				user.Email = *s.userEmail
			}

			if s.userName != nil {
				user.Username = *s.userName
			}

			scope.SetUser(user)
		}

		scope.SetLevel(s.level)
		scope.SetTags(s.fields)
	})

	localHub.CaptureException(err)
}

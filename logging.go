package payment

import (
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

// LoggingMiddleware logs method calls, parameters, results, and elapsed time.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next Service) Service {
		return loggingMiddleware{
			next:   next,
			logger: logger,
		}
	}
}

type loggingMiddleware struct {
	next   Service
	logger log.Logger
}

func (mw loggingMiddleware) Authorise(amount float32) (auth Authorisation, err error) {
	defer func(begin time.Time) {
		level.Info(mw.logger).Log(
			"method", "Authorise",
			"result", auth.Authorised,
			"took", time.Since(begin),
		)
	}(time.Now())
	auth, err = mw.next.Authorise(amount)
	if err != nil {
		level.Error(mw.logger).Log(
			"method", "Authorise",
			"result", "error",
			"error", err.Error(),
		)
	}
	return auth, err
}

func (mw loggingMiddleware) Health() (health []Health) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "Health",
			"result", len(health),
			"took", time.Since(begin),
		)
	}(time.Now())
	return mw.next.Health()
}

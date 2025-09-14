package logger

import (
	"time"

	"github.com/valyala/fasthttp"
)

type Logger interface {
	Info(args ...interface{})
}

func NewBaseController(logger Logger) *BaseControler {
	return &BaseControler{logger: logger}
}

type BaseControler struct {
	logger Logger
}

func (c BaseControler) LoggingMiddleware(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		start := time.Now()
		next(ctx)
		duration := time.Since(start)
		c.logger.Info(
			"method:", string(ctx.Method()),
			"path:", string(ctx.Path()),
			"status:", ctx.Response.StatusCode(),
			"duration:", duration,
			"user-agent:", string(ctx.UserAgent()),
			"remote_ip:", ctx.RemoteIP().String(),
		)
	}
}

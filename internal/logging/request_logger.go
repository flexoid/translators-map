package logging

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
)

func NewRequestLogger(logger *zap.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			fields := []zap.Field{}
			if reqID := middleware.GetReqID(r.Context()); reqID != "" {
				fields = append(fields, zap.String("req_id", reqID))
			}

			l := logger.With(fields...)
			r = r.WithContext(WithCtx(r.Context(), l.Sugar()))
			next.ServeHTTP(w, r)
		}
		return middleware.RequestLogger(&RequestLogger{})(http.HandlerFunc(fn))
	}
}

type RequestLogger struct {
}

func (l *RequestLogger) NewLogEntry(r *http.Request) middleware.LogEntry {
	entry := &RequestLoggerEntry{Logger: Ctx(r.Context()).Desugar()}

	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}
	fields := []zap.Field{
		zap.String("http_scheme", scheme),
		zap.String("http_proto", r.Proto),
		zap.String("http_method", r.Method),

		zap.String("remote_addr", r.RemoteAddr),
		zap.String("user_agent", r.UserAgent()),

		zap.String("uri", fmt.Sprintf("%s://%s%s", scheme, r.Host, r.RequestURI)),
	}

	entry.Logger = entry.Logger.With(fields...)

	entry.Logger.Info("Request started")

	return entry
}

type RequestLoggerEntry struct {
	Logger *zap.Logger
}

func (l *RequestLoggerEntry) Write(status, bytes int, header http.Header, elapsed time.Duration, extra interface{}) {
	l.Logger = l.Logger.With(
		zap.Int("resp_status", status),
		zap.Int("resp_bytes_length", bytes),
		zap.Float64("resp_elapsed_ms", float64(elapsed.Nanoseconds())/1000000.0),
	)

	l.Logger.Info("Request complete")
}

func (l *RequestLoggerEntry) Panic(v interface{}, stack []byte) {
	l.Logger = l.Logger.With(
		zap.String("stack", string(stack)),
		zap.String("panic", fmt.Sprintf("%+v", v)),
	)
}

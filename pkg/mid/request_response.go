package mid

import (
	"fmt"
	"net/http"
	"time"

	"github.com/ezzycreative1/toko-kue-system/pkg/glog"
	"github.com/labstack/echo/v4"
)

type reqResTrap struct {
	logger glog.Logger
}

func NewReqResMiddleware(logger glog.Logger) *reqResTrap {
	return &reqResTrap{
		logger: logger,
	}
}

func (rr *reqResTrap) Middle(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		req := c.Request()
		res := c.Response()
		start := time.Now()

		// traceID
		traceID := GetID(c)

		// referer
		referer := getFirstValueFromHeader(req, "Referer")
		clientID := getFirstValueFromHeader(req, "X-Client-Id")
		// source ip
		ipAddress := getFirstValueFromHeader(req, "X-Forwarded-For")

		// LOG for incoming request
		rr.logger.Info("request started",
			glog.String("referer", referer),
			glog.String("client_id", clientID),
			glog.String("source_ip", ipAddress),
			glog.String("path", req.URL.Path),
			glog.String("trace_id", traceID),
		)

		// LOG after request end
		defer func() {
			rr.logger.Info("request completed",
				glog.String("referer", referer),
				glog.String("client_id", clientID),
				glog.String("source_ip", ipAddress),
				glog.String("path", req.URL.Path),
				glog.String("trace_id", traceID),
				glog.Any("latency", fmt.Sprintf("%d ms", time.Since(start).Milliseconds())),
				glog.Int("status", res.Status),
			)
		}()

		if err = next(c); err != nil {
			c.Error(err)
		}
		return
	}
}

func getFirstValueFromHeader(req *http.Request, key string) string {
	vs, ok := req.Header[key]
	if ok {
		if len(vs) != 0 {
			return vs[0]
		}
	}
	return ""
}

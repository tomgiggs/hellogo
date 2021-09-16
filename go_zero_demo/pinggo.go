package main

import (
	"flag"
	"fmt"
	"gitlab.deepwisdomai.com/infra/go-zero/rest/httpx"
	"net/http"
	"os"
	"pinggo/internal/middleware"
	"pinggo/internal/types"
	"time"

	"pinggo/internal/config"
	"pinggo/internal/handler"
	"pinggo/internal/svc"

	"gitlab.deepwisdomai.com/infra/go-zero/core/conf"
	"gitlab.deepwisdomai.com/infra/go-zero/rest"
)

var configFile = flag.String("f", "etc/pinggo-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	if env := os.Getenv("ENV"); env != "" {
		*configFile = "etc/pinggo-api-" + env + ".yaml"
	}
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()
	// 初始化Prometheus监控
	middleware.InitMonitor()
	server.Use(func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			// 统一校验权限
			if !middleware.CheckToken(ctx, w, r) {
				httpx.FailJson(w, types.ErrorAuthCheck)
				return
			}
			next(w, r)
		}
	})

	// prometheus统计
	server.Use(func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			middleware.WebRequestCount(r.Method, r.URL.Path)
			beginTime := time.Now()
			next(w, r)
			costTime := time.Since(beginTime)
			middleware.RequestCostTimeObserve(r.URL.Path, costTime.Seconds())
		}
	})
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

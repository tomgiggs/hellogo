package middleware

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"pinggo/internal/svc"
)

var (
	registry        *prometheus.Registry
	WebRequestTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "web_reqeust_total",
			Help: "Number of hello requests in total",
		},
		// 设置两个标签
		[]string{"method", "endpoint"},
	)
	RespTime = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name: "response_time",
			Help: "cost time per request",
		},
		[]string{"costTime"},
	)
)

func InitMonitor() {
	registry = prometheus.NewRegistry()
	registry.MustRegister(WebRequestTotal) // 注册指标
	registry.MustRegister(RespTime)        // 注册指标
}

func PrometheusHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h := promhttp.HandlerFor(registry,
			promhttp.HandlerOpts{
				EnableOpenMetrics: true,
			})
		h.ServeHTTP(w, r)
	}
}

func WebRequestCount(method string, url string) {
	if url != "/metrics" {
		WebRequestTotal.With(prometheus.Labels{"method": method, "endpoint": url}).Inc()
	}
}

func RequestCostTimeObserve(url string, costTime float64) {
	if url != "/metrics" {
		RespTime.WithLabelValues(url).Observe(costTime)
	}
}

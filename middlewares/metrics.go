package middlewares

import (
	"github.com/penglongli/gin-metrics/ginmetrics"
)

func Metrics() *ginmetrics.Monitor {
	m := ginmetrics.GetMonitor()
	m.SetMetricPath("/metrics")
	m.SetSlowTime(10)
	m.SetDuration([]float64{0.1, 0.3, 1.2, 5, 10})
	return m
}
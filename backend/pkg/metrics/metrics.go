package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)


var (
  TimeToProcessGetPrice = promauto.NewHistogram(prometheus.HistogramOpts{
    Name: "tb_get_price_request_duration",
    Help: "Time needed to process GET /api/v1/price request (price service endpoint), ms",
    Buckets: prometheus.LinearBuckets(0, 50, 20), // 0..1000
  })
)

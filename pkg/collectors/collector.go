package collectors

import (
	"github.com/prometheus/client_golang/prometheus"
)

// Collector interface defines the methods that platform-specific collectors must implement
type Collector interface {
	prometheus.Collector
	Name() string
}

// BaseCollector provides common functionality for all collectors
type BaseCollector struct {
	namespace string
	up        prometheus.Gauge
}

func NewBaseCollector(namespace string) BaseCollector {
	return BaseCollector{
		namespace: namespace,
		up: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "up",
			Help:      "Was the last collection successful",
		}),
	}
}

func (c *BaseCollector) Up() prometheus.Gauge {
	return c.up
}

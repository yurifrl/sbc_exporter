package rock

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/yurifrl/sbc_exporter/pkg/collectors"
)

type Collector struct {
	collectors.BaseCollector
	temperature *prometheus.GaugeVec
}

func NewCollector() *Collector {
	namespace := "sbc_rock"
	return &Collector{
		BaseCollector: collectors.NewBaseCollector(namespace),
		temperature: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Name:      "temperature",
				Help:      "Temperature of the SoC in degree celsius",
			},
			[]string{"zone"},
		),
	}
}

func (c *Collector) Name() string {
	return "rockchip"
}

func (c *Collector) Describe(ch chan<- *prometheus.Desc) {
	// Self-describing metrics
}

func (c *Collector) Collect(ch chan<- prometheus.Metric) {
	// Placeholder implementation - just set a dummy temperature
	c.temperature.WithLabelValues("soc").Set(45.0)
}

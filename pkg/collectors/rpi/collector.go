package rpi

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/yurifrl/sbc_exporter/pkg/collectors"
)

type Collector struct {
	collectors.BaseCollector
	temperature *prometheus.GaugeVec
	frequency   *prometheus.GaugeVec
	voltage     *prometheus.GaugeVec
	memory      *prometheus.GaugeVec
	vcgencmd    string
}

func NewCollector() *Collector {
	namespace := "sbc_rpi"
	return &Collector{
		BaseCollector: collectors.NewBaseCollector(namespace),
		temperature: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Name:      "temperature",
				Help:      "Temperatures of the components in degree celsius",
			},
			[]string{"sensor", "type"},
		),
		frequency: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Name:      "frequency",
				Help:      "Clock frequencies of the components in hertz",
			},
			[]string{"component"},
		),
		voltage: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Name:      "voltage",
				Help:      "Voltages of the components in volts",
			},
			[]string{"component"},
		),
		memory: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Name:      "memory",
				Help:      "Memory split of CPU and GPU in bytes",
			},
			[]string{"component"},
		),
		vcgencmd: "vcgencmd",
	}
}

func (c *Collector) Name() string {
	return "raspberry_pi"
}

// Ensure Collector implements collectors.Collector
func (c *Collector) Describe(ch chan<- *prometheus.Desc) {
	c.Up().Describe(ch)
	c.temperature.Describe(ch)
	// ... describe other metrics
}

func (c *Collector) Collect(ch chan<- prometheus.Metric) {
	c.Up().Set(1) // Set up metric to 1 when collection starts

	// If any error occurs during collection:
	// c.Up().Set(0)

	c.Up().Collect(ch)
	c.temperature.Collect(ch)
	// ... collect other metrics
}

// ... rest of the existing RPi collector methods ...

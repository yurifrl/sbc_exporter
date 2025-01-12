package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/cobra"
	"github.com/yurifrl/sbc_exporter/pkg/collectors"
	"github.com/yurifrl/sbc_exporter/pkg/collectors/rock"
	"github.com/yurifrl/sbc_exporter/pkg/collectors/rpi"
)

var listenAddress string

func detectBoardType() string {
	if _, err := os.Stat("/sys/class/thermal/thermal_zone0/temp"); err == nil {
		if _, err := exec.LookPath("vcgencmd"); err == nil {
			return "rpi"
		}
		return "rock"
	}
	return "unknown"
}

func startExporter(collector collectors.Collector) error {
	prometheus.MustRegister(collector)
	http.Handle("/metrics", promhttp.Handler())
	log.Printf("Starting SBC metrics exporter on %s", listenAddress)
	return http.ListenAndServe(listenAddress, nil)
}

var rootCmd = &cobra.Command{
	Use:   "sbc_exporter",
	Short: "Prometheus exporter for Single Board Computers",
	RunE: func(cmd *cobra.Command, args []string) error {
		board := detectBoardType()
		log.Printf("Auto-detected board type: %s", board)

		switch board {
		case "rpi":
			return startExporter(rpi.NewCollector())
		case "rock":
			return startExporter(rock.NewCollector())
		default:
			return fmt.Errorf("unable to detect board type")
		}
	},
}

var rpiCmd = &cobra.Command{
	Use:   "rpi",
	Short: "Run exporter for Raspberry Pi",
	RunE: func(cmd *cobra.Command, args []string) error {
		return startExporter(rpi.NewCollector())
	},
}

var rockCmd = &cobra.Command{
	Use:   "rock",
	Short: "Run exporter for Rockchip",
	RunE: func(cmd *cobra.Command, args []string) error {
		return startExporter(rock.NewCollector())
	},
}

func init() {
	rootCmd.PersistentFlags().StringVar(&listenAddress, "web.listen-address", ":9110", "Address on which to expose metrics")
	rootCmd.AddCommand(rpiCmd, rockCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

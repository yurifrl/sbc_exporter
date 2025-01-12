# sbc_exporter

A Prometheus exporter for Single Board Computers (SBC) metrics. Currently supports:
- Raspberry Pi
- Rockchip (basic metrics)

## Usage

The exporter exposes metrics on port 9110 at `/metrics`.

## Supported Metrics

### Raspberry Pi
- Temperature
- Frequency
- Voltage
- Memory split

### Rockchip
- Basic system metrics (placeholder for now)

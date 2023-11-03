package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/MichaelS11/go-dht"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	listenAddress = flag.String("listen-address", ":10005", "The address to listen on for HTTP requests.")
	gpioPort      = flag.Int64("gpio-port", 2, "The GPIO port where DHT22 is connected.")
	interval      = flag.Duration("interval", 60, "Temperature and humidity evaluation interval.")
)

type metrics struct {
	temperature prometheus.Gauge
	humidity    prometheus.Gauge
}

func NewMetrics(reg prometheus.Registerer) *metrics {
	m := &metrics{
		temperature: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: "dht22",
			Name:      "temperature_celsius",
			Help:      "Current temperature in Celsius",
		}),
		humidity: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: "dht22",
			Name:      "humidity_percent",
			Help:      "Current humidity in percent",
		}),
	}
	reg.MustRegister(m.temperature, m.humidity)
	return m
}

func Gather() (float64, float64) {
	err := dht.HostInit()
	if err != nil {
		log.Fatalf("error initializing dht, %v", err)
	}

	dht, err := dht.NewDHT(fmt.Sprintf("GPIO%d", *gpioPort), dht.Celsius, "")
	if err != nil {
		log.Fatalf("error connecting to dht, %v", err)
	}

	humidity, temperature, err := dht.ReadRetry(22)
	if err != nil {
		log.Printf("error reading temperature and humidity, %v", err)
	}

	return humidity, temperature
}

func main() {
	flag.Parse()

	reg := prometheus.NewRegistry()
	reg.MustRegister(collectors.NewGoCollector())
	m := NewMetrics(reg)

	go func() {
		for {
			humidity, temperature := Gather()
			m.temperature.Set(float64(temperature))
			m.humidity.Set(float64(humidity))

			time.Sleep(*interval * time.Second)
		}
	}()

	http.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{Registry: reg}))
	log.Fatal(http.ListenAndServe(*listenAddress, nil))
}

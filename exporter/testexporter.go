package exporter

import (
	"log"

	"github.com/Scout24/cgroup-metrics-reporter/collector"
)

type TestExporter struct{}

func (e *TestExporter) TestCounter(c collector.Statsd) bool {
	// prefix every metric with the app name
	err := c.Count("test", 10, nil, 1)
	if err != nil {
		log.Fatal(err)
	}

	return true
}

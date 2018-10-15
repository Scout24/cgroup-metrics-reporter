package collector

import (
	"log"
	"time"

	"github.com/DataDog/datadog-go/statsd"
)

type Statsd interface {
	Count(name string, value int64, tags []string, rate float64) error
}

// type CollectorFunc func(Statsd) bool
type Collector interface {
	// Register function
	// accepts a function that receives statsd agent and a parameter
	// then function can call Count or Gauge of other metric methods
	// on them
	Register(Exporter)

	Start()
}

type LoopCollector struct {
	exporters []Exporter
	statsd    Statsd
}

func NewLoopCollector(hostport string, namespace string, tags []string) Collector {
	s, err := statsd.New(hostport)
	s.Namespace = namespace
	s.Tags = tags

	if err != nil {
		log.Fatal(err)
	}
	return &LoopCollector{
		statsd: s,
	}
}

func (c *LoopCollector) Register(e Exporter) {
	c.exporters = append(c.exporters, e)
}

func (c *LoopCollector) Start() {
	ticker := time.NewTicker(1 * time.Second)

	go func() {
		for _ = range ticker.C {
			for _, e := range c.exporters {
				e.Export(c.statsd)
			}
		}
	}()
}

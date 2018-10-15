package main

import (
	"fmt"
	"net/http"

	"github.com/Scout24/cgroup-metrics-reporter/aws"
	"github.com/Scout24/cgroup-metrics-reporter/collector"
	"github.com/Scout24/cgroup-metrics-reporter/exporter"
)

func main() {
	config := LoadConfig()

	verboseLogging(config.Verbose)

	handler := aws.NewAWSHandler()

	c := collector.NewLoopCollector(
		config.StatsdCollectorAddress,
		config.Namesapce,
		[]string{collector.CreateTag("instance_id", handler.InstanceId)},
	)

	e := exporter.NewCGroupExporter(handler)

	c.Register(e)
	c.Start()

	http.HandleFunc("/health", checkHealth)
	fmt.Println("Listening on " + config.ListenAddress)
	if err := http.ListenAndServe(config.ListenAddress, nil); err != nil {
		panic(err)
	}
}

package exporter

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/Scout24/cgroup-metrics-reporter/aws"
	"github.com/Scout24/cgroup-metrics-reporter/collector"
)

type CGroupExporter struct {
	awsHandler *aws.AWSHandler
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

const (
	cpuStatsPath = "/cgroup/cpu/ecs/%s/cpu.stat"
)

func NewCGroupExporter(aws *aws.AWSHandler) *CGroupExporter {
	return &CGroupExporter{
		awsHandler: aws,
	}
}

func (e *CGroupExporter) parseCPUStats(path string) map[string]int64 {
	log.Printf("path: %s\n", path)
	cpuStats := make(map[string]int64)

	inFile, _ := os.Open(path)
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		key := parts[0]
		value, _ := strconv.ParseInt(parts[1], 10, 64)
		cpuStats[key] = value
	}
	return cpuStats
}

func (e *CGroupExporter) Export(c collector.Statsd) bool {
	for _, task := range e.awsHandler.GetEcsTasksOnInstance() {
		var tags []string
		path := fmt.Sprintf(cpuStatsPath, task[0])
		cpuStats := e.parseCPUStats(path)
		log.Printf("cpuStats: %s\n", cpuStats)

		tags = append(tags, collector.CreateTag("service_name", task[1]))
		tags = append(tags, collector.CreateTag("task_id", task[0]))
		log.Printf("tags: %s\n", tags)

		c.Count("nr_periods", cpuStats["nr_periods"], tags, 1)
		c.Count("nr_throttled", cpuStats["nr_throttled"], tags, 1)
		c.Count("throttled_time", cpuStats["throttled_time"], tags, 1)
	}

	return true
}

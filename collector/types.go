package collector

type Exporter interface {
	Export(Statsd) bool
}

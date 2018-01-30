package main

import (
	"fmt"

	zsend "github.com/blacked/go-zabbix"
)

func makePrefix(key string) string {
	return fmt.Sprintf(
		"ngs.srv.nginx.%s", key,
	)

}
func createZabbixMetrics(
	hostname string,
	stats *nginxStats,
) []*zsend.Metric {

	var metrics []*zsend.Metric

	metrics = append(
		metrics,
		zsend.NewMetric(
			hostname,
			makePrefix(
				"current.active.connections",
			),
			stats.currentActiveConnections,
		),
	)
	metrics = append(
		metrics,
		zsend.NewMetric(
			hostname,
			makePrefix(
				"total.accepted.connections",
			),
			stats.totalAcceptedConnections,
		),
	)
	metrics = append(
		metrics,
		zsend.NewMetric(
			hostname,
			makePrefix(
				"total.handled.connections",
			),
			stats.currentActiveConnections,
		),
	)
	metrics = append(
		metrics,
		zsend.NewMetric(
			hostname,
			makePrefix(
				"total.handles.requests",
			),
			stats.totalHandlesRequests,
		),
	)
	metrics = append(
		metrics,
		zsend.NewMetric(
			hostname,
			makePrefix(
				"current.reading.connections",
			),
			stats.currentNginxReadHeaderConnections,
		),
	)
	metrics = append(
		metrics,
		zsend.NewMetric(
			hostname,
			makePrefix(
				"current.writing.connections",
			),
			stats.currentNginxWriteToClientConnections,
		),
	)
	metrics = append(
		metrics,
		zsend.NewMetric(
			hostname,
			makePrefix(
				"current.waiting.connections",
			),
			stats.currentWaitingConnections,
		),
	)

	return metrics
}

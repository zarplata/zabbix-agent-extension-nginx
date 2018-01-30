package main

import (
	"fmt"
	"regexp"

	ser "github.com/reconquest/ser-go"
)

type nginxStats struct {
	currentActiveConnections             string
	totalAcceptedConnections             string
	totalHandledConnections              string
	totalHandlesRequests                 string
	currentNginxReadHeaderConnections    string
	currentNginxWriteToClientConnections string
	currentWaitingConnections            string
}

func parseNginxStatResponse(
	statResponse string,
) (*nginxStats, error) {
	findedMetricsLen := 8
	statRe, err := regexp.Compile(
		`\D+(\d+)\D+(\d+) (\d+) (\d+)\D+(\d+)\D+(\d+)\D+(\d+)`,
	)
	if err != nil {
		return nil, ser.Errorf(
			err,
			"unable to compile regexp pattern",
		)
	}

	findedStatMetrics := statRe.FindStringSubmatch(statResponse)
	if len(findedStatMetrics) != findedMetricsLen {
		return nil, fmt.Errorf(
			"count of finded metrics not equal with default metrics length %d",
			findedMetricsLen,
		)
	}

	return &nginxStats{
		currentActiveConnections:             findedStatMetrics[1],
		totalAcceptedConnections:             findedStatMetrics[2],
		totalHandledConnections:              findedStatMetrics[3],
		totalHandlesRequests:                 findedStatMetrics[4],
		currentNginxReadHeaderConnections:    findedStatMetrics[5],
		currentNginxWriteToClientConnections: findedStatMetrics[6],
		currentWaitingConnections:            findedStatMetrics[7],
	}, nil

}

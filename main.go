package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strconv"

	zsend "github.com/blacked/go-zabbix"
	docopt "github.com/docopt/docopt-go"
)

var (
	version string = "[manual build]"
)

func main() {
	usage := `zabbix-agent-extension-nginx

Usage:
    zabbix-agent-extension-nginx [options]

Options:
    -n --nginx <address>          Address where nginx listen [default: 127.0.0.1:80]
    -s --stat <location>          Nginx stat location  [default: nginx_stats]
    -z --zabbix-host <address>    Hostname or IP address of zabbix server [default: 127.0.0.1]
    -p --zabbix-port <number>     Port of zabbix server [default: 10051]
    -h                            Show this screen.
`

	args, err := docopt.Parse(usage, nil, true, version, false)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	zabbixHost := args["--zabbix-host"].(string)
	zabbixPort, err := strconv.Atoi(args["--zabbix-port"].(string))
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	statsURL := fmt.Sprintf(
		"http://%s",
		path.Join(
			args["--nginx"].(string),
			args["--stat"].(string),
		),
	)

	statsResponse, err := http.Get(statsURL)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	defer statsResponse.Body.Close()

	if statsResponse.StatusCode != http.StatusOK {
		fmt.Printf("bad status code %d", statsResponse.StatusCode)
		os.Exit(1)
	}

	body, err := ioutil.ReadAll(statsResponse.Body)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	stats, err := parseNginxStatResponse(string(body))
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	zabbixMetrics := createZabbixMetrics(
		hostname,
		stats,
	)

	packet := zsend.NewPacket(zabbixMetrics)
	sender := zsend.NewSender(
		zabbixHost,
		zabbixPort,
	)
	sender.Send(packet)
	fmt.Println("OK")
}

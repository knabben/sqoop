package main

import (
	"os"
	"github.com/solo-io/go-utils/log"
	"github.com/solo-io/sqoop/pkg/setup"
	"github.com/solo-io/go-utils/stats"
)


const (
	START_STATS_SERVER = "START_STATS_SERVER"
)

func main() {
	if os.Getenv(START_STATS_SERVER) != "" {
		stats.StartStatsServerWithPort(stats.StartupOptions{Port: 9091})
	}
	if err := setup.Main(); err != nil {
		log.Fatalf("err in main: %v", err.Error())
	}
}

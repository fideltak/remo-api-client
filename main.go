package main

import (
	"os"
	"time"

	remo "github.com/nature_remo_api_client/remo"
	utils "github.com/nature_remo_api_client/utils"
)

const (
	endpoint = "https://api.nature.global/1/devices"
)

var logging = &utils.Logging{}

func main() {
	logging.Infoln("Start nature remo API client")
	remo := &remo.Client{
		Endpoint:   endpoint,
		Token:      os.Getenv("REMO_TOKEN"),
		DeviceName: os.Getenv("REMO_TARGET_DEVICE"),
		LogPath:    os.Getenv("REMO_LOG_PATH"),
	}

	i, err := time.ParseDuration(os.Getenv("REMO_INTERVAL") + "s")
	if err != nil {
		logging.Fatalf("%v", err)
	}
	remo.Interval = i

	remo.RetrieveDataPeriodic()
}

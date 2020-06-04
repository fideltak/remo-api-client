package remo

import (
	"encoding/json"
	"time"

	utils "github.com/nature_remo_api_client/utils"
)

var logging = &utils.Logging{}

type Client struct {
	Endpoint   string
	Token      string
	DeviceName string
	Interval   time.Duration
	LogPath    string
}

type Data struct {
	Name   string `json:"name"`
	Events Events `json:"newest_events"`
}

type Events struct {
	Humidity    SensorData `json:"hu,omitempty"`
	Illuminance SensorData `json:"il,omitempty"`
	Motion      SensorData `json:"mo,omitempty"`
	Temperature SensorData `json:"te,omitempty"`
}

type SensorData struct {
	Value     json.Number `json:"val"`
	Timestamp string      `json:"created_at"`
}

package remo

import (
	"fmt"
	"os"
	"time"
)

func (c *Client) SaveData(data *Data) {
	file, err := os.OpenFile(c.LogPath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		logging.Fatalf("Could not open log file: %v", err)
	}
	defer file.Close()

	//If you want to change output device name, you can set custome name by setting OS env value.
	remo_name := c.DeviceName
	if os.Getenv("REMO_CUSTOM_NAME") != "" {
		remo_name = os.Getenv("REMO_CUSTOM_NAME")
	}
	
	fmt.Fprintf(file, "%v %v %v %v %v %v %v %v %v %v\n",
		time.Now().UTC().Format(time.RFC3339),
		remo_name,
		data.Events.Humidity.Timestamp,
		data.Events.Humidity.Value,
		data.Events.Illuminance.Timestamp,
		data.Events.Illuminance.Value,
		data.Events.Motion.Timestamp,
		data.Events.Motion.Value,
		data.Events.Temperature.Timestamp,
		data.Events.Temperature.Value)
}

package remo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func (c *Client) RetrieveData() (data *Data, err error) {
	req, _ := http.NewRequest("GET", c.Endpoint, nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", c.Token))
	req.Header.Set("accept", "application/json")
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err

	} else if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Could retrieve data from endpoint.: %v", resp.StatusCode)

	} else {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		var data_list []Data
		err = json.Unmarshal(body, &data_list)
		if err != nil {
			return nil, err
		}

		for _, data := range data_list {
			if data.Name == c.DeviceName {
				return &data, nil
			}
		}
	}
	return nil, fmt.Errorf("Cloud not found device: %v", c.DeviceName)
}

func (c *Client) RetrieveDataPeriodic() {
//	ticker := time.NewTicker(time.Second * c.Interval)
ticker := time.NewTicker(c.Interval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			logging.Infoln("Trying to retrieve remo data...")
			data, err := c.RetrieveData()
			if err != nil {
				logging.Errorf("%v", err)
			}

			//If you want to change output device name, you can set custome name by setting OS env value.
			remo_name := c.DeviceName
			if os.Getenv("REMO_CUSTOM_NAME") != "" {
				remo_name = os.Getenv("REMO_CUSTOM_NAME")
			}

			logging.Infof("%v DeviceName:%v H_Timestamp:%v Humidity:%v I_Timestamp:%v Illuminance:%v M_Timestamp:%v Motion:%v T_Timestamp: %v Temperature:%v\n",
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

			if c.LogPath !=""{
				c.SaveData(data)
			}
		}
	}
}

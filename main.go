package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	remo "github.com/nature_remo_api_client/remo"
	utils "github.com/nature_remo_api_client/utils"
)

const (
	endpoint = "https://api.nature.global/1/devices"
)

var logging = &utils.Logging{}

func main() {
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

	if strings.ToLower(os.Getenv("PROXY")) == "yes" {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			res, err := remo.RetrieveData()
			if err != nil{
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			res_json, err := json.Marshal(res)
			if err != nil{
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.Header().Add("Content-Type", "application/json")
			fmt.Fprintf(w, "%s\n", string(res_json))
		})

		proxy_port := "8080"
		if os.Getenv("PROXY_PORT") != "" {
			proxy_port = os.Getenv("PROXY_PORT")
		}
		go func() {
			logging.Fatal(http.ListenAndServe(":"+proxy_port, nil))
		}()
	}

	logging.Infoln("Start nature remo Docker API client")
	remo.RetrieveDataPeriodic()
}

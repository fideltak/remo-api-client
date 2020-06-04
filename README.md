# Nature Remo API Client
This is api client for Nature Remo.  
You can retrieve sensor data such as Humidity, Illuminance, Motion and Temperature. And simply deploy this  client to Docker, k8s or so.

## Test Environment
- Remo Firmware: 1.0.77-g808448c

## Environment Values
You can set prameters as envronment values.  

|Key|Value Example|Require|Description|
|:---|:---|:---|:---|
|REMO\_TOKEN|Guht5z......|Yes|The token issued from Nature API web site.|
|REMO\_TARGET\_DEVICE|MyRemo01|Yes|Your Nature Remo device name which you want retrieve data.|
|REMO\_INTERVAL|10|Yes|Iteration seconds.|
|REMO\_LOG\_PATH|/tmp/sensor.log|Option|Set log file path if you want to save sensor data.|
|REMO\_CUSTOM\_NAME|sensor01|Option|If you want to change your device name in stdout data or log file, set this parameter.|


## How to run
If you are using docker, run like below.  

```
docker run --env REMO_TOKEN=<YOUR TOKEN> --env REMO_TARGET_DEVICE=<YOUR DEVICE NAME> --env REMO_INTERVAL=10 --env REMO_LOG_PATH=/tmp/sensor.log fideltak/remo-api-client
```

If you want to save log file on your docker host, attach volume.

```
docker run --env REMO_TOKEN=<YOUR TOKEN> --env REMO_TARGET_DEVICE=<YOUR DEVICE NAME> --env REMO_INTERVAL=10 --env REMO_LOG_PATH=/tmp/sensor.log -v /tmp/:/tmp fideltak/remo-api-client
```
If you are k8s user...  

```
apiVersion: apps/v1
kind: Deployment
metadata:
  name: remo-api-client
  namespace: remo
  labels:
    app: remo-api-client
spec:
  replicas: 1
  selector:
    matchLabels:
      app: remo-api-client
  template:
    metadata:
      labels:
        app: remo-api-client
    spec:
      containers:
        - name: remo-api-client
          image: fideltak/remo-api-client:latest
          volumeMounts:
            - name: tmp
              mountPath: /tmp
          env:
            - name: REMO_TOKEN
              value: <YOUR TOKEN>
            - name: REMO_TARGET_DEVICE
              value: <YOUR DEVICE NAME>
            - name: REMO_INTERVAL
              value: '600'
            - name: REMO_LOG_PATH
              value: '/tmp/sensor01.log'
            - name: REMO_CUSTOM_NAME
              value: 'sensor01'
      volumes:
        - name: tmp
          hostPath:
            path: /tmp
```

## Example  
stdout  

```
[Info] 2020/06/04 12:47:50 Trying to retrieve remo data...
[Info] 2020/06/04 12:47:50 2020-06-04T03:47:50Z DeviceName:RemoLiving H_Timestamp:2020-06-04T03:41:04Z Humidity:65 I_Timestamp:2020-06-04T01:49:12Z Illuminance:242 M_Timestamp:2020-06-04T01:48:45Z Motion:1 T_Timestamp: 2020-06-04T00:45:47Z Temperature:25.79
[Info] 2020/06/04 12:48:00 Trying to retrieve remo data...
[Info] 2020/06/04 12:48:00 2020-06-04T03:48:00Z DeviceName:RemoLiving H_Timestamp:2020-06-04T03:41:04Z Humidity:65 I_Timestamp:2020-06-04T01:49:12Z Illuminance:242 M_Timestamp:2020-06-04T01:48:45Z Motion:1 T_Timestamp: 2020-06-04T00:45:47Z Temperature:25.79
[Info] 2020/06/04 12:48:10 Trying to retrieve remo data...
[Info] 2020/06/04 12:48:10 2020-06-04T03:48:10Z DeviceName:RemoLiving H_Timestamp:2020-06-04T03:41:04Z Humidity:65 I_Timestamp:2020-06-04T01:49:12Z Illuminance:242 M_Timestamp:2020-06-04T01:48:45Z Motion:1 T_Timestamp: 2020-06-04T00:45:47Z Temperature:25.79
[Info] 2020/06/04 12:48:20 Trying to retrieve remo data...
[Info] 2020/06/04 12:48:20 2020-06-04T03:48:20Z DeviceName:RemoLiving H_Timestamp:2020-06-04T03:41:04Z Humidity:65 I_Timestamp:2020-06-04T01:49:12Z Illuminance:242 M_Timestamp:2020-06-04T01:48:45Z Motion:1 T_Timestamp: 2020-06-04T00:45:47Z Temperature:25.79
[Info] 2020/06/04 12:48:30 Trying to retrieve remo data...
[Info] 2020/06/04 12:48:30 2020-06-04T03:48:30Z DeviceName:RemoLiving H_Timestamp:2020-06-04T03:41:04Z Humidity:65 I_Timestamp:2020-06-04T01:49:12Z Illuminance:242 M_Timestamp:2020-06-04T01:48:45Z Motion:1 T_Timestamp: 2020-06-04T00:45:47Z Temperature:25.79
```  

log file

```
2020-06-04T03:39:01Z RemoLiving 2020-06-04T02:40:04Z 66 2020-06-04T01:49:12Z 242 2020-06-04T01:48:45Z 1 2020-06-04T00:45:47Z 25.79
2020-06-04T03:39:11Z RemoLiving 2020-06-04T02:40:04Z 66 2020-06-04T01:49:12Z 242 2020-06-04T01:48:45Z 1 2020-06-04T00:45:47Z 25.79
2020-06-04T03:39:21Z RemoLiving 2020-06-04T02:40:04Z 66 2020-06-04T01:49:12Z 242 2020-06-04T01:48:45Z 1 2020-06-04T00:45:47Z 25.79
2020-06-04T03:39:31Z RemoLiving 2020-06-04T02:40:04Z 66 2020-06-04T01:49:12Z 242 2020-06-04T01:48:45Z 1 2020-06-04T00:45:47Z 25.79
2020-06-04T03:39:41Z RemoLiving 2020-06-04T02:40:04Z 66 2020-06-04T01:49:12Z 242 2020-06-04T01:48:45Z 1 2020-06-04T00:45:47Z 25.79
2020-06-04T03:39:51Z RemoLiving 2020-06-04T02:40:04Z 66 2020-06-04T01:49:12Z 242 2020-06-04T01:48:45Z 1 2020-06-04T00:45:47Z 25.79
2020-06-04T03:40:01Z RemoLiving 2020-06-04T02:40:04Z 66 2020-06-04T01:49:12Z 242 2020-06-04T01:48:45Z 1 2020-06-04T00:45:47Z 25.79
2020-06-04T03:40:11Z RemoLiving 2020-06-04T02:40:04Z 66 2020-06-04T01:49:12Z 242 2020-06-04T01:48:45Z 1 2020-06-04T00:45:47Z 25.79
2020-06-04T03:40:21Z RemoLiving 2020-06-04T02:40:04Z 66 2020-06-04T01:49:12Z 242 2020-06-04T01:48:45Z 1 2020-06-04T00:45:47Z 25.79
2020-06-04T03:40:31Z RemoLiving 2020-06-04T02:40:04Z 66 2020-06-04T01:49:12Z 242 2020-06-04T01:48:45Z 1 2020-06-04T00:45:47Z 25.79
2020-06-04T03:40:41Z RemoLiving 2020-06-04T02:40:04Z 66 2020-06-04T01:49:12Z 242 2020-06-04T01:48:45Z 1 2020-06-04T00:45:47Z 25.79
2020-06-04T03:40:51Z RemoLiving 2020-06-04T02:40:04Z 66 2020-06-04T01:49:12Z 242 2020-06-04T01:48:45Z 1 2020-06-04T00:45:47Z 25.79
2020-06-04T03:41:01Z RemoLiving 2020-06-04T02:40:04Z 66 2020-06-04T01:49:12Z 242 2020-06-04T01:48:45Z 1 2020-06-04T00:45:47Z 25.79
2020-06-04T03:41:11Z RemoLiving 2020-06-04T03:41:04Z 65 2020-06-04T01:49:12Z 242 2020-06-04T01:48:45Z 1 2020-06-04T00:45:47Z 25.79
2020-06-04T03:41:21Z RemoLiving 2020-06-04T03:41:04Z 65 2020-06-04T01:49:12Z 242 2020-06-04T01:48:45Z 1 2020-06-04T00:45:47Z 25.79
2020-06-04T03:41:31Z RemoLiving 2020-06-04T03:41:04Z 65 2020-06-04T01:49:12Z 242 2020-06-04T01:48:45Z 1 2020-06-04T00:45:47Z 25.79
2020-06-04T03:41:41Z RemoLiving 2020-06-04T03:41:04Z 65 2020-06-04T01:49:12Z 242 2020-06-04T01:48:45Z 1 2020-06-04T00:45:47Z 25.79
```
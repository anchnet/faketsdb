# faketsdb
Forward data to InfluxDB with fake OpenTSDB protocol in Open-Falcon

## Usage
```sh
go get github.com/51idc/faketsdb
```

## Open-Falcon Transfer
// cfg.json
```
{
    //.....
    "tsdb": {
        "enabled": false,   # don't forget turn enable
        "batch": 200,
        "connTimeout": 1000,
        "callTimeout": 5000,
        "maxConns": 32,
        "maxIdle": 32,
        "retry": 3,
        "address": "127.0.0.1:8088"   # modify here 8089(default)
    }
}
```

## Helper
```sh
🍺 eagle [~] → faketsdb -h
Usage of faketsdb:
  -cache int
    	Number of batch items send to influx. (default 3)
  -influxAddr string
    	InfluxDB HTTP API address. (default "http://127.0.0.1:8086")
  -influxDatabase string
    	InfluxDB Database. (default "test")
  -port int
    	Fake proxy listen port. (default 8089)
 ```
 
 ## Daemon
 
 Recommended Supervisor 

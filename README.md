# Lightweight Metric Exporter

## Storing example

```bash
POST / HTTP/1.1
Accept: application/json, */*
Accept-Encoding: gzip, deflate
Connection: keep-alive
Content-Length: 160
Content-Type: application/json
Host: localhost:8080
User-Agent: HTTPie/1.0.3

{
    "date": 12312314,
    "description": "Crazy counter metric",
    "labels": {
        "asiyas_asa": "dfg",
        "label_2": "my"
    },
    "name": "metric_3",
    "type": "counter",
    "value": 1.3
}

HTTP/1.1 200 OK
Content-Length: 46
Content-Type: text/plain; charset=utf-8
Date: Thu, 03 Dec 2020 12:46:30 GMT

{
    "message": "Metric was updated!",
    "status": 200
}
```

## Expose metrics for your Prometheus
```bash
GET /metrics HTTP/1.1
Accept: */*
Accept-Encoding: gzip, deflate
Connection: keep-alive
Host: 127.0.0.1:8080
User-Agent: HTTPie/1.0.3



HTTP/1.1 200 OK
Content-Length: 228
Content-Type: text/plain; charset=utf-8
Date: Thu, 03 Dec 2020 12:54:35 GMT

# HELP metric_3_total Crazy counter metric
# TYPE metric_3_total counter
metric_3_total{asiyas_asa="dfg",label_2="my"} 6.500000
# HELP metric_3 Crazy metric
# TYPE metric_3 gauge
metric_3{asiyas_asa="dfg",label_2="my"} 1.300000
```

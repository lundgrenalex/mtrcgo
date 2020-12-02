# Lightweight Metric Exporter

## Storing example

```bash
POST / HTTP/1.1
Accept: application/json, */*
Accept-Encoding: gzip, deflate
Connection: keep-alive
Content-Length: 86
Content-Type: application/json
Host: localhost:8080
User-Agent: HTTPie/1.0.3

{
    "date": 12312314,
    "labels": {
        "asiyas-asa": "dfg"
    },
    "name": "ccssdd_sd",
    "value": 1.3
}

HTTP/1.1 200 OK
Content-Length: 46
Content-Type: text/plain; charset=utf-8
Date: Wed, 02 Dec 2020 10:13:38 GMT

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
Content-Length: 107
Content-Type: text/plain; charset=utf-8
Date: Wed, 02 Dec 2020 10:15:16 GMT

ccssdd_sd{asiyas-asa=dfg} 1.300000
```

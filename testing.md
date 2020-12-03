## Gauge-test

http POST localhost:8080 name="metric_3" labels:='{"asiyas_asa": "dfg", "label_2": "my"}' date:=12312314 value:=1.3 type="gauge" description="Crazy gauge metric" --verbose
http POST localhost:8080 name="metric_3" labels:='{"asiyas_asa": "dfg", "label_2": "my"}' date:=12312314 value:=1.3 type="counter" description="Crazy counter metric" --verbose

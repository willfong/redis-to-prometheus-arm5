# redis-to-prometheus

A simple Go app to serve as an HTTP endpoint for Prometheus metrics.

```
docker run \
    -d --name redis_to_prometheus \
    -p 8080:8080 \
    --net aqua \
    --restart unless-stopped \
    wfong/redis-to-prometheus
```


# Prometheus

## Metric Types
4 core metric types
differentiated in client library
    - not used by the prometheus server

### Counter
- monotonically increasing counter
- increase or reset to 0
- Uses
    - requests served
    - tasks completed
    - errors

### Gauge
- numeric value can go up and down
- Uses
    - temperature
    - memory usage
    - number of processes
    - number of concurrent requests

## Histogram
- sample observations (request durations or responses sizes etc.)
- count them in configurable buckets
- base metric <basename>
    - exposes multiple time series during a scrape
    - cummulative counters for the observation buckets
        - <basename>_bucket{le="<upper inclusive bound>"}
        - total sum for all observed values <basename>_sum
        - count of events <basename>_count
- histogram_quantile()
    - calculate quantiles from histogram
    - aggregations of histograms
    - apdex score
    - it is cummulative unlike summaries

## Summary
- similar to histogram does sampling
- calculates configurable amounts (sliding window)
    - not cumulative over entire time like histogram
- Metrics
    - phi-quantiles exposed as <basename>{quantile="<phi>"} where 0 <= phi <=1
    - total sum for all observed values <basename>_sum
    - count of events <basename>_count

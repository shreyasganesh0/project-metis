# Structured Logging

## JSON and KV logs

### Defining goals
- logging for development, debugging or BI
- affects flow and format of the logs
    - field names
    - no need to monitor exceptions if we need log for metrics
### Deciding what to log
- logging too much could cause
    - noise with unimportant data
    - expensive logging storage costs
- have to decide what you want to log and what to ignore
    - for example:
    - building alerts on top of logs - log actionalble data

### Selecting framework
- gives a mechanism to implement logging
- verbosity, define log levels, log rotation policy
- multiple connections or targets
- choose frameworks that are feature rich and easy to use
- also consider performance on the application

### Standardize logs
- understand when to use each log level
- create standard formatting and naming fields
- choose between json or kv pairs

### Formatting
- JSON
{
    "@timestamp": "2025-07025 17:03:12",
    "level": error,
    "message": "messaging details",
    "service": "service_name",
    "ip": "34.124.233.12",
}

- KVP
2025-07025 17:03:12 level=error message="messaging details"
service="service_name" ip=34.124.233.12

- choose based on analysis tool

### Provide context
- logs should be concise but clear
- adding context like category=cat_value instead of just having cat_val

### Unique id
- adding tags and ids where porrible makes navigations easy
- different systems being ided allows you to follow through any errors across apps
- transaction ids user id and account id


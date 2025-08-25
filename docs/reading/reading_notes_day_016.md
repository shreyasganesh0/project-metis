# Opentelemetry

## Traces
- path of request through an application
- represented as Spans JSON
```
{
    "name": "hello",
    "context": {
        "trace_id": "46723ac2342cd3424",
        "span_id": "423324234",
    },
    "parent_id": null,
    "start_time": "2022-04-29T18:53.23.1142013",
    "end_time": "2022-04-29T18:53.23.1142013",
    "attributes": {
        "http.route": "some_route1"
    },
    "events": {
        {
            "name": "hello world",
            "timestamp": "2022-04-29T18:43.58.1145021",
            "attributes": {
                "event_attributes": 1
            }
        }
    }
}
```
    - this is root trace has no parent id
```
{
  "name": "hello-greetings",
  "context": {
    "trace_id": "5b8aa5a2d2c872e8321cf37308d69df2",
    "span_id": "5fb397be34d26b51"
  },
  "parent_id": "051581bf3cb55c13",
  "start_time": "2022-04-29T18:52:58.114304Z",
  "end_time": "2022-04-29T22:52:58.114561Z",
  "attributes": {
    "http.route": "some_route2"
  },
  "events": [
    {
      "name": "hey there!",
      "timestamp": "2022-04-29T18:52:58.114561Z",
      "attributes": {
        "event_attributes": 1
      }
    },
    {
      "name": "bye now!",
      "timestamp": "2022-04-29T18:52:58.114585Z",
      "attributes": {
        "event_attributes": 1
      }
    }
  ]
}
```
    - span for specific tasks
    - parent is the hello (root) span and parent_id matches span_id of parent
    - same trace_id as teh root span

```
{
  "name": "hello-salutations",
  "context": {
    "trace_id": "5b8aa5a2d2c872e8321cf37308d69df2",
    "span_id": "93564f51e1abe1c2"
  },
  "parent_id": "051581bf3cb55c13",
  "start_time": "2022-04-29T18:52:58.114492Z",
  "end_time": "2022-04-29T18:52:58.114631Z",
  "attributes": {
    "http.route": "some_route3"
  },
  "events": [
    {
      "name": "hey there!",
      "timestamp": "2022-04-29T18:52:58.114561Z",
      "attributes": {
        "event_attributes": 1
      }
    }
  ]
}
```
    - child of hello span and sibiling of hello-greeting span
    - share same trace_id and parent_id represents the hierarchy

- spans are kind of like structured logs
- has context correlation, hierarchy
- can come from different services vms data services
    - This is the key to allow tracing to present the end-to-end view of a system
- Tracer Provider
    - factory for tracers
    - lifecycle matches the lifecycle of the application
    - tracer provider init includes resources and exporter init and is the first step
    in tracing with OpenTelemetry
- Tracer
    - creates spans containing info about the operation like a request
- Trace Exporters
    - sends traces to a consumer
        - consumers can be stdout, OpenTelemetry Collector or any open source backend
- Conext Propagation
    - spans can be correlated to each other and assembled into a trace
- Spans
    - unit of work/operation
    - Include information
        - Name
        - Parent span ID
        - Start and End timestamps
        - Span Context
        - Attributes
        - Span Events
        - Span Links
        - Span Status
- Span Context
    - immutable object on every span
    - contains
        - Trace ID
        - span's Span ID
        - Trace flags
        - Trace state: kv pairs with vendor specific trace info
    - part of a span that is propagated alongside Distributed Context and Baggage
    - used when creating Span Links
- Attributes
    - kv pairs with metadata used to annotate a span for info about the operation
    - capture user id and cart id for an operation that adds item to cart
    - add during creation so SDK sampling can use the values
    - keys must be non null string
    - values must be non null string boolean floating point int or an array of those
    - semantic attributes which are known naming conventions for metadata in
    common operations
- Span Events
    - like a strucutred log message on a Span
    - used to denote a single point during the spans duration
        - denoting when a page becomes interactive
    - spans events vs span attributes
        - when a operation completes and you want to add some data from the
        operation to the telemetry
        - if timestamp during compeletion is meaningful - use span events
        - use span attributes if the timestamp isnt relevant

- Span Links
    - associate open span with one or more spans
    - usually usefule when one operation spawns other asyn operations
    - to trace both we associate two seperate traces using a span link
        - link last span from the first trace to the first span of the second trace
- Span Status
    - 3 possible values
        - Unset
            - default value
            - completed without error
        - Error
            - some error occured in the operation it tracks
        - Ok
            - span was marked explicitlly as error-free
            - useful to finalize interpretation of the span
- Span Kind
    - 5 Kinds
        - Client
            - sync outgoing remote call
            - not queued later for processing
        - Server
            - sync incoming remote call
        - Internal
            - operations that dont cross process boundary
        - Producer
            - creation of a async job
        - Consumer
            - processing of a job created by a producer
    - provides a hint to the backend of trace assembly
        - parent of server span is remote client span
        - child of client span is usually server span
        - parent of a consumer span is the producer
        - child of the producer span is the consumer span
    - if not provided it is assumed to be "internal"


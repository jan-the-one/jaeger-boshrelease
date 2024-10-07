# jaeger-boshrelease

## TODO 

- Connect Jaeger to Kafka instead of Cassandra (docker compose) -> https://stackoverflow.com/q/64391505 [x]
- Add prebuilt binaries to /src [x]
- Make a simple grpcClient/httpClient for Jaeger [x] (easier to use HTTP based clients)
- Write job scripts [x]
- Configure a Concourse setup [x]
- Add Jaeger Dashboard [x] 
- Test out & fix bindings using a connector App [x]
- - Curling healthcheck endpoints [x]
- - Adding some spans (using some custom-made app) [x]
- Add BOSH properties []
- - The env variables to wire everything need to be made parametric!
- - In particular, the "storage" configuration is important
- - Must test out the TLS setup for Cassandra. 
- - Kafka needs the "ingester" setup and is thus not supported at this time
- Add Makefile to build the release and manage the deployment []
- - Add some target to open tunnels, or some `.dsf` script
- Add Concourse automation to produce Jaeger binaries on-demand and prepare some PR []
- Focus on graphite Exports of Traces + Metrics !! []
----------- ----------- ----------- ----------- ----------- ----------- ----------- -----------

- Explore NGINX-based tracing []
- Explore OpenTelemetry relations to Prometheus/Graphite []

- Design the "processor" app as a simple trace pusher
- * In particular, add properties to specify details of Jaeger Query (dashboard), and external storages like Cassandra

---

- We don't need a co-located "Kafka" cluster but we can benefit from an output stream to somewhere
- - We need to implement some "mock" storage which can then make the traces available somehow for consumption
- In a test deployment we can try to have a single VM in which the Jaeger setup is fully running. 
- To test it properly we would need to add a consul agent and run it accordingly.
- **NOTE**: We need to try the custom gRPC server in a local setup with Docker. Once it is ready we can try to include it in the Boshrelease

- SAMPLING API @ Jaeger Collector: https://www.jaegertracing.io/docs/1.60/sampling/#remote-sampling
- TRACE API @ Jaeger Query https://www.jaegertracing.io/docs/1.60/apis/#grpcprotobuf-stable
- STORAGE API @ Internal component addressed by Jaeger Collector: https://www.jaegertracing.io/docs/1.60/apis/#remote-storage-api-stable
  - Jaeger Remote Storage is an in-memory solution for traces
  - High degree of customization
  - Tutorial @ https://github.com/jaegertracing/jaeger/blob/main/plugin/storage/grpc/README.md

---

### Sampling Strategies & Extensions of Jaeger

https://www.jaegertracing.io/docs/1.61/sampling/

- Native Connection to NGINX? [MUST DO]
- Translation into metrics, e.g. Graphite format? [MUST DO]

- TODO: Add a custom storage-handler which can map traces to _metrics_. It can define custom metrics and make them available in Graphite format.
- TODO: Create the Graphite mappings and streams
- TODO: Create a small client that reads traces from the agent API.


### Components

- ~~Pre-processor AGENT (sessionings)~~
- Discovery / Policy / Instigator components for Fault Injection

#### ~~Session-based Tracing~~

- Open a "Trace" at an app that keeps some local storage (remote sessions)
- Add "metrics" to it from the service; must have some localized "agent" to do such monitoring
- - Number of Traces
- - CPU / MEM increase
- - The trace events should serve as some probe-time delineator; e.g. telling the instrumentation code 
    when to record events and so on. 


#### Split-Trace Fault-Injection & Stress Testing

- Test some existing traced path
- Test using policies
- For each "span" we need some endpoint to hit or way of identifying a node
- Or we traverse the trace and kill each span-node 

#### Tracing-based fault Injection

- For each span, add some faulty code as baggage; 
- "Traverse" the trace again
- There has to be some way of performing the assigned Fault;
  The span-reporting code must be able to invoke some callback
  * Null Pointer Exceptions
  * Paging issues (huge files)
  * Bad Kernel invokations
  * Lockspinning

#### Budget Based Sampling && Distance Based Sampling

::: Summary :::

The `pusher` app should be able to import synthetic traces, and implement `hot-spot` based tracing 
based on some "distance" metric. It could also rely on some "budget", or "depth" limit

A -> B         --        ------> C -> F -> G
       -> E ->   -> D[x] ->
                  -> J[x]       

- First step would be to enforce service-level traces only !!
- Example: the J span is ommited eagerly. That's because the depth might be too much.
- [post-processing; metrics; budget] Example: The D and J spans are ommitted due to the parent services not having been traversed
  often enough. The more a service is frequented, the larger it's budget to spend on "child" spans. Budget could also be set "manually" for each new Trace.
- [post-processing; metrics] Example: [Requires a fully-known trace] Once the entire trace exists, spans are omitted based on some "distance", e.g. nr of hops from specific hotspots. Nr. of hops would be a network mesh 
  parameter, not a span-based one.

# jaeger-boshrelease

## TODO 

- Connect Jaeger to Kafka instead of Cassandra (docker compose) -> https://stackoverflow.com/q/64391505 [x]
- Add prebuilt binaries to /src [x]
- Make a simple grpcClient/httpClient for Jaeger [x] (easier to use HTTP based clients)
- Add the grpcClient to Jaeger as a package []
- Write job scripts []
- Add properties for the configs of each component []


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

- TODO: Add a custom storage-handler which can map traces to _metrics_. It can define custom metrics and make them available in Graphite format.
- TODO: Create the Graphite mappings and streams
- TODO: Create a small client that reads traces from the agent API.

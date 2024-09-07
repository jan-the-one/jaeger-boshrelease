
docker run --rm -v/Users/gerdalliu/Workspace/jaeger:/jaeger -w/jaeger \
    jaegertracing/protobuf:0.2.0 "-I/jaeger -Iidl/proto/api_v2 -I/usr/include/github.com/gogo/protobuf -Iplugin/storage/grpc/proto --go_out=/jaeger/code plugin/storage/grpc/proto/storage.proto"

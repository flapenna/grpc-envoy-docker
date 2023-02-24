# grpc-envoy-docker

## Setup using Docker containers

The project is designed to use docker compose and run two containers:

* Golang service
* Envoy

## How to start

### Compile the proto 
Before running the application we have to compile the [proto file user_management.proto](proto/user_management.proto).\
To do that run the script [generate_api.sh](generate_api.sh) located under the root.\
Make sure you have installed protoc on your machine, otherwise check the documentation https://grpc.io/docs/protoc-installation/ .

### Run the application
Start docker containers : `docker-compose up` \
Stop docker containers  : `docker-compose down`

### Test it
Now you should have the gRPC service running and serving at 9091 port, and envoy listening at port 9090. \
To test the application you can use Postman, curl or any other http client to make http request at port 9090.

The Envoy filter allows a RESTful JSON API client to send requests to Envoy over HTTP and get proxied to a gRPC service.

[Envoy proxy documentation](https://www.envoyproxy.io/docs/envoy/latest/configuration/http/http_filters/grpc_json_transcoder_filter)

#### Bonus
Keep in my mind the port 9091 isn't mapped in the docker-compose.yaml, so you won't be able to hit the gRPC server from a network outside the Docker one.\
To hit the service with gRPC request make sure the port is mapped in the docker-compose.yaml.

```
services:
  grpcservice:
    build:
      context: .
      dockerfile: Dockerfile-service
    expose:
      - "9091"
```

to 

```
services:
  grpcservice:
    build:
      context: .
      dockerfile: Dockerfile-service
    ports:
      - "9091:9091"
```



## Authors
- [Federico La Penna](https://github.com/flapenna)

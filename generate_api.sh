#!/usr/bin/env bash

protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
  --include_imports --include_source_info --descriptor_set_out=./pb/user_service_definition.pb \
  --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
  proto/*.proto

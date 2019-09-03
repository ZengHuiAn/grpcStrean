#!/usr/bin/env bash


protoc  --go_out=plugins=grpc:. MessagePack.proto ProtoPack.proto ServiceProto.proto ServiceRegist.proto
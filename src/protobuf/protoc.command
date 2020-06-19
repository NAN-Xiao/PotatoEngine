#!/bin/bash
cd $(dirname $BASH_SOURCE) || {
    echo Error getting script directory >&2
    exit 1
}
#ls $mydir

#outPath="$(pwd)"
#echo $(pwd)
file= "$(pwd)/test.proto"
protoc --go_out=. *.proto

go:generate protoc --go_out=plugins=grpc:../routeguide ../routeguide/route_guide.proto protoc
protoc --proto_path=src --go_out=build/gen src/foo.proto src/bar/baz.proto
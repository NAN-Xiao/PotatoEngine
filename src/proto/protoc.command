#!/bin/bash
cd $(dirname $BASH_SOURCE) || {
    echo Error getting script directory >&2
    exit 1
}
#ls $mydir
#file= "$(pwd)/test.proto"
protoc --proto_path=./ --go_out=./ --go_opt=paths=source_relative *.proto
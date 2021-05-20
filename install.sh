#!/bin/bash
while getopts d: OPT; do
    case "$OPT" in
      d)
        DIR="$OPTARG" ;;
      [?])
        # got invalid option
        echo "Usage: $0 [-d work directory]" >&2
        exit 1 ;;
    esac
done

go test ./...
go get ./...

mkdir -p $DIR
go build -o $DIR/ts

cp service.yaml $DIR/

mkdir -p $DIR/data/mapping
mkdir -p $DIR/data/ontology
mkdir -p $DIR/data/result/report
mkdir -p $DIR/data/result/sent
mkdir -p $DIR/data/source/inprogress
mkdir -p $DIR/data/source/processed

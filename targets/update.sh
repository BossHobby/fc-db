#! /bin/bash

rm json/*.json
pushd json
go run ../../cmd/convert/main.go ../../unified-targets/configs/default/
popd

rm -rf header/*
pushd header
go run ../../cmd/header/main.go ../json/
popd
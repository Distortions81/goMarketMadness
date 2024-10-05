#!/bin/bash

GOOS=js GOMAXPROCS=1 GOARCH=wasm go build -trimpath -gcflags=all="-B" -ldflags="-s -w" -o start.wasm

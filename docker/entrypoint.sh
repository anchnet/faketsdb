#!/usr/bin/env bash

if [ ! $PORT ]; then
    PORT=8089
fi

if [ ! $CACHE ]; then
    CACHE=200
fi

if [ ! $INFLUXDB_ADDR ]; then
    INFLUXDB_ADDR=http://influxdb:8086
fi

if [ ! $INFLUXDB_DATABASE ]; then
    INFLUXDB_DATABASE=test
fi


if [ ! $DEBUG ]; then
    DEBUG='-debug'
fi

faketsdb \
    $DEBUG \
    -port=$PORT \
    -cache=$CACHE \
    -influxAddr=$INFLUXDB_ADDR \
    -influxDatabase=$INFLUXDB_DATABASE
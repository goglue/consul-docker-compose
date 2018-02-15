#!/usr/bin/env bash

SUBNET=10.0.0.0/16   # my defaults were always 172.18, using 19 only to test that this works
GATEWAY=10.0.0.4
docker network create --subnet=$SUBNET --gateway $GATEWAY \
  -o com.docker.network.bridge.name=docker_gwbridge \
  -o com.docker.network.bridge.enable_icc=true \
  -o com.docker.network.bridge.enable_ip_masquerade=true \
  docker_gwbridge
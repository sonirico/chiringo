#!/usr/bin/env bash

for index in $(seq 2000 2100)
do
  echo http post localhost:2001/peers host=localhost port="$index"
  http post localhost:2001/peers host=localhost port="$index"
  sleep 1;
done
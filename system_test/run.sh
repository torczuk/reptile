#!/bin/bash

wait_for_port() {
  local port=$1
  local status=1

  while [ $status -ne 0 ]; do
    echo "wait for localhost:2600"
    sleep 1
    nc -z localhost 2600
    status=$(echo $?)
  done
}
docker-compose build
docker-compose up -d

go test
readonly result=$?

docker-compose down

exit $result
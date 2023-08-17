#!/usr/bin/env bash

while true; do
  for pod in $(kubectl get pods -l app=hello_app -o jsonpath="{.items[*].metadata.name}")
  do
    vol=$(kubectl get pod $pod -o jsonpath="{..image}" | sort | sed -e 's/:latest//' | awk '{print $1}')
    echo "$pod: $vol"
  done
  sleep 3
  echo "----"
done
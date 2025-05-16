#!/bin/bash

set -eu

dir="$(dirname "$0")"

FLY="${FLY_CLI:-fly}"

"$FLY" -t "${CONCOURSE_TARGET:-bosh-ecosystem}" \
  sp -p os-conf-release \
  -c "$dir/pipeline.yml"


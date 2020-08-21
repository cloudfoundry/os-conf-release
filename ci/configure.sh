#!/bin/bash

set -exu

dir="$(dirname "$0")"

FLY="${FLY_CLI:-fly}"

"$FLY" -t "${CONCOURSE_TARGET:-production}" \
  sp -p os-conf-release \
  -c "$dir/pipeline.yml" \
  -l <(lpass show --notes 'os-conf-release pipeline vars') \


#!/bin/bash

set -exuo pipefail

export GOPATH="${PWD}/os-conf-release"
export PATH="${GOPATH}/bin:$PATH"

apt-get update
apt-get install -y netcat-openbsd

jumpbox_url=${BOSH_JUMPBOX_URL:-${BOSH_JUMPBOX_IP}:22}
jumpbox_private_key_path=$(mktemp)
chmod 600 ${jumpbox_private_key_path}
echo "${BOSH_JUMPBOX_PRIVATE_KEY}" > ${jumpbox_private_key_path}

export BOSH_ALL_PROXY=ssh+socks5://${BOSH_JUMPBOX_USER}@${jumpbox_url}?private-key=${jumpbox_private_key_path}

bosh upload-stemcell ${PWD}/stemcell/*.tgz
bosh create-release --dir ${PWD}/os-conf-release --timestamp-version --tarball=release.tgz
bosh upload-release release.tgz

export BOSH_BINARY_PATH=$(which bosh)

pushd "${PWD}/os-conf-release/src/os-conf-acceptance-tests"
  go run github.com/onsi/ginkgo/ginkgo/v2 -v
popd

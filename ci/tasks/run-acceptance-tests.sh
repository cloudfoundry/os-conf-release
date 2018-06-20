#!/bin/bash

set -exuo pipefail

export GOPATH="${PWD}/os-conf-release"
export PATH="${GOPATH}/bin:$PATH"

BBL_VERSION=6.7.10
BBL_SHA=53bd673f2967b5b82a3db4862c823230f55fe7d2f57bcec8241318e4ba9d8e60
curl -fSL https://github.com/cloudfoundry/bosh-bootloader/releases/download/v${BBL_VERSION}/bbl-v${BBL_VERISON}_linux_x86-64 -o /usr/bin/bbl \
  && echo "$BBL_SHA  /usr/bin/bbl" | shasum -c - \
  && chmod +x /usr/bin/bbl

eval "$(bbl --state-dir "${PWD}/bbl-state" print-env)"

BOSH_VERSION=4.0.1
BOSH_SHA=0045668aa6582ddd518cfb2b6134eba439bbb94d
curl -fSL https://s3.amazonaws.com/bosh-cli-artifacts/bosh-cli-${BOSH_VERSION}-linux-amd64 -o /usr/bin/bosh \
  && echo "$BOSH_SHA  /usr/bin/bosh" | shasum -c - \
  && chmod +x /usr/bin/bosh

bosh upload-stemcell ${PWD}/stemcell/*.tgz
bosh create-release --dir ${PWD}/os-conf-release --timestamp-version --tarball=release.tgz
bosh upload-release release.tgz

export BOSH_BINARY_PATH=/usr/bin/bosh

pushd "${PWD}/os-conf-release/src/os-conf-acceptance-tests"
  go install vendor/github.com/onsi/ginkgo/gingko
  ginkgo -v
popd

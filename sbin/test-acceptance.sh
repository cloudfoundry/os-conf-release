#!/usr/bin/env bash

set -e

BASEDIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )/.." && pwd )"

export BOSH_BINARY_PATH=${BOSH_BINARY_PATH:-/usr/local/bin/bosh}
export BOSH_DEPLOYMENT=${BOSH_DEPLOYMENT:-os-conf-xenial}
export BOSH_STEMCELL=${BOSH_STEMCELL:-ubuntu-xenial}

if [ -z ${BBL_STATE_DIR} ]; then
  echo "'BBL_STATE_DIR' not set. You need to specify the path to the 'bbl-state.json' of a working bbl environment."
  exit 1
fi

if [ -z ${STEMCELL} ]; then
  echo "'STEMCELL' not set. You need to specify the path to a stemcell of type '${BOSH_STEMCELL}'."
  exit 1
fi


eval "$(bbl --state-dir "${BBL_STATE_DIR}" print-env)"

bosh upload-stemcell $STEMCELL
bosh create-release --dir ${BASEDIR} --timestamp-version --tarball=release.tgz --force
bosh upload-release release.tgz


pushd "${BASEDIR}/src/os-conf-acceptance-tests"
  go install ./vendor/github.com/onsi/ginkgo/ginkgo
  ginkgo -v
popd

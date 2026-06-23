#!/usr/bin/env bash
set -eu -o pipefail

REPO_ROOT="$( cd "$( dirname "${BASH_SOURCE[0]}" )/../.." && pwd )"
REPO_PARENT="$( cd "${REPO_ROOT}/.." && pwd )"

if [[ -n "${DEBUG:-}" ]]; then
  set -x
  export BOSH_LOG_LEVEL=debug
  export BOSH_LOG_PATH="${BOSH_LOG_PATH:-${REPO_PARENT}/bosh-debug.log}"
fi

pushd bbl-state
  # Run bbl plan to initialise the state directory (downloads bosh-deployment
  # from GitHub as a side-effect). Then replace the auto-downloaded copy with
  # the pinned bosh-deployment resource so we get reproducible director builds.
  bbl plan

  rm -rf bosh-deployment
  cp -r "${REPO_PARENT}/bosh-deployment" .

  bbl --debug up
popd

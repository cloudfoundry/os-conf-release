#!/usr/bin/env bash
set -eu -o pipefail

REPO_ROOT="$( cd "$( dirname "${BASH_SOURCE[0]}" )/../.." && pwd )"
REPO_PARENT="$( cd "${REPO_ROOT}/.." && pwd )"

if [[ -n "${DEBUG:-}" ]]; then
  set -x
  export BOSH_LOG_LEVEL=debug
  export BOSH_LOG_PATH="${BOSH_LOG_PATH:-${REPO_PARENT}/bosh-debug.log}"
fi

apt-get update
apt-get install -y netcat-openbsd

if [ -n "${BOSH_JUMPBOX_PRIVATE_KEY}" ]; then
  jumpbox_url=${BOSH_JUMPBOX_URL:-${BOSH_JUMPBOX_IP}:22}
  jumpbox_private_key_path=$(mktemp)
  chmod 600 "${jumpbox_private_key_path}"
  echo "${BOSH_JUMPBOX_PRIVATE_KEY}" > "${jumpbox_private_key_path}"

  export BOSH_ALL_PROXY="ssh+socks5://${BOSH_JUMPBOX_USER}@${jumpbox_url}?private-key=${jumpbox_private_key_path}"
fi

bosh_cli=$(which bosh)

"${bosh_cli}" upload-stemcell "${REPO_PARENT}/stemcell/"*.tgz
"${bosh_cli}" create-release --dir "${REPO_ROOT}" --timestamp-version --tarball=release.tgz
"${bosh_cli}" upload-release release.tgz

pushd "${REPO_ROOT}/src/os-conf-acceptance-tests"
  export BOSH_BINARY_PATH="${bosh_cli}"
  go run github.com/onsi/ginkgo/v2/ginkgo -v
popd

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
  if [[ ! -f bbl-state.json ]]; then
    echo "No bbl-state.json found; bbl up never completed. Nothing to tear down."
    exit 0
  fi

  echo "Pre-cleanup: deleting any remaining BOSH deployments before tearing down the environment..."
  bbl_env="$(bbl print-env 2>/dev/null || true)"
  if [[ -n "${bbl_env}" ]]; then
    eval "${bbl_env}" || true
    deployments="$(timeout 60 bosh deployments --json 2>/dev/null \
      | jq -r '.Tables[0].Rows[].name' 2>/dev/null || true)"
    if [[ -n "${deployments}" ]]; then
      echo "${deployments}" | while IFS= read -r dep; do
        echo "Pre-cleanup: deleting deployment '${dep}'..."
        timeout 600 bosh -d "${dep}" delete-deployment --force -n 2>&1 \
          || echo "WARNING: failed to delete deployment '${dep}'"
      done
    else
      echo "No BOSH deployments found; skipping pre-cleanup."
    fi
  else
    echo "WARNING: could not source BOSH env; skipping deployment pre-cleanup."
  fi

  bbl --debug down --no-confirm
popd

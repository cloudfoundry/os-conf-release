#!/usr/bin/env bash

set -e

# Version info
semver_version="$(cat os-conf-version/number)"
echo "$semver_version" > released-version/semver-version
echo "v$semver_version" > released-version/prefixed-semver-version

pushd os-conf-release
  echo "${PRIVATE_YML}" > config/private.yml

  bosh create-release --final --version ${semver_version}

  rm config/private.yml

  git diff | cat
  git add .

  git config --global user.email "ci@localhost"
  git config --global user.name "CI Bot"
  git commit -m "OS Conf Release v${semver_version}"
popd

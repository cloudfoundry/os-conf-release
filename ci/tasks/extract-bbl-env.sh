#!/usr/bin/env bash
set -euo pipefail

pushd bbl-state
  eval "$(bbl print-env)"
  JUMPBOX_PRIVATE_KEY_CONTENT="$(cat "${JUMPBOX_PRIVATE_KEY:-${BOSH_GW_PRIVATE_KEY}}")"
popd

# bbl print-env sets BOSH_GW_HOST (jumpbox IP or hostname) and BOSH_GW_USER.
JUMPBOX_IP="${BOSH_GW_HOST:-$(echo "${BOSH_ALL_PROXY}" | cut -d"@" -f2 | cut -d":" -f1)}"
JUMPBOX_USER="${BOSH_GW_USER:-jumpbox}"

cat > bosh-env/metadata.yml << EOF
INSTANCE_JUMPBOX_PRIVATE: |-
$(echo "${JUMPBOX_PRIVATE_KEY_CONTENT}" | sed -E 's/(-+(BEGIN|END) (RSA |EC |OPENSSH )?PRIVATE KEY-+) *| +/\1\n/g' | sed 's/^/  /')
INSTANCE_JUMPBOX_USER: "${JUMPBOX_USER}"
INSTANCE_JUMPBOX_EXTERNAL_IP: "${JUMPBOX_IP}"
BOSH_CLIENT: "${BOSH_CLIENT}"
BOSH_CLIENT_SECRET: "${BOSH_CLIENT_SECRET}"
BOSH_ENVIRONMENT: "${BOSH_ENVIRONMENT}"
BOSH_CA_CERT: |-
$(echo "${BOSH_CA_CERT}" | sed -E 's/(-+(BEGIN|END) CERTIFICATE-+) *| +/\1\n/g' | sed 's/^/  /')
EOF

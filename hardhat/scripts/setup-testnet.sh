#!/bin/bash
set -e

# Install the Avalanche CLI
# curl -sSfL https://raw.githubusercontent.com/ava-labs/avalanche-cli/main/scripts/install.sh | sh -s
# export PATH=~/bin:$PATH
avalanche subnet import file infra/ava/conf/lilypad.json
cp infra/ava/key/* ~/.avalanche-cli/key
avalanche subnet deploy lilypad -l
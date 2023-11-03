#!/bin/bash

set -e
sudo chown -Rh vscode:vscode /workspaces/resource-provider-api/.devcontainer/nsc

echo "Dumping NATS user creds file"
nsc --data-dir=/workspaces/resource-provider-api/.devcontainer/nsc/nats/nsc/stores generate creds -a RESP -n USER > /tmp/user.creds

echo "Dumping NATS sys creds file"
nsc --data-dir=/workspaces/resource-provider-api/.devcontainer/nsc/nats/nsc/stores generate creds -a SYS -n sys > /tmp/sys.creds

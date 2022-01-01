#!/bin/sh

set -e

. tfvars.sh

terraform destroy -auto-approve

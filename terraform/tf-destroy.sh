#!/bin/bash

set -e

cd cluster-config
terraform destroy -auto-approve
rm config || echo "config not found, ignoring..."

cd ../cluster-init
terraform destroy --auto-approve
rm config || echo "config not found, ignoring..."

cd ../digitalocean
terraform destroy --auto-approve

cd ..

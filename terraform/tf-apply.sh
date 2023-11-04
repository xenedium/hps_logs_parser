#!/bin/bash

# These ENV VARS are required
# TF_VAR_ingress_prod_host
# TF_VAR_ingress_staging_host
# TF_VAR_acme_email
# TF_VAR_do_token
# S3_bucket
# S3_access_key
# S3_secret_key

set -e

cd digitalocean
terraform init -backend-config="bucket=$S3_bucket" -backend-config="access_key=$S3_access_key" -backend-config="secret_key=$S3_secret_key"
terraform apply -auto-approve

cd ..
cp digitalocean/config cluster-config/config
cp digitalocean/config cluster-init/config

cd cluster-init
terraform init -backend-config="bucket=$S3_bucket" -backend-config="access_key=$S3_access_key" -backend-config="secret_key=$S3_secret_key"
terraform apply -auto-approve

cd ../cluster-config
terraform init -backend-config="bucket=$S3_bucket" -backend-config="access_key=$S3_access_key" -backend-config="secret_key=$S3_secret_key"
terraform apply -auto-approve

cd ..

# key-value-store
Key value store service

## Prerequisites
* Route53 Hosted Zone
* SSL Cert ARN
* EC2 key pair Name
* Pipeline Source bucket (versioning enabled)
* Pipeline bucket (versioning enabled)
* artifact bucket with consul binary

## Build, create pipeline, Deploy infrastructure and test
./scripts/build_parameter_files.sh
./scripts/setup_and_run_deployment_pipeline.sh


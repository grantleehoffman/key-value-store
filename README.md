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
```
./scripts/build_parameter_files.sh
./scripts/setup_and_run_deployment_pipeline.sh
```
## Teardown

Stacks must be tore down in the following order and each stack should be fully destroyed before destroying the next stack.
```
aws cloudformation delete-stack --stack-name consul
aws cloudformation delete-stack --stack-name persistent-resources
aws cloudformation delete-stack --stack-name kv-pipeline

```


# key-value-store
Key value store service

## Prerequisites
* dep Golang Dependency manager
* awscli
* Route53 Hosted Zone
* SSL Cert ARN
* EC2 key pair name
* Pipeline source bucket (versioning enabled)
* Pipeline action bucket (versioning enabled)
* Artifact bucket with unzipped [consul binary](https://releases.hashicorp.com/consul/1.1.0/consul_1.1.0_linux_amd64.zip)

## Build, create pipeline, Deploy infrastructure and test
Currently must be run in us-east-1 due to ami hardcoded
```
./scripts/setup_and_run_deployment_pipeline.sh -s pipeline-source-bucket -a codepipeline-bucket -p my-profile -r us-east-1 -c
```
## Teardown

Stacks must be tore down in the following order and each stack should be fully destroyed before destroying the next stack.
```
aws cloudformation delete-stack --stack-name kv-consul
aws cloudformation delete-stack --stack-name kv-persistent-resources
aws cloudformation delete-stack --stack-name kv-pipeline

```

## TODO
ami id region map in cfn

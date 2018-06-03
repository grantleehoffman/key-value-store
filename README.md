# key-value-store
Key value store service

## Prerequisites
* dep Golang dependency manager
* awscli
* Route53 hosted zone
* SSL cert ARN
* EC2 key pair name
* Pipeline source bucket (versioning enabled)
* Pipeline action bucket (versioning enabled)
* Artifact bucket with unzipped [consul binary](https://releases.hashicorp.com/consul/1.1.0/consul_1.1.0_linux_amd64.zip)

## Build, create pipeline, Deploy infrastructure and test
must be run in us-east-1 due to ami hardcoding
```
./scripts/setup_and_run_deployment_pipeline.sh -s pipeline-source-bucket -a codepipeline-bucket -p my-profile -r us-east-1 -c
```
## Teardown

Stacks must be tore down in the following order and each stack should be fully destroyed before destroying the next stack.
```
aws cloudformation delete-stack --stack-name kv-consul-cluster --profile my-profile --region us-east-1
aws cloudformation delete-stack --stack-name kv-persistent-resources --profile my-profile --region us-east-1
aws cloudformation delete-stack --stack-name kv-pipeline --profile my-profile --region us-east-1

```

## CLI

#### Build
```
build.sh
```

Mac OSX binary can be located at:
```
./cli/bin/darwin/key-value
```

Linux binary can be located at:
```
./cli/bin/linux/key-value
```

## TODO
ami id region map in cfn

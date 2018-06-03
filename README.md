# key-value-store
Key value store service

## Prerequisites
* dep Golang dependency manager
* awscli

#### AWS Account Setup Steps
1. Setup consul artifact bucket
2. Download [consul zip](https://releases.hashicorp.com/consul/1.1.0/consul_1.1.0_linux_amd64.zip) extract `consul` binary and upload to artifact bucket
3. Setup pipeline source bucket, enable versioning
4. Setup pipeline action bucket, enable versioning
5. Create EC2 Key Pair
6. Register Route53 Domain
7. Create SSL Cert through AWS ACM for the registered domain in step 4


## Build, create pipeline, deploy infrastructure and test
* must run in us-east-1 due to ami hardcoding
```
./scripts/setup_and_run_deployment_pipeline.sh -s pipeline-source-bucket -a codepipeline-bucket -p my-profile -r us-east-1 -c
```

* You will be promped for the resource names from the account setup steps

## Teardown

Stacks must be tore down in the following order, each stack should be fully destroyed before deleting the following stack otherwise you may get resource dependency conflicts.

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
* Front Consul cluster with custom API
* Add Multi-AZ

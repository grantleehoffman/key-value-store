# key-value-store
Key value store service

## Prerequisites
* go 1.10.x
* dep Golang dependency manager
* awscli

#### AWS Account Setup Steps
1. Setup consul artifact bucket
2. Download [consul zip](https://releases.hashicorp.com/consul/1.1.0/consul_1.1.0_linux_amd64.zip) extract `consul` binary and upload to artifact bucket
3. Setup pipeline source bucket, enable versioning
4. Setup pipeline action bucket, enable versioning
5. Create EC2 Key Pair
6. Register Route53 Domain
7. Create SSL Cert through AWS ACM for the registered domain in the last step

* Resources must all be created in the same region
  * tested in us-east-1 & us-west-2


## Build, create pipeline, deploy infrastructure and test

```
./scripts/setup_and_run_deployment_pipeline.sh -s pipeline-source-bucket -a codepipeline-bucket -p my-profile -r us-east-1 -c
```

* You will be promped for the resource names from the account setup steps

## Teardown

Stacks must be tore down in the following order, each stack should be fully destroyed before deleting the following stack otherwise you may receive resource dependency errors.

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

#### Integration test script

A basic integration test scipt is located at `scripts/key-value-test.sh`

* This is the same script the Codebuild test job runs.

This script requires you have the key-value binary in your current working directory and set the env `server_url` with your server URI.

* Example: `export server_url="kvdemo.thehoff.xyz"`


## TODO
* Front Consul cluster with custom API
* Add Multi-AZ

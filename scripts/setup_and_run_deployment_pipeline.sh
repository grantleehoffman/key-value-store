#!/bin/bash -e
./build.sh
export AWS_DEFAULT_REGION=us-east-1
aws s3 cp deployment-artifact.zip s3://craft-demo-pipeline-source-bucket-us-east-1/deployment-artifact.zip
aws cloudformation create-stack --stack-name kv-pipeline --template-body file://cfn/pipeline.json --capabilities CAPABILITY_IAM

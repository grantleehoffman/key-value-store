#!/bin/bash -e
../build.sh
AWS_DEFAULT_REGION=us-west-1
aws s3 cp deployment-artifact.zip s3://craft-demo-pipeline-source-bucket-us-east-1/deployment-artifact.zip
aws cloudformation create-stack --template-body file://../cfn/pipeline.json

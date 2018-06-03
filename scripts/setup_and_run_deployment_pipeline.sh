#!/bin/bash -e

usage() {
  echo -e "\nUsage:"
  echo "  -s source_bucket [required]"
  echo "  -a pipeline_action_bucket [required]"
  echo "  -p profile [required]"
  echo "  -r region [required]"
  echo "  -c [optional] - interactively set up your cfn parameter values"
  echo -e "\nExample: $0 -s my-s3-bucket -a my-pipeline-bucket -p my-aws-profile -r us-east-1 -c\n"
  exit 1
}

while getopts "s:a:p:r:c" OPT
do
  case $OPT in
    "s")
      source_bucket=${OPTARG}
      ;;
    "a")
      pipeline_action_bucket=${OPTARG}
      ;;
    "p")
      profile=${OPTARG}
      ;;
    "r")
      region=${OPTARG}
      ;;
    "c")
      cfn_parameters=true
      ;;
    "h"|*)
      usage
      ;;
  esac
done

if [[ -z ${source_bucket} || -z ${profile} || -z ${region} || -z ${pipeline_action_bucket} ]]; then
  usage
fi

if [[ ${cfn_parameters} ]];then
  printf "===== Build CFN Parameters =====\n"
  scripts/build_parameter_files.sh
  printf "================================\n\n\n"
fi

set +e
stack_status=$(aws cloudformation describe-stacks --stack-name kv-pipeline --profile "${profile}" --region "${region}" --query Stacks[].StackStatus --output text 2> /dev/null)
set -e

printf "=== Build Deployment Artifact ====\n"
./build.sh
printf "==================================\n\n\n"

echo "==== Upload Deployment Artifact ===="
aws s3 cp deployment-artifact.zip "s3://${source_bucket}/deployment-artifact.zip" --profile "${profile}"
printf "==================================\n\n\n"

if [[ ${stack_status} != "CREATE_COMPLETE" ]] && [[ ${stack_status} != "UPDATE_COMPLETE" ]];then
printf "==== Pipeline Stack not found ====\n"
printf "==== Creating pipeline stack =====\n"
  aws cloudformation create-stack --stack-name kv-pipeline --template-body file://cfn/pipeline.json \
     --capabilities CAPABILITY_IAM --region "${region}" --parameters \
     "[{\"ParameterKey\":\"SourceArtifactBucket\",\"ParameterValue\":\"${source_bucket}\"},\
     {\"ParameterKey\":\"SourceArtifactKey\",\"ParameterValue\":\"deployment-artifact.zip\"},\
     {\"ParameterKey\":\"PipelineBucket\",\"ParameterValue\":\"${pipeline_action_bucket}\"}]"
printf "==================================\n\n\n"
fi
echo "Tasks complete, new pipeline revision queued"


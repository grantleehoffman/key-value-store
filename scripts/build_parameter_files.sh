#!/bin/bash -e
echo "Please provide the following values to create the cfn parameters"

get_args() {
while [ "$input" == "" ]
do
   read -p "EC2 Key Pair name to login to stacks: " input
   KeyName=("$input")
   read -p "The s3 bucket that contains the 'consul' binary: " input
   ArtifactBucket=("$input")
   read -p "Cidr block allowed to access key value api, use 0.0.0.0/0 if testing from codebuild: " input
   AllowedIpCidrBlock=("$input")
   read -p "Route53 registered domain name: " input
   ELBHostedZoneName=("$input")
   read -p "A SSL Cert ARN to terminate SSL on the ELB: " input
   SSLCertARN=("$input")
done
}

get_args
echo "$KeyName $ArtifactBucket $AllowedIpCidrBlock $ELBHostedZoneName $SSLCertARN"

echo "{\"Parameters\" : {\"KeyName\" : \"${KeyName}\", \"ArtifactBucket\": \"$ArtifactBucket\" } }" > cfn/consul-parameters.json
echo "{\"Parameters\" : {\"ELBHostedZoneName\" : \"${ELBHostedZoneName}\", \"SSLCertARN\": \"$SSLCertARN\", \"AllowedIpCidrBlock\": \"$AllowedIpCidrBlock\" } }" > cfn/persistent-resources-parameters.json

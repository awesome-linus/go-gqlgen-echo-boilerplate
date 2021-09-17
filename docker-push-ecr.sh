#!/usr/bin/env bash

if [ "$1" != "dev" ] && [ "$1" != "qa" ] && [ "$1" != "stg" ] && [ "$1" != "prod" ]
then
  echo  "Invalid argument! Please set one of 'qa' or 'dev' or 'stg' or 'prd'."
  exit 1
fi

if [ "$2" = "" ]; then
  echo  "Invalid argument! Please set tag."
  exit 1
fi

if [ "${AWS_ACCOUNT_ID}" = "" ]; then
  echo  "Invalid Environment variable! Please set Your AWS_ACCOUNT_ID."
  exit 1
fi

deployStage="$1"
repositoryName="$2"
imageTag="$3"

repositoryUri="${AWS_ACCOUNT_ID}.dkr.ecr.ap-northeast-1.amazonaws.com/${deployStage}-${repositoryName}"

$(aws ecr get-login --no-include-email --region ap-northeast-1 --profile awesome-terraform)

docker build --no-cache --rm -f docker/prd/go/Dockerfile -t "${repositoryUri}" .
docker tag "${repositoryUri}":latest "${repositoryUri}":"${imageTag}"
docker push "${repositoryUri}":latest
docker push "${repositoryUri}":"${imageTag}"

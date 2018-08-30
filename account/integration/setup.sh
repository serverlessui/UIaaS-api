#!/bin/bash

set -e

##define inputs
STACK=$1
COLLECTION=$2

##define variables
MACOS="Mac"

NEWURL=`aws cloudformation describe-stacks \
            --stack-name $STACK \
            --query "Stacks[0].[Outputs[? starts_with(OutputKey, 'ServiceEndpoint')]][0][*].{OutputValue:OutputValue}" \
            --output text`

echo $NEWURL

##Check OS
unameOut="$(uname -s)"
case "${unameOut}" in
    Linux*)     machine=Linux;;
    Darwin*)    machine=Mac;;
    CYGWIN*)    machine=Cygwin;;
    MINGW*)     machine=MinGw;;
    *)          machine="UNKNOWN:${unameOut}"
esac
if [ "$machine" == "$MACOS" ]
then
    echo "You are on device with MAC OS, running sed with '' "
    sed  -i '' "s#{{URL}}#$NEWURL#g" $COLLECTION
else
    echo "You are not on a mac, running sed normally"
    sed  -i "s#{{URL}}#$NEWURL#g" $COLLECTION
fi
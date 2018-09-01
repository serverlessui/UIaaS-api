#!/bin/bash

##define inputs
STACK=$1
COLLECTION=$2
ENVIRONMENT=$3

integrationTests() {
  ./setup.sh "$STACK-$ENVIRONMENT" "$COLLECTION" 
  ./run.sh $COLLECTION
}

integrationTests
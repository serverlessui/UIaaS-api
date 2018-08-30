#!/bin/bash

##define inputs
STACK="uiaas-api-account"
COLLECTION=$1
ENVIRONMENT=$2

accountIntegration () {
  cd account && ./setup.sh "$STACK-$ENVIRONMENT" "$COLLECTION-$ENVIRONMENT" 
}

integrationTests() {
    accountIntegration
}

integrationTests
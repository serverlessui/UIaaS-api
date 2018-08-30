#!/bin/bash

##define inputs
STACK="uiaas-api-account"
COLLECTION=$1
ENVIRONMENT=$2

accountIntegration () {
  cd account && make integ STACK=$STACK COLLECTION="$COLLECTION-$ENVIRONMENT" 
}

integrationTests() {
    accountIntegration

}
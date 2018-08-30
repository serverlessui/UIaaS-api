#!/bin/bash

ENVIRONMENT=$1

deployStacks () {
  cd stacks && make deploy ENVIRONMENT=$ENVIRONMENT
}

deployAccount() {
    cd account && make deploy ENVIRONMENT=$ENVIRONMENT
}

deployStats() {
    cd stats && make deploy ENVIRONMENT=$ENVIRONMENT
}

deployFunctions() {
    deployStacks
    cd ..
    deployAccount
    cd ..
    deployStats
}

deployFunctions
